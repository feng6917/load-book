package utils

import "gorm.io/gorm"

func Page(db *gorm.DB, limit, offset int) *gorm.DB {
	if limit < 1 {
		limit = 10
	}
	return db.Limit(limit).Offset(offset)
}
