package main

import (
	"net"
	"time"
	"webot/config"
	botService "webot/grpc"
	"webot/logger"
	"webot/logger/hooks/sentryhook"
	botpb "webot/proto/bot/v1"
	sentry "webot/sentry"
	agentdata "webot/store/postgres/agentdata"
	agentsessiondata "webot/store/postgres/agentsession"
	categorydata "webot/store/postgres/category"
	chathistorydata "webot/store/postgres/chathistory"
	opsessiondata "webot/store/postgres/opsession"

	"webot/store/postgres"

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

	postgresClient, err := postgres.NewClient(config.Postgres())
	if err != nil {
		log.Panicf("main::error creating postgres client err=%+v", err)
	}

	agentRepo, err := agentdata.New(postgresClient)
	if err != nil {
		log.Panicf("main::error creating repo err=%+v", err)
	}

	agentSessionRepo, err := agentsessiondata.New(postgresClient)
	if err != nil {
		log.Panicf("main::error creating repo err=%+v", err)
	}

	chatHistoryRepo, err := chathistorydata.New(postgresClient)
	if err != nil {
		log.Panicf("main::error creating repo err=%+v", err)
	}

	opSessionRepo, err := opsessiondata.New(postgresClient)
	if err != nil {
		log.Panicf("main::error creating repo err=%+v", err)
	}

	categoryRepo, err := categorydata.New(postgresClient)
	if err != nil {
		log.Panicf("main::error creating repo err=%+v", err)
	}

	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	serv := botService.New(agentSessionRepo, agentRepo, chatHistoryRepo, opSessionRepo, categoryRepo)
	botpb.RegisterChatServiceServer(grpcServer, serv)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
