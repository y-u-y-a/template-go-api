package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Inquiry struct {
	id        int
	email     string
	name      string
	tel       string
	content   string
	isConfirm bool
}

func NewInquiry(id int, email string, name string, tel string, content string, isConfirm bool) *Inquiry {
	return &Inquiry{
		id:        id,
		email:     email,
		name:      name,
		tel:       tel,
		content:   content,
		isConfirm: isConfirm,
	}
}

func (i *Inquiry) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(
			&i.name,
			validation.Required.Error("名前は必須入力です"),
		),
		validation.Field(
			&i.email,
			validation.Required.Error("メールアドレスは必須入力です"),
			is.Email.Error("メールアドレスを入力して下さい"),
		),
		validation.Field(
			&i.tel,
			validation.Required.Error("電話番号は必須入力です"),
		),
		validation.Field(
			&i.content,
			validation.Required.Error("お問い合わせ内容は必須入力です"),
		),
		// TODO: isConfirmのバリデーションをどうするか
	)
}

func (i *Inquiry) ID() int {
	return i.id
}

func (i *Inquiry) Email() string {
	return i.email
}

func (i *Inquiry) Name() string {
	return i.name
}

func (i *Inquiry) Tel() string {
	return i.tel
}

func (i *Inquiry) IsConfirm() bool {
	return i.isConfirm
}

func (i *Inquiry) Content() string {
	return i.content
}
