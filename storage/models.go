package storage

import (
	"gorm.io/gorm"
	"time"
)

// Prediction 结构体，存储 AI 预测结果
type Prediction struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	OTCReportID uint       `gorm:"not null" json:"otc_report_id"`            // 外键字段
	OTCReport   *OTCReport `gorm:"foreignKey:OTCReportID" json:"otc_report"` // 通过外键字段关联 OTCReport	Advice      string     `gorm:"size:200"`
	ImageID     string     `gorm:"size:200" json:"image_id"`
	Probability string     `gorm:"size:200" json:"probability"`
}

type Patient struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Name       string      `gorm:"type:varchar(100);not null" json:"name"`
	Age        int         `gorm:"not null" json:"age"`
	Gender     string      `gorm:"type:varchar(10);not null" json:"gender"`
	OTCReports []OTCReport `gorm:"foreignKey:PatientID;references:ID" json:"otc_reports"` // 一对多关系
}

type OTCReport struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	PatientID  uint       `gorm:"not null" json:"patientID"`
	Patient    Patient    `gorm:"foreignKey:PatientID" json:"patient"`
	ReportDate time.Time  `gorm:"not null" json:"reportDate"`
	Prediction Prediction `gorm:"foreignKey:OTCReportID" json:"prediction"` // 一对一关系
}

type Doctor struct {
	gorm.Model            // 自动包含 ID、CreatedAt、UpdatedAt 和 DeletedAt
	Name       string     `gorm:"size:100;not null" json:"name"`              // 设置列的大小并确保不为 null
	Email      string     `gorm:"size:100;not null;unique" json:"email"`      // 设置唯一的 Email
	Contact    string     `gorm:"size:20;not null" json:"contact"`            // 设置联系字段的大小并确保不为 null
	Address    string     `gorm:"size:255" json:"address"`                    // 地址字段，可以根据需要设置大小
	Profession string     `gorm:"size:100" json:"profession"`                 // 职业字段的大小
	Patients   []*Patient `gorm:"many2many:doctor_patients;" json:"patients"` // 关联 doctor_patients 连接表

}

// 指定 GORM 需要的表名
func (Doctor) TableName() string {
	return "doctors"
}

func (Patient) TableName() string {
	return "patients"
}
