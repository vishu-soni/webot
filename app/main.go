package main

import (
	"net"
	"time"
	logger "webot/logger"
	sentryhook "webot/logger/hooks/sentryhook"
	botpb "webot/proto/bot/v1"
	sentry "webot/sentry"
	service "webot/service"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := sentry.Init(); err != nil {
		log.WithError(err).Error("sentry initialization failed")
	} else {
		logger.AddHook(sentryhook.New([]log.Level{log.ErrorLevel, log.PanicLevel, log.FatalLevel}))
		defer sentry.Flush(2 * time.Second)
		log.Info("main::sentry init success")
	}

	// postgresClient, err := postgres.NewClient(config.Postgres())
	// if err != nil {
	// 	log.Panicf("main::error creating postgres client err=%+v", err)
	// }

	// agentRepo, err := agentdata.New(postgresClient)
	// if err != nil {
	// 	log.Panicf("main::error creating repo err=%+v", err)
	// }

	// agentSessionRepo, err := agentsessiondata.New(postgresClient)
	// if err != nil {
	// 	log.Panicf("main::error creating repo err=%+v", err)
	// }

	// chatHistoryRepo, err := chathistorydata.New(postgresClient)
	// if err != nil {
	// 	log.Panicf("main::error creating repo err=%+v", err)
	// }

	// opSessionRepo, err := opsessiondata.New(postgresClient)
	// if err != nil {
	// 	log.Panicf("main::error creating repo err=%+v", err)
	// }

	// categoryRepo, err := categorydata.New(postgresClient)
	// if err != nil {
	// 	log.Panicf("main::error creating repo err=%+v", err)
	// }

	handler := service.New()
	Server := grpc.NewServer()
	botpb.RegisterServiceServer(Server, handler)
	reflection.Register(Server)

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	Server.Serve(lis)

}
