package handler

import (
	"AD/service"
	"AD/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type PredictHandler struct {
	PredictService *service.PredicService
}

func NewPredictHandler(uploadService *service.PredicService) *PredictHandler {
	return &PredictHandler{PredictService: uploadService}
}

func (h *PredictHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
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

	err = h.PredictService.UploadImage(fileBytes)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload successful"))
}

func (h *PredictHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中获取文件名
	vars := mux.Vars(r)
	fileName := vars["fileName"]

	if fileName == "" {
		http.Error(w, "文件名不能为空", http.StatusBadRequest)
		return
	}

	// 从 MinIO 桶中获取文件
	bucketName := "ad-project" // MinIO 桶名称
	fileBytes, err := utils.GetFile(bucketName, fileName)
	if err != nil {
		log.Printf("无法获取文件: %v", err)
		http.Error(w, "无法获取文件", http.StatusInternalServerError)
		return
	}

	// 返回文件内容
	w.Header().Set("Content-Type", "image/jpeg") // 根据文件类型设置 MIME 类型
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}
