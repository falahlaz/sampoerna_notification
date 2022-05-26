package firebase

import (
	"context"
	"fmt"
	"time"

	firebasego "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"gitlab.com/sholludev/sampoerna_notification/pkg/log"
)

var app *firebasego.App
var err error

func Init(ctx context.Context) {
	app, err = firebasego.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}
}

func (sn *SingleNotification) Send(ctx context.Context) (string, error) {
	var response string
	oneHour := time.Duration(1) * time.Hour

	client, err := app.Messaging(ctx)
	if err != nil {
		return response, err
	}

	response, err = client.Send(ctx, &messaging.Message{
		Token: sn.FCMToken,
		Data: map[string]string{
			"action":  sn.Message.Action,
			"message": sn.Message.Message,
			"type":    sn.Message.Type,
			"data":    fmt.Sprintf("%v", sn.Message.Data),
		},
		Android: &messaging.AndroidConfig{
			TTL:      &oneHour,
			Priority: "normal",
		},
	})

	// log
	log.InsertLogActivity(ctx, &log.LogActivity{
		Row:       response,
		Action:    "Send Single Notification",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return response, err
}
