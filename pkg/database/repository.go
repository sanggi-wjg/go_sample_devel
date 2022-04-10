package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go_sample_devel/pkg/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"path"
)

type Repository struct {
	db           *gorm.DB
	logger       *logrus.Logger
	defaultJoins []string
}

func NewRepository(db *gorm.DB, defaultJoins ...string) *Repository {
	return &Repository{db, logrus.New(), defaultJoins}
}

func CreateMockRepository() *Repository {
	mockDB, err := gorm.Open(
		sqlite.Open(path.Join(util.GetBasePath(), "gorm.db")),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalf("fail to open gorm db: %v", err)
	}
	return &Repository{db: mockDB, logger: logrus.New()}
}

var (
	// ErrNotFound is a convenience reference for the actual GORM error
	ErrNotFound = gorm.ErrRecordNotFound
)

// FindAll find all entities
func (r *Repository) FindAll(entity interface{}) error {
	res := r.db.
		Unscoped().
		Find(entity)
	return r.handleError(res)
}

// FindById find by id entity
func (r *Repository) FindById(entity interface{}, id uint64) error {
	res := r.db.
		Where("id = ?", id).
		First(entity)
	return r.handleOneError(res)
}

func (r *Repository) FindPaged(entity interface{}, limit, offset int) error {
	res := r.db.
		Unscoped().
		Limit(limit).
		Offset(offset).
		Find(entity)
	return r.handleError(res)
}

func (r *Repository) FindWhere(entity interface{}, condition string) error {
	res := r.db.
		Where(condition).
		Find(entity)
	return r.handleError(res)
}

// Create : create entity
func (r *Repository) Create(entity interface{}) error {
	res := r.db.Create(entity)
	return r.handleError(res)
}

// Update : update entity
func (r *Repository) Update(entity interface{}) error {
	res := r.db.Save(entity)
	return r.handleError(res)
}

func (r Repository) Delete(entity interface{}) error {
	res := r.db.Delete(entity)
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
