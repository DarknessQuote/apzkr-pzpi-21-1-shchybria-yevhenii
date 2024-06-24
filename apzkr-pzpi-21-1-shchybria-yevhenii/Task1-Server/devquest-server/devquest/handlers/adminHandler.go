package handlers

import (
	"devquest-server/config"
	"devquest-server/devquest/infrastructure"
	"devquest-server/devquest/utils"
	"fmt"
	"net/http"
	"os"
)

type AdminHttpHandler struct{
	DB infrastructure.Database
	Config *config.Config
}

func NewAdminHttpHandler(db infrastructure.Database, conf *config.Config) *AdminHttpHandler {
	return &AdminHttpHandler{DB: db, Config: conf}
}

func (a *AdminHttpHandler) GetDatabaseBackup(w http.ResponseWriter, r *http.Request) {
	dumpResult := a.DB.CreateBackup(a.Config)
	if dumpResult.Error != nil {
		utils.ErrorJSON(w, dumpResult.Error.Err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		utils.ErrorJSON(w, dumpResult.Error.Err)
		return
	}

	backupFilePath := fmt.Sprintf("%s\\backups\\%s", currentDir, dumpResult.File)
	w.Header().Add("Content-Type", "application/x-tar")
	http.ServeFile(w, r, backupFilePath)
}