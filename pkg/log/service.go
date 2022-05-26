package log

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
)

const (
	INDEX_LOG_ERROR    = "sampoerna_master_sampoerna_log_error"
	INDEX_LOG_ACTIVITY = "sampoerna_master_sampoerna_log_activity"
	INDEX_LOG_LOGIN    = "sampoerna_master_sampoerna_log_login"
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

func Insert(c echo.Context, index string, log interface{}) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: shortid.MustGenerate(),
		Body:       strings.NewReader(util.ToJSON(log)),
		Refresh:    "true",
	}

	if _, err := req.Do(c.Request().Context(), client); err != nil {
		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot insert data",
			"Index":         index,
			"Data":          log,
		}).Error(err.Error())

		return err
	}

	return nil
}

func InsertLogActivity(c echo.Context, log *LogActivity) error {
	return Insert(c, INDEX_LOG_ACTIVITY, log)
}

func InsertLogError(c echo.Context, log *LogError) error {
	return Insert(c, INDEX_LOG_ERROR, log)
}

func InsertLogLogin(c echo.Context, log *LogLogin) error {
	return Insert(c, INDEX_LOG_LOGIN, log)
}
