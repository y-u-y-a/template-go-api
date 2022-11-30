//go:generate mockgen -source=repositories.go -destination=../mock_domain/repositories.go

package domain

import (
	"context"
)

type DatabaseRepository interface {
	QueryInquiries(ctx context.Context) ([]*Inquiry, error)
}
