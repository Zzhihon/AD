package storage

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// DB 全局数据库连接池
var DB *sqlx.DB

// InitDB 初始化数据库连接
func InitDB(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 设置连接池参数
	DB.SetMaxOpenConns(50) // 最大打开连接数
	DB.SetMaxIdleConns(10) // 最大空闲连接数
}

// SavePrediction 保存预测结果
func SavePrediction(prediction Prediction) error {
	query := `INSERT INTO predictions (image_id, probability, created_at) VALUES (:image_id, :probability, NOW())`
	_, err := DB.NamedExec(query, prediction)
	return err
}

// GetPredictionByImageID 通过 image_id 获取预测结果
func GetPredictionByImageID(imageID string) (*Prediction, error) {
	var prediction Prediction
	query := `SELECT id, image_id, probability, created_at FROM predictions WHERE image_id = ?`
	err := DB.Get(&prediction, query, imageID)
	if err != nil {
		return nil, err
	}
	return &prediction, nil
}
