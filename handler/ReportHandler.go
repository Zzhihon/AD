package handler

import (
	"AD/dto"
	"AD/service"
	"AD/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// ReportHandler 负责 HTTP 请求
type ReportHandler struct {
	ReportService *service.ReportService
}

// NewReportHandler 创建新的 ReportHandler 实例
func NewReportHandler(reportService *service.ReportService) *ReportHandler {
	return &ReportHandler{ReportService: reportService}
}

func (h *ReportHandler) FindByPatientID(w http.ResponseWriter, r *http.Request) {
	// 从 URL 参数中获取 patientID
	vars := mux.Vars(r)
	patientID := vars["patient_id"]

	// 调用 Service 层
	reports, err := h.ReportService.FindByPatientID(patientID)
	if err != nil {
		log.Println("Failed to find OTC reports:", err)
		http.Error(w, "Failed to find OTC reports", http.StatusInternalServerError)
		return
	}

	// 返回结果
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reports); err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ReportHandler) Search(w http.ResponseWriter, r *http.Request) {
	var req dto.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 调用 Service 层
	reports, err := h.ReportService.Search(req)
	if err != nil {
		log.Println("Failed to search OTC reports:", err)
		http.Error(w, "Failed to search OTC reports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reports); err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	} else {
		log.Println(reports)
		log.Println("Report search succeeded")
	}
}

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var report dto.OTCFormRequest
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.ReportService.CreateReport(&report)
	if err != nil {
		http.Error(w, "Error creating report", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(report)
}

// GetReportByID 处理根据 ReportID 获取医生的请求
func (h *ReportHandler) GetReportByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reportID := vars["report_id"]

	report, err := h.ReportService.GetReportByID(reportID)
	if err != nil || report == nil {
		http.Error(w, "Report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}

// UpdateReport 处理更新医生信息的请求
func (h *ReportHandler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	var report storage.OTCReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.ReportService.UpdateReport(&report)
	if err != nil {
		http.Error(w, "Error updating report", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}
