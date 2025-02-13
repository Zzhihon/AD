package service

import (
	"AD/storage"
	"errors"
	"log"
)

type DoctorService struct {
	DoctorRepo *storage.DoctorRepository // 注入 DoctorRepository
}

// NewDoctorService 创建 DoctorService
func NewDoctorService(doctorRepo *storage.DoctorRepository) *DoctorService {
	return &DoctorService{DoctorRepo: doctorRepo}
}

// CreateDoctor 添加新医生
func (s *DoctorService) CreateDoctor(doctor *storage.Doctor) error {
	// 这里可以加一些业务逻辑，比如数据验证等
	if doctor.Name == "" {
		return errors.New("name cannot be empty")
	}
	return s.DoctorRepo.CreateDoctor(doctor)
}

// GetDoctorByID 通过 DoctorID 获取医生
func (s *DoctorService) GetDoctorByID(doctorID string) (*storage.Doctor, error) {
	return s.DoctorRepo.GetDoctorByID(doctorID)
}

func (s *DoctorService) UpdateDoctor(doctor *storage.Doctor) error {
	// 可以加入一些业务验证或处理逻辑
	return s.DoctorRepo.UpdateDoctor(doctor)
}

func (s *DoctorService) GetPatientsByDoctorID(doctorID uint) ([]*storage.Patient, error) {
	log.Printf("Fetching patients for doctor ID: %d", doctorID)
	patients, err := s.DoctorRepo.GetPatientsByDoctorID(doctorID)
	if err != nil {
		return nil, err
	}
	return patients, nil
}
