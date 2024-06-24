package handlers

import (
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/usecases"
	"devquest-server/devquest/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ProjectHttpHandler struct {
	projectUsecase usecases.ProjectUsecase
}

func NewProjectHttpHandler(pUsecase usecases.ProjectUsecase) *ProjectHttpHandler {
	return &ProjectHttpHandler{projectUsecase: pUsecase}
}

func (p *ProjectHttpHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	project, err := p.projectUsecase.GetProjectByID(projectID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, project)
}

func (p *ProjectHttpHandler) GetProjectsOfManager(w http.ResponseWriter, r *http.Request) {
	managerID, err := uuid.Parse(chi.URLParam(r, "manager_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	projects, err := p.projectUsecase.GetManagerProjects(managerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, projects)
}

func (p *ProjectHttpHandler) GetProjectsOfDeveloper(w http.ResponseWriter, r *http.Request) {
	developerID, err := uuid.Parse(chi.URLParam(r, "developer_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	projects, err := p.projectUsecase.GetDeveloperProjects(developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, projects)
}

func (p *ProjectHttpHandler) GetProjectDevelopers(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "project_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developers, err := p.projectUsecase.GetProjectDevelopers(projectID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, developers)
}

func (p *ProjectHttpHandler) AddProject(w http.ResponseWriter, r *http.Request) {		
	var createProjectDTO models.CreateProjectDTO
	err := utils.ReadJSON(w, r, &createProjectDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = p.projectUsecase.CreateNewProject(createProjectDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "project successfully created",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (p *ProjectHttpHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var updateProjectDTO models.UpdateProjectDTO
	err = utils.ReadJSON(w, r, &updateProjectDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = p.projectUsecase.UpdateProject(projectID, updateProjectDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "project successfully updated",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (p *ProjectHttpHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = p.projectUsecase.DeleteProject(projectID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "project successfully deleted",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (p *ProjectHttpHandler) AddDeveloperToProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(r.URL.Query().Get("projectID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developerID, err := uuid.Parse(r.URL.Query().Get("developerID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = p.projectUsecase.AddDeveloperToProject(projectID, developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "developer added to project",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (p *ProjectHttpHandler) RemoveDeveloperFromProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(r.URL.Query().Get("projectID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developerID, err := uuid.Parse(r.URL.Query().Get("developerID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = p.projectUsecase.RemoveDeveloperFromProject(projectID, developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "developer removed from project",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}
