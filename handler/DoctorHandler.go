package handler

import (
	"AD/service"
	"AD/storage"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DoctorHandler 负责 HTTP 请求
type DoctorHandler struct {
	DoctorService *service.DoctorService
}

// NewDoctorHandler 创建新的 DoctorHandler 实例
func NewDoctorHandler(doctorService *service.DoctorService) *DoctorHandler {
	return &DoctorHandler{DoctorService: doctorService}
}

func (h *DoctorHandler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor storage.Doctor
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = h.DoctorService.CreateDoctor(&doctor)
	if err != nil {
		http.Error(w, "Error creating doctor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

// GetDoctorByID 处理根据 DoctorID 获取医生的请求
func (h *DoctorHandler) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorID := vars["doctor_id"]

	doctor, err := h.DoctorService.GetDoctorByID(doctorID)
	if err != nil || doctor == nil {
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doctor)
}

// UpdateDoctor 处理更新医生信息的请求
func (h *DoctorHandler) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor storage.Doctor
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.DoctorService.UpdateDoctor(&doctor)
	if err != nil {
		http.Error(w, "Error updating doctor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doctor)
}

// GetPatients 获取指定医生的所有病人
func (h *DoctorHandler) GetPatients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorIDStr := vars["doctor_id"]
	if doctorIDStr == "" {
		http.Error(w, "doctor_id is required", http.StatusBadRequest)
		return
	}

	doctorID, err := strconv.Atoi(doctorIDStr)
	if err != nil || doctorID <= 0 {
		http.Error(w, "Invalid doctor_id", http.StatusBadRequest)
		return
	}

	// 调用 service 层方法获取病人数据
	patients, err := h.DoctorService.GetPatientsByDoctorID(uint(doctorID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching patients: %v", err), http.StatusInternalServerError)
		return
	}

	// 设置响应的 Content-Type 为 application/json
	w.Header().Set("Content-Type", "application/json")

	// 返回病人数据
	if err := json.NewEncoder(w).Encode(patients); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
