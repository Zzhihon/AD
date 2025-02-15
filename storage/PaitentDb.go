package storage

import (
	"AD/dto"
	"errors"
	"gorm.io/gorm"
	"log"
)

// PatientRepository 负责数据库操作
type PatientRepository struct {
	db *gorm.DB
}

// NewPatientRepository 创建 PatientRepository
func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetAllPatients() (*Patient, error) {
	var patients Patient
	err := r.db.Find(&patients).Error
	if err != nil {
		log.Println("Error fetching patients:", err)
		return nil, err
	}
	return &patients, nil
}

func (r *PatientRepository) GetPatientByID(patientID string) (*Patient, error) {
	var patient Patient
	log.Println("医生的id" + patientID)
	err := r.db.Where("id = ?", patientID).First(&patient).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		log.Println("Error fetching patient:", err)
		return nil, err
	}
	return &patient, nil
}

// CreatePatient 添加新医生
func (r *PatientRepository) CreatePatient(request *dto.CreatePatientRequest) error {
	// 创建病人
	patient := Patient{
		Name:   request.Name,
		Age:    request.Age,
		Gender: request.Gender,
	}

	// 插入病人到数据库
	if err := r.db.Create(&patient).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	// 查找指定的医生
	var doctor Doctor
	if err := r.db.First(&doctor, request.DoctorID).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	// 将病人与医生关联（即将病人插入关联表 doctor_patients 中）
	if err := r.db.Model(&patient).Association("Doctors").Append(&doctor); err != nil {
		log.Println("Association" + err.Error())
		return err
	}

	return nil
}

// UpdatePatient 更新医生信息
func (r *PatientRepository) UpdatePatient(patient *Patient) error {
	err := r.db.Save(patient).Error
	if err != nil {
		log.Println("Error updating patient:", err)
	}
	return err
}

// DeletePatient 删除医生
func (r *PatientRepository) DeletePatient(patientID string) error {
	err := r.db.Where("patient_id = ?", patientID).Delete(&Patient{}).Error
	if err != nil {
		log.Println("Error deleting patient:", err)
	}
	return err
}
