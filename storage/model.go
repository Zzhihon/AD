package storage

// Prediction 结构体，存储 AI 预测结果
type Prediction struct {
	ID          int64   `db:"id"`
	ImageID     string  `db:"image_id"`
	Probability float64 `db:"probability"`
	Advice      string  `db:"advice"`
	CreatedAt   string  `db:"created_at"`
}

type Doctor struct {
	DoctorID   string `db:"doctor_id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Contact    string `db:"contact"`
	Address    string `db:"address"`
	Profession string `db:"profession"`
}

type Patient struct {
	PatientID string `db:"patient_id"`
	Name      string `db:"name"`
	Gender    string `db:"gender"`
	Age       int    `db:"age"`
	Height    int    `db:"height"`
	Weight    int    `db:"weight"`
	BloodType string `db:"blood_type"`
}

type OTCReport struct {
	ReportID   string `db:"report_id"`
	PatientID  string `db:"patient_id"`
	ImageID    string `db:"image_id"`
	HostDoctor string `db:"host_doctor_id"`
}
