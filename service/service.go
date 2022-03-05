package grpc

import (
	"context"
	botpb "webot/proto/bot/v1"

	log "github.com/sirupsen/logrus"
)

type message struct {
	UserCode string
	Body     string
}

type handler struct {
	botpb.ServiceServer
	messages     chan message
	userMapping  map[string]string
	agentMapping map[string]string
	userChannel  map[string]chan message
}

func New() *handler {
	h := &handler{
		messages:     make(chan message, 128),
		userMapping:  map[string]string{},
		agentMapping: map[string]string{},
		userChannel:  map[string]chan message{},
	}
	go h.handleMessages()
	return h
}

func (h *handler) Chit(ctx context.Context, req *botpb.ChitRequest) (*botpb.ChitResponse, error) {
	log.Infof("chit:user-%s msg-%s", req.GetUserCode(), req.GetMessage())
	h.messages <- message{
		UserCode: req.GetUserCode(),
		Body:     req.GetMessage(),
	}
	return &botpb.ChitResponse{
		Response: "ok",
	}, nil
}

func (h *handler) Chat(req *botpb.ChatRequest, stream botpb.Service_ChatServer) error {
	channel := make(chan message, 10)
	h.userChannel[req.GetUserCode()] = channel
	for msg := range channel {
		stream.Send(&botpb.ChatResponse{
			Message: msg.Body,
		})
	}
	return nil
}

func (h *handler) InitiateChat(ctx context.Context, req *botpb.InitiateChatRequest) (*botpb.IntiateChatResponse, error) {
	h.userMapping[req.GetUser()] = req.GetAgent()
	h.userMapping[req.GetAgent()] = req.GetUser()
	return &botpb.IntiateChatResponse{Response: "ok"}, nil
}

func (h *handler) handleMessages() {
	for msg := range h.messages {
		log.Infof("handleMessages:user-%s msg-%s", msg.UserCode, msg.Body)
		agentCode, ok := h.userMapping[msg.UserCode]
		if ok {
			agentChannel, ok := h.userChannel[agentCode]
			if !ok {
				log.Infof("handleMessages:anf-%s", agentCode)
				continue
			}
			agentChannel <- msg
			continue
		}
		userCode, ok := h.userMapping[msg.UserCode]
		if ok {
			userChannel, ok := h.userChannel[userCode]
			if !ok {
				log.Infof("handleMessages:aunf-%s", userCode)
				continue
			}
			userChannel <- msg
			continue
		}

	}
}
