package storage

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

// ReportRepository 负责数据库操作
type ReportRepository struct {
	db *gorm.DB
}

// NewReportRepository 创建 ReportRepository
func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetAllReports() (*OTCReport, error) {
	var reports OTCReport
	err := r.db.Find(&reports).Error
	if err != nil {
		log.Println("Error fetching reports:", err)
		return nil, err
	}
	return &reports, nil
}

func (r *ReportRepository) GetReportByID(reportID string) (*OTCReport, error) {
	var report OTCReport
	log.Println("医生的id" + reportID)
	err := r.db.Where("id = ?", reportID).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		log.Println("Error fetching report:", err)
		return nil, err
	}
	return &report, nil
}

// CreateReport 添加新医生
func (r *ReportRepository) CreateReport(report *OTCReport) error {
	log.Println("Creating report:", report)
	err := r.db.Create(report).Error
	if err != nil {
		log.Println("Error creating report:", err)
	}
	return err
}

// UpdateReport 更新医生信息
func (r *ReportRepository) UpdateReport(report *OTCReport) error {
	err := r.db.Save(report).Error
	if err != nil {
		log.Println("Error updating report:", err)
	}
	return err
}

// DeleteReport 删除医生
func (r *ReportRepository) DeleteReport(reportID string) error {
	err := r.db.Where("report_id = ?", reportID).Delete(&OTCReport{}).Error
	if err != nil {
		log.Println("Error deleting report:", err)
	}
	return err
}
