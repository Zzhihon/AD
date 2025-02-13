package service

import (
	"AD/storage"
)

type ReportService struct {
	ReportRepo *storage.ReportRepository // 注入 ReportRepository
}

// NewReportService 创建 ReportService
func NewReportService(reportRepo *storage.ReportRepository) *ReportService {
	return &ReportService{ReportRepo: reportRepo}
}

// CreateReport 添加新医生
func (s *ReportService) CreateReport(report *storage.OTCReport) error {
	// 这里可以加一些业务逻辑，比如数据验证等

	return s.ReportRepo.CreateReport(report)
}

// GetReportByID 通过 ReportID 获取医生
func (s *ReportService) GetReportByID(reportID string) (*storage.OTCReport, error) {
	return s.ReportRepo.GetReportByID(reportID)
}

func (s *ReportService) UpdateReport(report *storage.OTCReport) error {
	// 可以加入一些业务验证或处理逻辑
	return s.ReportRepo.UpdateReport(report)
}
