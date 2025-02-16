package handler

import (
	"AD/service"
	"io/ioutil"
	"log"
	"net/http"
)

type UploadHandler struct {
	UploadService *service.UploadService
}

func NewUploadHandler(uploadService *service.UploadService) *UploadHandler {
	return &UploadHandler{UploadService: uploadService}
}

func (h *UploadHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// 解析 multipart/form-data 请求
	err := r.ParseMultipartForm(10 << 20) // 限制上传文件大小为 10MB
	if err != nil {
		log.Println(err)
		http.Error(w, "无法解析表单", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "File read error", http.StatusInternalServerError)
		return
	}

	err = h.UploadService.UploadImage(fileBytes)
	if err != nil {
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload successful"))
}
