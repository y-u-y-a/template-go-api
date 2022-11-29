// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"y-u-y-a/template-go/ent/inquiry"

	"entgo.io/ent/dialect/sql"
)

// Inquiry is the model entity for the Inquiry schema.
type Inquiry struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 名前
	Name string `json:"name,omitempty"`
	// メールアドレス
	Email string `json:"email,omitempty"`
	// 電話番号
	Tel string `json:"tel,omitempty"`
	// お問い合わせ内容
	Content string `json:"content,omitempty"`
	// 対応フラグ
	IsConfirm bool `json:"is_confirm,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Inquiry) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case inquiry.FieldIsConfirm:
			values[i] = new(sql.NullBool)
		case inquiry.FieldID:
			values[i] = new(sql.NullInt64)
		case inquiry.FieldName, inquiry.FieldEmail, inquiry.FieldTel, inquiry.FieldContent:
			values[i] = new(sql.NullString)
		case inquiry.FieldCreatedAt, inquiry.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Inquiry", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Inquiry fields.
func (i *Inquiry) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case inquiry.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case inquiry.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case inquiry.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case inquiry.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case inquiry.FieldEmail:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[j])
			} else if value.Valid {
				i.Email = value.String
			}
		case inquiry.FieldTel:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tel", values[j])
			} else if value.Valid {
				i.Tel = value.String
			}
		case inquiry.FieldContent:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[j])
			} else if value.Valid {
				i.Content = value.String
			}
		case inquiry.FieldIsConfirm:
			if value, ok := values[j].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_confirm", values[j])
			} else if value.Valid {
				i.IsConfirm = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Inquiry.
// Note that you need to call Inquiry.Unwrap() before calling this method if this Inquiry
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Inquiry) Update() *InquiryUpdateOne {
	return (&InquiryClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Inquiry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Inquiry) Unwrap() *Inquiry {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Inquiry is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Inquiry) String() string {
	var builder strings.Builder
	builder.WriteString("Inquiry(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(i.Email)
	builder.WriteString(", ")
	builder.WriteString("tel=")
	builder.WriteString(i.Tel)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(i.Content)
	builder.WriteString(", ")
	builder.WriteString("is_confirm=")
	builder.WriteString(fmt.Sprintf("%v", i.IsConfirm))
	builder.WriteByte(')')
	return builder.String()
}

// Inquiries is a parsable slice of Inquiry.
type Inquiries []*Inquiry

func (i Inquiries) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
