package main

import (
	"AD/handler"
	"AD/mq_consumer"
	"AD/service"
	"AD/storage"
	"AD/utils"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// 创建 uploads 目录
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		fmt.Println("无法创建 uploads 目录:", err)
		return
	}

	//连接minio
	minioClient()

	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "*"}, // 允许的前端源
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handlers := c.Handler(router)
	// 读取配置
	//config.LoadConfig()

	// 初始化数据库
	db := dbclient()

	// 创建 Repository
	doctorRepo := storage.NewDoctorRepository(db)
	patientRepo := storage.NewPatientRepository(db)
	reportRepo := storage.NewReportRepository(db)
	predictionRepo := storage.NewPredictionRepository(db)

	// 创建 Service，将 Repository 注入其中
	doctorService := service.NewDoctorService(doctorRepo)
	patientService := service.NewPatientService(patientRepo)
	reportService := service.NewReportService(reportRepo)
	predictionService := service.NewPredictService(predictionRepo)

	go mq_consumer.StartConsumer(predictionService)

	// 创建 Handler，将 Service 注入其中
	doctorHandler := handler.NewDoctorHandler(doctorService)
	patientHandler := handler.NewPatientHandler(patientService)
	reportHandler := handler.NewReportHandler(reportService)
	predictHandler := handler.NewPredictHandler(predictionService)

	// 启动 HTTP 服务器

	apipath := os.Getenv("API_PATH")

	//router.HandleFunc("/upload/", service.UploadHandler).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/ws/", service.WebSocketHandler).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/AddDoctor/", doctorHandler.CreateDoctor).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/GetDoctor/{doctor_id:[0-9]+}/", doctorHandler.GetDoctorByID).Methods(http.MethodGet)
	router.HandleFunc(apipath+"/UpdateDoctor/", doctorHandler.UpdateDoctor).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/GetPatients/{doctor_id:[0-9]+}/", doctorHandler.GetPatients).Methods(http.MethodGet)

	router.HandleFunc(apipath+"/AddPatient/", patientHandler.CreatePatient).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/GetPatient/{patient_id:[0-9]+}/", patientHandler.GetPatientByID).Methods(http.MethodGet)
	router.HandleFunc(apipath+"/UpdatePatient/", patientHandler.UpdatePatient).Methods(http.MethodPost)

	router.HandleFunc(apipath+"/AddReport/", reportHandler.CreateReport).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/GetReport/{report_id:[0-9]+}/", reportHandler.GetReportByID).Methods(http.MethodGet)
	router.HandleFunc(apipath+"/UpdateReport/", reportHandler.UpdateReport).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/FindReportsByID/{patient_id:[0-9]+}/", reportHandler.FindByPatientID).Methods(http.MethodGet)
	router.HandleFunc(apipath+"/Search/", reportHandler.Search).Methods(http.MethodPost)

	router.HandleFunc(apipath+"/UploadImage/", predictHandler.UploadImage).Methods(http.MethodPost)
	router.HandleFunc(apipath+"/GetImage/{fileName:[A-Za-z0-9!@#-.]*}/", predictHandler.GetImage).Methods(http.MethodGet)

	log.Println("Server started on :9088")
	log.Fatal(http.ListenAndServe(":9088", handlers))
}

func dbclient() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	log.Println(dbName)
	// PostgreSQL 连接信息
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPasswd, dbName, dbAddr, dbPort)

	// 连接 PostgreSQL 数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库
	err = db.AutoMigrate(&storage.Doctor{}, &storage.OTCReport{}, &storage.Patient{}, &storage.Prediction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}

func minioClient() {
	minioAdress := os.Getenv("MINIO_ADDR")
	minioAccess := os.Getenv("MINIO_ACCESS")
	minioSecret := os.Getenv("MINIO_SECRET")

	utils.InitMinioClient(minioAdress, minioAccess, minioSecret, false)

}
