package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repository struct {
	db           *gorm.DB
	logger       *logrus.Logger
	defaultJoins []string
}

func NewRepository(db *gorm.DB, defaultJoins ...string) *repository {
	return &repository{db, logrus.New(), defaultJoins}
}

func (r *repository) FindAll(model interface{}) error {
	res := r.db.Unscoped().Find(model)
	return r.HandleError(res)
}

func (r *repository) HandleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("%w", res.Error)
		r.logger.Error(err)
		return res.Error
	}
	return nil
}
