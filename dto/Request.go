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

type SearchRequest struct {
	PatientName      string `json:"patient_name"`      // 患者名字（模糊搜索）
	Gender           string `json:"gender"`            // 性别（男或女）
	AgeRange         string `json:"age_range"`         // 年龄区间（56-60，61-65，66-70）
	OTCImageStatus   *int   `json:"otc_image_status"`  // OTC图像状态（0: 未上传, 1: 已上传）
	PredictionStatus *int   `json:"prediction_status"` // 报告状态（0: 未开始, 1: 生成中, 2: 生成异常）
}
