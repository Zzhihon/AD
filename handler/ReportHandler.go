package handler

import (
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

func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request) {
	var report storage.OTCReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println()
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
