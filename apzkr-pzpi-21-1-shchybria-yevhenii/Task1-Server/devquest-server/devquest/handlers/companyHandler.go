package handlers

import (
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/usecases"
	"devquest-server/devquest/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CompanyHttpHandler struct {
	companyUsecase usecases.CompanyUsecase
}

func NewCompanyHttpHandler(companyUsecase usecases.CompanyUsecase) *CompanyHttpHandler {
	return &CompanyHttpHandler{companyUsecase: companyUsecase}
}

func (c *CompanyHttpHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := c.companyUsecase.FindAllCompanies()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, companies)
}

func (c *CompanyHttpHandler) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	company, err := c.companyUsecase.FindCompanyByID(companyID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, company)
}

func (c *CompanyHttpHandler) AddCompany (w http.ResponseWriter, r *http.Request) {
	var insertCompanyDTO models.InsertCompanyDTO

	err := utils.ReadJSON(w, r, &insertCompanyDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	newCompany, err := c.companyUsecase.CreateNewCompany(&insertCompanyDTO)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, newCompany)
}

func (c *CompanyHttpHandler) UpdateCompany (w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var payload models.UpdateCompanyDTO
	err = utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = c.companyUsecase.UpdateCompany(companyID, payload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "company successfully updated",
	}

	utils.WriteJSON(w, http.StatusAccepted, res)
}

func (c *CompanyHttpHandler) DeleteCompany (w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = c.companyUsecase.DeleteCompany(companyID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "company successfully deleted",
	}

	utils.WriteJSON(w, http.StatusAccepted, res)
}