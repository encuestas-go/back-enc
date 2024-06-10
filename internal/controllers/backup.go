package controllers

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type BackupController struct {
	fsys fs.FS
}

func InitBackupController() *BackupController {
	return &BackupController{
		fsys: os.DirFS("backups"),
	}
}

type BackupFiles struct {
	Name string `json:"name"`
}

func (b *BackupController) Create(c echo.Context) error {
	//date := "2024-06-09"
	return nil
}

func (b *BackupController) Get(c echo.Context) error {
	//backupsDir := "../backups"
	backupsDir := filepath.Join("internal", "backups")

	dir, err := os.ReadDir(backupsDir)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var backupFiles []string
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		backupFiles = append(backupFiles, entry.Name())
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Backup files successfully retrieved",
		Data:       backupFiles,
	})
}
