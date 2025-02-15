package storage

import (
	"AD/dto"
	"errors"
	"gorm.io/gorm"
	"log"
	"strings"
)

// ReportRepository 负责数据库操作
type ReportRepository struct {
	db *gorm.DB
}

func (r *ReportRepository) Search(req dto.SearchRequest) ([]OTCReport, error) {
	// 显式指定 otc_reports 表和 patients 表的别名
	query := r.db.Table("otc_reports").
		Joins("LEFT JOIN patients ON otc_reports.patient_id = patients.id")

	// 患者名字（模糊搜索）
	if req.PatientName != "" {
		query = query.Where("patients.name LIKE ?", "%"+req.PatientName+"%")
	}

	// 性别（男或女）
	if req.Gender != "" {
		query = query.Where("patients.gender = ?", req.Gender)
	}

	// 年龄区间（56-60，61-65，66-70）
	if req.AgeRange != "" {
		ageRange := strings.Split(req.AgeRange, "-")
		if len(ageRange) == 2 {
			query = query.Where("patients.age BETWEEN ? AND ?", ageRange[0], ageRange[1])
		}
	}

	// OTC图像状态
	if req.OTCImageStatus != nil {
		query = query.Where("otc_reports.otc_image_status = ?", *req.OTCImageStatus)
	}

	// 报告状态
	if req.PredictionStatus != nil {
		query = query.Where("otc_reports.report_status = ?", *req.PredictionStatus)
	}

	// 执行查询
	var reports []OTCReport
	if err := query.Preload("Patient").Find(&reports).Error; err != nil {
		log.Println("Failed to execute search query:", err)
		return nil, err
	}

	return reports, nil
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
