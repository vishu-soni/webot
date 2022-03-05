package chathistory

import (
	"context"
	"time"
	model "webot/proto/bot/v1"
	store "webot/store"
	"webot/store/postgres"
)

// Emailtemplates is schema for email_templates table in postgres
type ChatHistory struct {
	tableName   struct{}  `pg:"chat_history"`
	ID          int       `pg:"id,pk"`
	CreatedAt   time.Time `pg:"created_at"`
	UpdatedAt   time.Time `pg:"updated_at"`
	Deleted     bool      `pg:"deleted,use_zero"`
	SentFrom    string    `pg:"sent_from"`
	SentTo      string    `pg:"sent_to"`
	Message     string    `pg:"message"`
	SentTime    string    `pg:"sent_time"`
	ChannelName string    `pg:"channel_name"`
	Status      string    `pg:"status"`
}

// Data is struct used for emailRepo
type Data struct {
	store.ChatRepo
	client *postgres.Client
}

// Generates new emailRepo to handle DB queries
func New(client *postgres.Client) (store.ChatRepo, error) {
	return &Data{
		client: client,
	}, nil
}

// QueryByTemplateName gives values for a specific template
func (d *Data) QueryByUserCodeAndType(ctx context.Context, userCode string, userType string) (*model.ChatUsers, error) {
	
	return nil, nil
}
