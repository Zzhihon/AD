package service

import (
	"AD/mq_producer"
	"AD/storage"
	"AD/utils"
	"log"
)

type PredicService struct {
	PredictionRepo *storage.PredictionRepository
}

func NewPredictService(predictionRepo *storage.PredictionRepository) *PredicService {
	return &PredicService{PredictionRepo: predictionRepo}
}

// UploadImage 处理文件上传，保存文件并推送路径到 RabbitMQ
func (s *PredicService) UploadImage(fileBytes []byte) error {
	// 保存图片并获取路径
	filePath, err := utils.SaveFile(fileBytes)
	if err != nil {
		return err
	}
	log.Println(filePath)
	filePath = "/var/www/docker_prod/ad/upload/normal.jpg"

	// 推送路径到 RabbitMQ
	err = mq_producer.PublishTask(filePath)
	if err != nil {
		return err
	}

	return nil
}

// ProcessPrediction 处理预测结果并保存到数据库
func (s *PredicService) ProcessPrediction(filePath string) error {
	// 调用远程服务器 API 获取预测结果
	predictionResponse, err := utils.CallAIPrediction(filePath)
	if err != nil {
		log.Println(predictionResponse)
		return err
	}

	// 将预测结果保存到数据库
	prediction := &storage.Prediction{
		OTCReportID: 1, // 假设 OTCReportID 为 1
		Probability: predictionResponse.Prediction,
		Advice:      predictionResponse.Advice,
		LR:          predictionResponse.LR,
		SVM:         predictionResponse.SVM,
		DT:          predictionResponse.DT,
		Final:       predictionResponse.Final,
	}
	return s.PredictionRepo.SavePrediction(prediction)
}
