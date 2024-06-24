package handlers

import (
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/usecases"
	"devquest-server/devquest/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AchievementHttpHandler struct {
	achievementUsecase usecases.AchievementUsecase
}

func NewAchievementHttpHandler(aUsecase usecases.AchievementUsecase) *AchievementHttpHandler {
	return &AchievementHttpHandler{achievementUsecase: aUsecase}
}

func (a *AchievementHttpHandler) GetProjectAchievements(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "project_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	projectsAchievements, err := a.achievementUsecase.GetProjectAchievements(projectID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, projectsAchievements)
}

func (a *AchievementHttpHandler) GetDeveloperAchievements(w http.ResponseWriter, r *http.Request) {
	developerID, err := uuid.Parse(chi.URLParam(r, "developer_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developerAchievements, err := a.achievementUsecase.GetDeveloperAchievements(developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, developerAchievements)
}

func (a *AchievementHttpHandler) AddAchievementToProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(chi.URLParam(r, "project_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var createAchievementDTO models.CreateAchievementDTO
	createAchievementDTO.ProjectID = projectID
	err = utils.ReadJSON(w, r, &createAchievementDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = a.achievementUsecase.AddAchievementToProject(createAchievementDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "achievement successfully created",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (a *AchievementHttpHandler) UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	achievementID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var updateAchievementDTO models.UpdateAchievementDTO
	err = utils.ReadJSON(w, r, &updateAchievementDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = a.achievementUsecase.UpdateAchievement(achievementID, updateAchievementDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "achievement successfully updated",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (a *AchievementHttpHandler) DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	achievementID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = a.achievementUsecase.DeleteAchievement(achievementID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "achievement successfully deleted",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (a *AchievementHttpHandler) GiveAchievementToDeveloper(w http.ResponseWriter, r *http.Request) {
	projectID, err := uuid.Parse(r.URL.Query().Get("projectID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	achievementID, err := uuid.Parse(r.URL.Query().Get("achievementID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	developerID, err := uuid.Parse(r.URL.Query().Get("developerID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = a.achievementUsecase.GiveAchievementToDeveloper(projectID, achievementID, developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "achievement given to developer",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}