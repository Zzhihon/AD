package service

import (
	"AD/dto"
	"AD/storage"
	"fmt"
	"strconv"
	"time"
)

type ReportService struct {
	ReportRepo *storage.ReportRepository // 注入 ReportRepository
}

// NewReportService 创建 ReportService
func NewReportService(reportRepo *storage.ReportRepository) *ReportService {
	return &ReportService{ReportRepo: reportRepo}
}

// CreateReport 添加新医生
func (s *ReportService) CreateReport(request *dto.OTCFormRequest) error {
	// 这里可以加一些业务逻辑，比如数据验证等
	i, err := strconv.Atoi(request.PatientID)
	if err != nil {
		fmt.Println("转换失败:", err)
		return err
	}

	// 将 int 转换为 uint
	patientID := uint(i)
	report := &storage.OTCReport{
		PatienName: request.PatientName,
		DoctorName: request.DoctorName,
		PatientID:  patientID,
		ReportDate: time.Now(),
	}

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
