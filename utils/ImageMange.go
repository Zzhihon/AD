package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
)

// SaveFile 将文件保存到 MinIO，并返回文件路径
func SaveFile(fileBytes []byte) (string, error) {
	if MinioClient == nil {
		return "", fmt.Errorf("MinIO client is not initialized")
	}

	// 生成唯一的文件名
	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), "report.jpg")

	// 上传文件到 MinIO
	bucketName := "ad-project" // MinIO 桶名称
	_, err := MinioClient.PutObject(context.Background(), bucketName, fileName, bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("无法上传文件到 MinIO: %v", err)
	}

	// 返回文件的访问路径
	filePath := fmt.Sprintf("http://183.6.97.121:9002/%s/%s", bucketName, fileName)
	return filePath, nil
}

// GetFile 从 MinIO 桶中获取文件
func GetFile(bucketName, objectName string) ([]byte, error) {
	if MinioClient == nil {
		return nil, fmt.Errorf("MinIO client is not initialized")
	}

	// 获取文件对象
	object, err := MinioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("无法获取文件: %v", err)
	}
	defer object.Close()

	// 读取文件内容
	fileBytes, err := io.ReadAll(object)
	if err != nil {
		return nil, fmt.Errorf("无法读取文件内容: %v", err)
	}

	return fileBytes, nil
}
