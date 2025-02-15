package service

import (
	"AD/dto"
	"AD/storage"
	"errors"
)

type PatientService struct {
	PatientRepo *storage.PatientRepository // 注入 PatientRepository
}

// NewPatientService 创建 PatientService
func NewPatientService(patientRepo *storage.PatientRepository) *PatientService {
	return &PatientService{PatientRepo: patientRepo}
}

// CreatePatient 添加新医生
func (s *PatientService) CreatePatient(request *dto.CreatePatientRequest) error {
	// 这里可以加一些业务逻辑，比如数据验证等
	if request.Name == "" {
		return errors.New("name cannot be empty")
	}

	return s.PatientRepo.CreatePatient(request)
}

func (s *PatientService) GetPatientByID(patientID string) (*storage.Patient, error) {
	return s.PatientRepo.GetPatientByID(patientID)
}

func (s *PatientService) UpdatePatient(patient *storage.Patient) error {
	// 可以加入一些业务验证或处理逻辑
	return s.PatientRepo.UpdatePatient(patient)
}
