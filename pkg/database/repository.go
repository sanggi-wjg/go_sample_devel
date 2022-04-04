package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type Repository struct {
	db           *gorm.DB
	logger       *logrus.Logger
	defaultJoins []string
}


func NewRepository(db *gorm.DB, defaultJoins ...string) *Repository {
	return &Repository{db, logrus.New(), defaultJoins}
}

var (
	// ErrNotFound is a convenience reference for the actual GORM error
	ErrNotFound = gorm.ErrRecordNotFound
)

// FindAll find all entities
func (r *Repository) FindAll(entity interface{}) error {
	res := r.db.Unscoped().Find(entity)
	return r.handleError(res)
}


// FindById find by id entity
func (r *Repository) FindById(entity interface{}, id uint64) error {
	res := r.db.Where("id = ?", id).First(entity)
	return r.handleOneError(res)
}


// Create : create entity
func (r *Repository) Create(entity interface{}) error {
	res := r.db.Create(entity)
	return r.handleError(res)
}


// Upsert update or create entity
func (r *Repository) Upsert(entity interface{}) error {
	res := r.db.Save(entity)
	return r.handleError(res)
}

// handleOneError handle one case
func (r *Repository) handleOneError(res *gorm.DB) error {
	if err := r.handleError(res); err != nil {
		return err
	}
	if res.RowsAffected != 1 {
		return ErrNotFound
	}
	return nil
}

// handleError handle common case
func (r *Repository) handleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("%w", res.Error)
		r.logger.Error(err)
		return res.Error
	}
	return nil
}
