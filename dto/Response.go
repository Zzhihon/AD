package dto

type PredictionResponse struct {
	LR         float64 `json:"lr"`
	SVM        float64 `json:"svm"`
	DT         float64 `json:"dt"`
	Final      float64 `json:"final"`
	Prediction string  `json:"prediction"`
	Advice     string  `json:"advice"`
}
