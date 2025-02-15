package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 处理文件上传
func ImageUpload(w http.ResponseWriter, r *http.Request) {
	// 解析 multipart/form-data 请求
	err := r.ParseMultipartForm(10 << 20) // 限制上传文件大小为 10MB
	if err != nil {
		http.Error(w, "无法解析表单", http.StatusBadRequest)
		return
	}

	// 获取上传的文件
	file, handler, err := r.FormFile("file") // "file" 是前端上传文件的字段名
	if err != nil {
		http.Error(w, "无法获取文件"+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 打印文件信息
	fmt.Printf("上传的文件名: %s\n", handler.Filename)
	fmt.Printf("文件大小: %d bytes\n", handler.Size)
	fmt.Printf("MIME 类型: %s\n", handler.Header.Get("Content-Type"))

	// 创建目标文件
	dst, err := os.Create("./uploads/" + handler.Filename) // 保存到 uploads 目录
	if err != nil {
		http.Error(w, "无法创建文件", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "无法保存文件", http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "文件上传成功: %s\n", handler.Filename)
}
