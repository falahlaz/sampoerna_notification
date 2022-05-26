package log

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
)

const (
	INDEX_LOG_ERROR    = "sampoerna_notification_log_error"
	INDEX_LOG_ACTIVITY = "sampoerna_notification_log_activity"
)

var client *elasticsearch.Client

func Init() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			environment.Get("ELASTIC_HOST_1"),
			environment.Get("ELASTIC_HOST_2"),
			environment.Get("ELASTIC_HOST_3"),
		},
		Username: environment.Get("ELASTIC_USER"),
		Password: environment.Get("ELASTIC_PASS"),
	}
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
}

func Insert(c context.Context, index string, log interface{}) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: shortid.MustGenerate(),
		Body:       strings.NewReader(util.ToJSON(log)),
		Refresh:    "true",
	}

	if _, err := req.Do(c, client); err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot insert data",
			"Index":         index,
			"Data":          log,
		}).Error(err.Error())

		return err
	}

	return nil
}

func InsertLogActivity(c context.Context, log *LogActivity) error {
	return Insert(c, INDEX_LOG_ACTIVITY, log)
}

func InsertLogError(c context.Context, log *LogError) error {
	return Insert(c, INDEX_LOG_ERROR, log)
}
