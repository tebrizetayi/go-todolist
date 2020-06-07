package todolist

import (
	"go-todolist/config"

	"github.com/jinzhu/gorm"
)

type Service struct {
	cnf *config.Config
	db  *gorm.DB
}

func NewService(cnf *config.Config, db *gorm.DB) *Service {
	return &Service{
		cnf: cnf,
		db:  db,
	}
}

// GetConfig returns config.Config instance
func (s *Service) GetConfig() *config.Config {
	return s.cnf
}
