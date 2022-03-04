package userdata

import (
	"context"
	"time"
	store "webot/store"
	"webot/store/postgres"
)

// Data is struct used for emailRepo
type Data struct {
	store.AgentSessionRepo
	client *postgres.Client
}

// Generates new emailRepo to handle DB queries
func New(client *postgres.Client) (store.AgentSessionRepo, error) {
	return &Data{
		client: client,
	}, nil
}

type AgentSession struct {
	tableName     struct{}  `pg:"agent_session"`
	ID            int       `pg:"id,pk"`
	CreatedAt     time.Time `pg:"created_at"`
	UpdatedAt     time.Time `pg:"updated_at"`
	Deleted       bool      `pg:"deleted,use_zero"`
	AgentCode     string    `pg:"agent_code"`
	FromTime      int64     `pg:"from_time"`
	ToTime        int64     `pg:"to_time"`
	SessionStatus bool      `pg:"session_status,use_zero"`
}

func (d *Data) SaveAgentSession(ctx context.Context, userCode string) error {
	agentSession := new(AgentSession)
	agentSession.AgentCode = userCode
	current_time := time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	from_time := current_time.Unix()
	agentSession.FromTime = from_time
	agentSession.SessionStatus = true
	_, err := d.client.Model(agentSession).Insert()
	if err != nil {

		return nil
	}
	return nil
}

func (d *Data) UpdateAgentSession(ctx context.Context, userCode string) error {

	agentSession := new(AgentSession)
	err := d.client.Model(agentSession).Where("agent_code = ? and session_status = true", userCode).Select()
	if err != nil {
		return nil
	}

	current_time := time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	agentSession.UpdatedAt = current_time
	to_time := current_time.Unix()
	agentSession.ToTime = to_time
	agentSession.SessionStatus = false
	_, err = d.client.Model(agentSession).Where("id = ?", agentSession.ID).Update()
	if err != nil {
		return nil
	}
	return nil
}
