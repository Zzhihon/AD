package dto

type CreatePatientRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	DoctorID uint   `json:"doctor_id"` // 传递医生的ID
}

type OTCFormRequest struct {
	DoctorName  string `json:"doctor_name"`
	PatientName string `json:"patient_name"`
	PatientID   string `json:"patient_id"`
}
