module github.com/VMAnalytic/alarm-bot

go 1.16

require (
	cloud.google.com/go/firestore v1.5.0
	firebase.google.com/go v3.13.0+incompatible
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	github.com/joho/godotenv v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/api v0.54.0
	google.golang.org/grpc v1.39.1
	gopkg.in/tucnak/telebot.v2 v2.3.5
)
