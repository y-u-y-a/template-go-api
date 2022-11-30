package usecase

import (
	"context"
	"y-u-y-a/template-go/domain"

	"go.uber.org/zap"
)

type InquiryUsecase struct {
	logger             *zap.Logger
	databaseRepository domain.DatabaseRepository
}

func NewInquiryUsecase(
	logger *zap.Logger,
	databaseRepository domain.DatabaseRepository,
) *InquiryUsecase {
	return &InquiryUsecase{
		logger,
		databaseRepository,
	}
}

// お問い合わせ情報を全件取得する
func (u *InquiryUsecase) GetInquiries(ctx context.Context) ([]*domain.Inquiry, error) {
	return u.databaseRepository.QueryInquiries(ctx)
}
