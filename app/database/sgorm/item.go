package sgorm

import "gorm.io/gorm"

type iRepository struct {
	db *gorm.DB
}
