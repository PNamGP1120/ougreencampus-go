package system

import "gorm.io/gorm"

type ReportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{db: db}
}

func (r *ReportService) Count(table string) (int64, error) {
	var count int64
	err := r.db.Table(table).Count(&count).Error
	return count, err
}
