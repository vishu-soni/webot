package userdata

import (
	"context"
	"time"
	store "webot/store"
	"webot/store/postgres"
)

// Data is struct used for emailRepo
type Data struct {
	store.OpSessionRepo
	client *postgres.Client
}

// Generates new emailRepo to handle DB queries
func New(client *postgres.Client) (store.OpSessionRepo, error) {
	return &Data{
		client: client,
	}, nil
}

type OperatorSession struct {
	tableName    struct{}  `pg:"op_session"`
	ID           int       `pg:"id,pk"`
	CreatedAt    time.Time `pg:"created_at"`
	UpdatedAt    time.Time `pg:"updated_at"`
	Deleted      bool      `pg:"deleted,use_zero"`
	AgentCode    string    `pg:"agent_code"`
	FromTime     int64     `pg:"from_time"`
	ToTime       int64     `pg:"to_time"`
	OperatorCode string    `pg:"operator_code"`
}

func (d *Data) SaveOperatorSession(ctx context.Context, opCode string, agentCode string) error {
	operatorSession := new(OperatorSession)
	operatorSession.AgentCode = agentCode
	operatorSession.OperatorCode = opCode
	current_time := time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	from_time := current_time.Unix()
	operatorSession.FromTime = from_time
	_, err := d.client.Model(operatorSession).Insert()
	if err != nil {
		return nil
	}
	return nil
}

func (d *Data) UpdateOperatorSession(ctx context.Context, opCode string, agentCode string) error {
	operatorSession := new(OperatorSession)
	err := d.client.Model(operatorSession).Where("agent_code = ? AND operator_code = ?", agentCode, opCode).Select()
	if err != nil {
		return nil
	}
	current_time := time.Now().Add(time.Hour*time.Duration(5) + time.Second*1800)
	operatorSession.UpdatedAt = current_time
	to_time := current_time.Unix()
	operatorSession.ToTime = to_time
	_, err = d.client.Model(operatorSession).Where("id = ?", operatorSession.ID).Update()
	if err != nil {
		return nil
	}
	return nil
}
