package handler

import (
	"AD/dto"
	"AD/service"
	"AD/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// PatientHandler 负责 HTTP 请求
type PatientHandler struct {
	PatientService *service.PatientService
}

// NewPatientHandler 创建新的 PatientHandler 实例
func NewPatientHandler(patientService *service.PatientService) *PatientHandler {
	return &PatientHandler{PatientService: patientService}
}

func (h *PatientHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var request dto.CreatePatientRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}
	err = h.PatientService.CreatePatient(&request)
	if err != nil {

		http.Error(w, "Error creating patient"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}

// GetPatientByID 处理根据 PatientID 获取医生的请求
func (h *PatientHandler) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID := vars["patient_id"]

	patient, err := h.PatientService.GetPatientByID(patientID)
	if err != nil || patient == nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}

// UpdatePatient 处理更新医生信息的请求
func (h *PatientHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	var patient storage.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.PatientService.UpdatePatient(&patient)
	if err != nil {
		http.Error(w, "Error updating patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}
