package handler

import (
	"fmt"
	"net/http"
	"y-u-y-a/template-go/usecase"

	chi "github.com/go-chi/chi/v5"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
)

type AdminHandler struct {
	logger         *zap.Logger
	inquiryUsecase *usecase.InquiryUsecase
}

func NewAdminHandler(
	logger *zap.Logger,
	InquiryUsecase *usecase.InquiryUsecase,
) *AdminHandler {
	return &AdminHandler{
		logger,
		InquiryUsecase,
	}
}

func (h *AdminHandler) Route() http.Handler {
	r := chi.NewRouter()
	r.Get("/inquiries", h.getInquiries)
	return r
}

// お問い合わせ情報を全件取得する
func (h *AdminHandler) getInquiries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	inquiries, err := h.inquiryUsecase.GetInquiries(ctx)
	if err != nil {
		fmt.Println("お問い合わせ情報の取得に失敗しました", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	pp.Println("inquiries", inquiries)
}
