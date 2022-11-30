package repository

import "y-u-y-a/template-go/domain"

type DatabaseRepository struct{}

func NewDatabaseRepository() domain.DatabaseRepository {
	return &DatabaseRepository{}
}
