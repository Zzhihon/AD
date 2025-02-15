package main

import (
	"AD/handler"
	"AD/mq"
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
	// 启动 RabbitMQ 消费者
	go mq.StartConsumer()

	// 创建 Repository
	doctorRepo := storage.NewDoctorRepository(db)
	patientRepo := storage.NewPatientRepository(db)
	reportRepo := storage.NewReportRepository(db)

	// 创建 Service，将 Repository 注入其中
	doctorService := service.NewDoctorService(doctorRepo)
	patientService := service.NewPatientService(patientRepo)
	reportService := service.NewReportService(reportRepo)

	// 创建 Handler，将 Service 注入其中
	doctorHandler := handler.NewDoctorHandler(doctorService)
	patientHandler := handler.NewPatientHandler(patientService)
	reportHandler := handler.NewReportHandler(reportService)

	// 启动 HTTP 服务器
	//router.HandleFunc("/upload/", service.UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/ws/", service.WebSocketHandler).Methods(http.MethodPost)
	router.HandleFunc("/AddDoctor/", doctorHandler.CreateDoctor).Methods(http.MethodPost)
	router.HandleFunc("/GetDoctor/{doctor_id:[0-9]+}/", doctorHandler.GetDoctorByID).Methods(http.MethodGet)
	router.HandleFunc("/UpdateDoctor/", doctorHandler.UpdateDoctor).Methods(http.MethodPost)
	router.HandleFunc("/GetPatients/{doctor_id:[0-9]+}/", doctorHandler.GetPatients).Methods(http.MethodGet)

	router.HandleFunc("/AddPatient/", patientHandler.CreatePatient).Methods(http.MethodPost)
	router.HandleFunc("/GetPatient/{patient_id:[0-9]+}/", patientHandler.GetPatientByID).Methods(http.MethodGet)
	router.HandleFunc("/UpdatePatient/", patientHandler.UpdatePatient).Methods(http.MethodPost)

	router.HandleFunc("/AddReport/", reportHandler.CreateReport).Methods(http.MethodPost)
	router.HandleFunc("/GetReport/{report_id:[0-9]+}/", reportHandler.GetReportByID).Methods(http.MethodGet)
	router.HandleFunc("/UpdateReport/", reportHandler.UpdateReport).Methods(http.MethodPost)

	router.HandleFunc("/ImageUpload/", utils.ImageUpload).Methods(http.MethodPost)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers))
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
