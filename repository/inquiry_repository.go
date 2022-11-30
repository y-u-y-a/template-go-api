package repository

import (
	"context"
	"fmt"
	"y-u-y-a/template-go/domain"
	"y-u-y-a/template-go/ent"

	"github.com/k0kubun/pp"
)

// お問い合わせ情報を全件取得する
func (d *DatabaseRepository) QueryInquiries(ctx context.Context) ([]*domain.Inquiry, error) {
	client := ent.NewClient()
	inquiries, err := client.Inquiry.Query().All(ctx)
	if err != nil {
		fmt.Println("err >>>", err)
		return nil, err
	}
	pp.Println("inquiries", inquiries)
	// newInquiries := make([]*domain.Inquiry, len(inquiries))
	// for i, v := range inquiries {
	// 	inquiry := domain.NewInquiry(v.ID, v.Email, v.Name, v.Tel, v.Content, v.IsConfirm)
	// 	newInquiries[i] = inquiry
	// }
	// return newInquiries, nil
	return []*domain.Inquiry{}, nil
}
