package grpc

// import (
// 	"io"
// 	"math/rand"
// 	"sync"
// 	botpb "webot/proto/bot/v1"
// 	"webot/store"

// 	log "github.com/sirupsen/logrus"
// )

// var m sync.Mutex

// type chatServiceServer struct {
// 	botpb.ChatServiceServer

// 	opChannel    map[string]string
// 	agChannel    map[string]string
// 	activeUsers  map[string]*botpb.ChatService_ChitChatServer
// 	activeAgents map[string]*botpb.ChatService_ChitChatServer

// 	chatRepo         store.ChatRepo
// 	agentRepo        store.AgentRepo
// 	agentSessionRepo store.AgentSessionRepo
// 	opSessionRepo    store.OpSessionRepo
// }

// func New(agentSessionRepo store.AgentSessionRepo, agentRepo store.AgentRepo, chatRepo store.ChatRepo, opSessionRepo store.OpSessionRepo) botpb.ChatServiceServer {
// 	s := &chatServiceServer{
// 		opChannel:        make(map[string]string),
// 		agChannel:        make(map[string]string),
// 		activeUsers:      make(map[string]*botpb.ChatService_ChitChatServer),
// 		activeAgents:     make(map[string]*botpb.ChatService_ChitChatServer),
// 		chatRepo:         chatRepo,
// 		agentRepo:        agentRepo,
// 		agentSessionRepo: agentSessionRepo,
// 		opSessionRepo:    opSessionRepo,
// 	}
// 	return s
// }

// func (s *chatServiceServer) getRandomAgent() string {
// 	m := s.activeAgents
// 	mapKeys := make([]string, 0, len(m))
// 	for key := range m {
// 		mapKeys = append(mapKeys, key)
// 	}
// 	return mapKeys[rand.Intn(len(mapKeys))]
// }

// func (s *chatServiceServer) ChitChat(ping botpb.ChatService_ChitChatServer) error {
// 	ctx := ping.Context()

// 	for {

// 		// exit if context is done
// 		// or continue
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 		}

// 		// receive data from stream
// 		req, err := ping.Recv()
// 		if err == io.EOF {
// 			// return will close stream from server side
// 			log.Info("ChitChat exit")
// 			return nil
// 		}
// 		if err != nil {
// 			log.Info("ChitChat: received error %v", err)
// 			continue
// 		}
// 		if req.GetSenderType() == "OPERATOR" && req.GetType() == "INFO" {
// 			userCode := req.GetSender()
// 			s.activeUsers[userCode] = &ping
// 			newAgent := s.getRandomAgent()
// 			s.opChannel[userCode] = newAgent
// 			s.agChannel[newAgent] = userCode
// 			continue
// 		} else if req.GetSenderType() == "AGENT" && req.GetType() == "INFO" {
// 			userCode := req.GetSender()
// 			s.activeAgents[userCode] = &ping
// 			continue
// 		} else if req.GetSenderType() == "OPERATOR" {
// 			agent := s.opChannel[req.GetSender()]
// 			agentChannel := s.activeAgents[agent]
// 			res := botpb.Chat{
// 				Sender:     req.GetMessage(),
// 				SenderType: req.GetSenderType(),
// 				Message:    req.GetMessage(),
// 			}
// 			ag := *agentChannel
// 			m.Lock()
// 			ag.Send(&res)
// 			m.Unlock()
// 			ag = nil

// 		} else if req.GetSenderType() == "AGENT" {
// 			user := s.agChannel[req.GetSender()]
// 			userChannel := s.activeUsers[user]
// 			res := botpb.Chat{
// 				Sender:     req.GetSender(),
// 				SenderType: req.GetSenderType(),
// 				Message:    req.GetMessage(),
// 			}
// 			ug := *userChannel
// 			m.Lock()
// 			ug.Send(&res)
// 			m.Unlock()
// 			ug = nil
// 		}

// 		resp := botpb.Chat{
// 			Sender:     req.GetSender(),
// 			SenderType: req.GetSenderType(),
// 			Message:    req.GetMessage(),
// 		}
// 		m.Lock()
// 		if err := ping.Send(&resp); err != nil {
// 			log.Info("send error %v", err)
// 		}
// 		m.Unlock()
// 	}
// }
