package grpc

import (
	"context"
	"io"
	"math/rand"
	botpb "webot/proto/bot/v1"
	"webot/store"
)

type chatServiceServer struct {
	botpb.UnimplementedChatServiceServer
	channel          map[string][]chan *botpb.Message
	activeUsers      map[string]chan *botpb.Message
	activeAgents     map[string]chan *botpb.Message
	chatRepo         store.ChatRepo
	agentRepo        store.AgentRepo
	agentSessionRepo store.AgentSessionRepo
	opSessionRepo    store.OpSessionRepo
}

func New(agentSessionRepo store.AgentSessionRepo, agentRepo store.AgentRepo, chatRepo store.ChatRepo, opSessionRepo store.OpSessionRepo) *chatServiceServer {
	s := &chatServiceServer{
		UnimplementedChatServiceServer: botpb.UnimplementedChatServiceServer{},
		channel:                        map[string][]chan *botpb.Message{},
		activeUsers:                    map[string]chan *botpb.Message{},
		activeAgents:                   map[string]chan *botpb.Message{},
		chatRepo:                       chatRepo,
		agentRepo:                      agentRepo,
		agentSessionRepo:               agentSessionRepo,
		opSessionRepo:                  opSessionRepo,
	}
	return s
}

func (s *chatServiceServer) RemoveChannel(channelSlice []chan *botpb.Message, channel chan *botpb.Message) []chan *botpb.Message {
	for i, elem := range channelSlice {
		if elem == channel {
			return append(channelSlice[:i], channelSlice[i+1:]...)
		}
	}
	return channelSlice
}

func (s *chatServiceServer) getRandomAgent() string {
	m := s.activeAgents
	mapKeys := make([]string, 0, len(m))
	for key := range m {
		mapKeys = append(mapKeys, key)
	}
	return mapKeys[rand.Intn(len(mapKeys))]
}

func (s *chatServiceServer) JoinChannel(ch *botpb.Channel, msgStream botpb.ChatService_JoinChannelServer) error {
	invite := botpb.Message{}
	invite.Sender = ch.GetUserCode()
	invite.Message = "wait"
	invite.Channel = ch
	invite.SenderType = ch.GetUserType()

	ctx := context.Background()
	var randomAgent string
	msgChannel := make(chan *botpb.Message)
	msgStream.Send(&invite)
	if ch.UserType == "OPERATOR" {
		s.activeUsers[ch.UserCode] = msgChannel
		randomAgent = s.getRandomAgent()
		agentMessageChannel := s.activeAgents[randomAgent]
		s.channel[ch.UserCode] = append(s.channel[ch.UserCode], agentMessageChannel)
		s.channel[randomAgent] = append(s.channel[ch.UserCode], msgChannel)
		s.agentRepo.UpdateChannelCount(ctx, randomAgent, len(s.channel[randomAgent])-1)
		s.opSessionRepo.SaveOperatorSession(ctx, ch.UserCode, randomAgent)
	} else if ch.UserType == "AGENT" {
		s.agentRepo.SaveChatUserStatus(ctx, ch)
		s.agentSessionRepo.SaveAgentSession(ctx, ch.UserCode)
		s.activeAgents[ch.UserCode] = msgChannel
	}

	// doing this never closes the stream

	for {
		select {
		case <-msgStream.Context().Done():
			if ch.UserType == "AGENT" {
				delete(s.activeAgents, ch.UserCode)
				s.agentRepo.UpdateChannelCount(ctx, ch.UserCode, 0)
				s.agentSessionRepo.UpdateAgentSession(ctx, ch.UserCode)
				s.agentRepo.DropChatUserStatus(ctx, ch.UserCode)
			} else if ch.UserType == "OPERATOR" {
				delete(s.activeUsers, ch.UserCode)
				s.channel[randomAgent] = s.RemoveChannel(s.channel[randomAgent], msgChannel)
				s.agentRepo.UpdateChannelCount(ctx, randomAgent, len(s.channel[randomAgent])-1)
				s.opSessionRepo.UpdateOperatorSession(ctx, ch.UserCode, randomAgent)
			}
			return nil
		case msg := <-msgChannel:
			msgStream.Send(msg)
		}
	}

}

func (s *chatServiceServer) SendMessage(msgStream botpb.ChatService_SendMessageServer) error {
	msg, err := msgStream.Recv()

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	ack := botpb.MessageAck{Status: "SENT"}
	msgStream.SendAndClose(&ack)

	go func() {
		streams := s.channel[msg.Channel.UserCode]

		for _, msgChan := range streams {
			msgChan <- msg
		}
	}()

	return nil
}
