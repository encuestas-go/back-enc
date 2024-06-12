package controllers

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
	currentDate := time.Now().Format("2006-01-02")

	zipFileName := currentDate + "_database.zip"

	command := "zip -r " + zipFileName + " mysql-data && mv " + zipFileName + " backups/"

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    "Backup file successfully created",
	})
}

func (b *BackupController) Get(c echo.Context) error {
	backupsDir := filepath.Join("backups")

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
