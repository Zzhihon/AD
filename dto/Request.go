package dto

type CreatePatientRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	DoctorID uint   `json:"doctor_id"` // 传递医生的ID
}
