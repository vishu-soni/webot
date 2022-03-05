package userdata

import (
	"context"
	"fmt"
	"time"
	store "webot/store"
	"webot/store/postgres"
)

// Data is struct used for emailRepo
type Data struct {
	store.CategoryRepo
	client *postgres.Client
}

// Generates new emailRepo to handle DB queries
func New(client *postgres.Client) (store.CategoryRepo, error) {
	return &Data{
		client: client,
	}, nil
}

// Emailtemplates is schema for email_templates table in postgres
type Category struct {
	tableName           struct{}  `pg:"category"`
	ID                  int       `pg:"id,pk"`
	CreatedAt           time.Time `pg:"created_at"`
	UpdatedAt           time.Time `pg:"updated_at"`
	Deleted             bool      `pg:"deleted,use_zero"`
	ProductType         string    `pg:"product_type"`
	Query               string    `pg:"query"`
	Response            []string  `pg:"response,array"`
	AutoTicketGenerated bool      `pg:"auto_ticket_generated"`
}

func (d *Data) GetQueryResponse(ctx context.Context, query string) ([]string, error) {
	categoryData := new(Category)
	err := d.client.Model(categoryData).Where("query = ?", query).Select()
	fmt.Printf("categoryData %v", categoryData.Response)
	if err != nil {
		fmt.Printf("error---- %v", err)
		return nil, fmt.Errorf("NOT_EXISTS")
	}

	return categoryData.Response, nil
}
