package storage

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

// DoctorRepository 负责数据库操作
type DoctorRepository struct {
	db *gorm.DB
}

// NewDoctorRepository 创建 DoctorRepository
func NewDoctorRepository(db *gorm.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) GetAllDoctors() (*Doctor, error) {
	var doctors Doctor
	err := r.db.Find(&doctors).Error
	if err != nil {
		log.Println("Error fetching doctors:", err)
		return nil, err
	}
	return &doctors, nil
}

func (r *DoctorRepository) GetDoctorByID(doctorID string) (*Doctor, error) {
	var doctor Doctor
	log.Println("医生的id" + doctorID)
	err := r.db.Where("id = ?", doctorID).First(&doctor).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		log.Println("Error fetching doctor:", err)
		return nil, err
	}
	return &doctor, nil
}

// CreateDoctor 添加新医生
func (r *DoctorRepository) CreateDoctor(doctor *Doctor) error {
	log.Println("Creating doctor:", doctor)
	err := r.db.Create(doctor).Error
	if err != nil {
		log.Println("Error creating doctor:", err)
	}
	return err
}

// UpdateDoctor 更新医生信息
func (r *DoctorRepository) UpdateDoctor(doctor *Doctor) error {
	err := r.db.Save(doctor).Error
	if err != nil {
		log.Println("Error updating doctor:", err)
	}
	return err
}

// DeleteDoctor 删除医生
func (r *DoctorRepository) DeleteDoctor(doctorID string) error {
	err := r.db.Where("doctor_id = ?", doctorID).Delete(&Doctor{}).Error
	if err != nil {
		log.Println("Error deleting doctor:", err)
	}
	return err
}

func (r *DoctorRepository) GetPatientsByDoctorID(doctorID uint) ([]*Patient, error) {
	var patients []*Patient
	// 查询 doctor_patients 表获取所有关联的病人ID
	if err := r.db.Table("doctor_patients").
		Where("doctor_id = ?", doctorID).
		Joins("JOIN patients ON patients.id = doctor_patients.patient_id").
		Select("patients.id, patients.name"). // 选择需要的字段
		Scan(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
