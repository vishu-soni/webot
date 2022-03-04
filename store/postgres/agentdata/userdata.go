package userdata

import (
	"context"
	"fmt"
	"time"
	model "webot/proto/bot/v1"
	store "webot/store"
	"webot/store/postgres"
)

// Data is struct used for emailRepo
type Data struct {
	store.AgentRepo
	client *postgres.Client
}

// Generates new emailRepo to handle DB queries
func New(client *postgres.Client) (store.AgentRepo, error) {
	return &Data{
		client: client,
	}, nil
}

// Emailtemplates is schema for email_templates table in postgres
type ChatUserData struct {
	tableName struct{}  `pg:"agent_live_data"`
	ID        int       `pg:"id,pk"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	Deleted   bool      `pg:"deleted,use_zero"`
	AgentCode string    `pg:"agent_code"`
	AgentType string    `pg:"agent_type"`
	Status    bool      `pg:"status,use_zero"`
	Channels  int       `pg:"channels"`
}

func (d *Data) SaveChatUserStatus(ctx context.Context, userData *model.Channel) (int, error) {
	agentLiveData := new(ChatUserData)
	err := d.client.Model(agentLiveData).Where("agent_code = ?", userData.UserCode).Select()
	if err != nil {
		agentLiveData.AgentCode = userData.UserCode
		agentLiveData.AgentType = userData.UserType
		agentLiveData.Status = userData.Status
		if _, err := d.client.Model(agentLiveData).Insert(); err != nil {
			return 0, nil
		}
		return 0, nil
	}

	agentLiveData.UpdatedAt = time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	agentLiveData.Status = true
	_, err = d.client.Model(agentLiveData).Where("agent_code = ?", userData.UserCode).Update()
	if err != nil {
		return 0, nil
	}
	return 0, nil
}

func (d *Data) DropChatUserStatus(ctx context.Context, userCode string) error {
	agentLiveData := new(ChatUserData)
	err := d.client.Model(agentLiveData).Where("agent_code = ?", userCode).Select()
	if err != nil {
		return nil
	}
	agentLiveData.UpdatedAt = time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	agentLiveData.Status = false
	agentLiveData.Channels = 0

	_, err = d.client.Model(agentLiveData).Where("id = ?", agentLiveData.ID).Update()
	if err != nil {
		return nil
	}
	return nil
}

func (d *Data) UpdateChannelCount(ctx context.Context, userCode string, n int) (int, error) {
	fmt.Printf("userChannels are %v\n", n)
	agentLiveData := new(ChatUserData)
	err := d.client.Model(agentLiveData).Where("agent_code = ?", userCode).Select()
	if err != nil {
		return 0, nil
	}

	agentLiveData.UpdatedAt = time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	agentLiveData.Channels = n
	_, err = d.client.Model(agentLiveData).Where("id = ?", agentLiveData.ID).Update()
	if err != nil {
		return 0, nil
	}
	return 0, nil
}
