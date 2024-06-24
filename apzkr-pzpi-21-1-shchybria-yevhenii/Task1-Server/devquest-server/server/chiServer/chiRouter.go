package chiServer

import (
	"devquest-server/server/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)
func getRoutes() http.Handler {
	mux := chi.NewMux()
	authSettings := GetChiServer().AuthSettings

	mux.Use(middleware.EnableCORS)

	mux.Group(func(r chi.Router) {
		r.Post("/auth/login", userHttpHandler.Login(authSettings))
		r.Post("/auth/register", userHttpHandler.Register(authSettings))
		r.Post("/auth/refresh", userHttpHandler.RefreshToken(authSettings))
		r.Delete("/auth/logout", userHttpHandler.Logout(authSettings))
		r.Get("/auth/roles", userHttpHandler.GetRolesForRegistration)
		r.Get("/user/{id}", userHttpHandler.GetUserByID)
		r.Get("/role/{id}", userHttpHandler.GetRoleByID)

		r.Get("/companies", companyHttpHandler.GetAllCompanies)
		r.Get("/companies/{id}", companyHttpHandler.GetCompanyByID)

		r.Put("/measure/add-owner", measurementHttpHander.AddOwnerToDevice)
		r.Post("/measure", measurementHttpHander.AddMeasurementResult)
		r.Get("/measure/{developer_id}", measurementHttpHander.GetLatestMeasurementsForDeveloper)
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.RolesRequired(*authSettings, "Admin"))

		r.Post("/companies", companyHttpHandler.AddCompany)
		r.Put("/companies/{id}", companyHttpHandler.UpdateCompany)
		r.Delete("/companies/{id}", companyHttpHandler.DeleteCompany)
		r.HandleFunc("/admin/data-backup", adminHttpHander.GetDatabaseBackup)
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.RolesRequired(*authSettings, "Manager,Developer"))

		r.Get("/projects/{id}", projectHttpHandler.GetProjectByID)
		r.Get("/projects/developers/{project_id}", projectHttpHandler.GetProjectDevelopers)

		r.Get("/tasks/{project_id}", taskHttpHandler.GetProjectTasks)
		r.Get("/tasks/categories/{id}", taskHttpHandler.GetTaskCategoryByID)
		r.Get("/tasks/status/{id}", taskHttpHandler.GetTaskStatusByID)

		r.Get("/achievements/project/{project_id}", achievementHttpHandler.GetProjectAchievements)
		r.Get("/achievements/developer/{developer_id}", achievementHttpHandler.GetDeveloperAchievements)
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.RolesRequired(*authSettings, "Manager"))

		r.Get("/projects/available-developers/{manager_id}", userHttpHandler.GetDevelopersForManager)

		r.Get("/projects/manager/{manager_id}", projectHttpHandler.GetProjectsOfManager)
		r.Put("/projects/{id}", projectHttpHandler.UpdateProject)
		r.Post("/projects", projectHttpHandler.AddProject)
		r.Delete("/projects/{id}", projectHttpHandler.DeleteProject)
		r.Post("/projects/developers", projectHttpHandler.AddDeveloperToProject)
		r.Delete("/projects/developers", projectHttpHandler.RemoveDeveloperFromProject)

		r.Post("/tasks/{project_id}", taskHttpHandler.CreateNewTask)
		r.Put("/tasks/{id}", taskHttpHandler.UpdateTask)
		r.Delete("/tasks/{id}", taskHttpHandler.DeleteTask)
		r.Get("/tasks/categories", taskHttpHandler.GetTaskCategories)

		r.Post("/achievements/{project_id}", achievementHttpHandler.AddAchievementToProject)
		r.Put("/achievements/{id}", achievementHttpHandler.UpdateAchievement)
		r.Delete("/achievements/{id}", achievementHttpHandler.DeleteAchievement)
		r.Post("/achievements/give", achievementHttpHandler.GiveAchievementToDeveloper)
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.RolesRequired(*authSettings, "Developer"))

		r.Get("/projects/developer/{developer_id}", projectHttpHandler.GetProjectsOfDeveloper)
		r.Put("/tasks/accept", taskHttpHandler.AcceptTask)
		r.Put("/tasks/complete", taskHttpHandler.CompleteTask)
	})

	return mux
}