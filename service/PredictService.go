package service

import (
	"AD/mq"
	"io/ioutil"
	"net/http"
)

// 处理上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "File read error", http.StatusInternalServerError)
		return
	}
	// 推送到 RabbitMQ
	err = mq.PublishTask(imageData)
	if err != nil {
		http.Error(w, "Failed to enqueue task", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Upload successful"))
}
