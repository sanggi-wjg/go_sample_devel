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

var (
	// ErrNotFound is a convenience reference for the actual GORM error
	ErrNotFound = gorm.ErrRecordNotFound
)

func (r *repository) FindAll(entity interface{}) error {
	res := r.db.Unscoped().Find(entity)
	return r.handleError(res)
}

func (r *repository) FindById(entity interface{}, id uint64) error {
	res := r.db.Where("id = ?", id).First(entity)
	return r.handleOneError(res)
}

func (r *repository) Create(entity interface{}) error {
	res := r.db.Create(entity)
	return r.handleError(res)
}

func (r *repository) handleOneError(res *gorm.DB) error {
	if err := r.handleError(res); err != nil {
		return err
	}
	if res.RowsAffected != 1 {
		return ErrNotFound
	}
	return nil
}

func (r *repository) handleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("%w", res.Error)
		r.logger.Error(err)
		return res.Error
	}
	return nil
}
