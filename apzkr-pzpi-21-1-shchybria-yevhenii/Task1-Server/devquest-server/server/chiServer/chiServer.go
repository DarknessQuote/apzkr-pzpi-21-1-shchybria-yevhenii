package chiServer

import (
	"devquest-server/config"
	"devquest-server/devquest/handlers"
	"devquest-server/devquest/infrastructure"
	"devquest-server/devquest/infrastructure/postgres"
	"devquest-server/devquest/usecases"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type chiServer struct {
	Config *config.Config
	Database *infrastructure.Database
	AuthSettings *infrastructure.Auth
}

var (
	once sync.Once;
	serverInstance *chiServer
	userHttpHandler *handlers.UserHttpHandler
	adminHttpHander *handlers.AdminHttpHandler
	companyHttpHandler *handlers.CompanyHttpHandler
	projectHttpHandler *handlers.ProjectHttpHandler
	achievementHttpHandler *handlers.AchievementHttpHandler
	taskHttpHandler *handlers.TaskHttpHandler
	measurementHttpHander *handlers.MeasurementHttpHandler
)

func NewChiServer(conf *config.Config, db *infrastructure.Database, auth *infrastructure.Auth) *chiServer {
	once.Do( func() {
		serverInstance = &chiServer {
		Config: conf,
		Database: db,
		AuthSettings: auth,
	}
	})
	return serverInstance
}

func GetChiServer() *chiServer {
	return serverInstance
}

func (s *chiServer) Start() {
	initializeHttpHandlers()

	port := s.Config.Server.Port
	serverUrl := fmt.Sprintf(":%d", port)	
	router := getRoutes()
	
	log.Printf("Starting application on port %d", port)
	if err := http.ListenAndServe(serverUrl, router); err != nil {
		log.Fatal(err)
	}
}

func initializeHttpHandlers() {
	userRepository := postgres.NewUserPostgresRepo(*serverInstance.Database)
	companyRepository := postgres.NewCompanyPostgresRepo(*serverInstance.Database)
	projectRepository := postgres.NewProjectPostgresRepo(*serverInstance.Database)
	achievementRepository := postgres.NewAchievementPostgresRepo(*serverInstance.Database)
	taskRepository := postgres.NewTaskPostgresRepo(*serverInstance.Database)
	measurementRepository := postgres.NewMeasurementPostgresRepo(*serverInstance.Database)

	userUsecase := usecases.NewUserUsecase(userRepository, companyRepository)
	companyUsecase := usecases.NewCompanyUsecase(companyRepository)
	projectUsecase := usecases.NewProjectUsecase(projectRepository, userRepository, companyRepository)
	achievementUsecase := usecases.NewAchievementUsecase(achievementRepository, projectRepository, userRepository)
	taskUsecase := usecases.NewTaskUsecase(taskRepository, projectRepository, userRepository)
	measurementUsecase := usecases.NewMeasurementUsecase(measurementRepository, userRepository)

	userHttpHandler = handlers.NewUserHttpHandler(*userUsecase)
	adminHttpHander = handlers.NewAdminHttpHandler(*serverInstance.Database, serverInstance.Config)
	companyHttpHandler = handlers.NewCompanyHttpHandler(*companyUsecase)
	projectHttpHandler = handlers.NewProjectHttpHandler(*projectUsecase)
	achievementHttpHandler = handlers.NewAchievementHttpHandler(*achievementUsecase)
	taskHttpHandler = handlers.NewTaskHttpHandler(*taskUsecase)
	measurementHttpHander = handlers.NewMeasurementHttpHandler(*measurementUsecase)
}