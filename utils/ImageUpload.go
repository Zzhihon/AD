package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// SaveFile 将文件保存到本地，并返回文件路径
func SaveFile(fileBytes []byte) (string, error) {
	// 确保上传目录存在
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.Mkdir(uploadDir, 0755) // 创建目录，权限为 0755
		if err != nil {
			return "", fmt.Errorf("无法创建上传目录: %v", err)
		}
	}

	// 生成唯一的文件名
	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), "image.jpg") // 使用时间戳作为文件名前缀
	filePath := filepath.Join(uploadDir, fileName)

	// 将文件保存到本地
	err := ioutil.WriteFile(filePath, fileBytes, 0644) // 权限为 0644
	if err != nil {
		return "", fmt.Errorf("无法保存文件: %v", err)
	}

	// 返回文件路径
	return filePath, nil
}
