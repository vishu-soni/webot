package store

import (
	"context"
	model "webot/proto/bot/v1"
)

type ChatRepo interface {
	QueryByUserCodeAndType(ctx context.Context, userCode string, userType string) (*model.ChatUsers, error)
}

type AgentRepo interface {
	SaveChatUserStatus(ctx context.Context, userData *model.Channel) (int, error)
	UpdateChannelCount(ctx context.Context, userData string, n int) (int, error)
	DropChatUserStatus(ctx context.Context, userCode string) error
}

type AgentSessionRepo interface {
	SaveAgentSession(ctx context.Context , userCode string) error
	UpdateAgentSession(ctx context.Context, userCode string) error
}

type OpSessionRepo interface {
	SaveOperatorSession(ctx context.Context, opCode string , agent string ) error
	UpdateOperatorSession(ctx context.Context, opCode string , agent string ) error
}