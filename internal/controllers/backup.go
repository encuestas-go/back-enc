package controllers

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	zipFileName := currentDate + "_backup.sql"

	command := fmt.Sprintf("mysqldump -h 127.0.0.1 -P 3306 -u root -p ENCUESTA > backups/%s", zipFileName)

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

func (b *BackupController) Restore(c echo.Context) error {
	backupFile := c.QueryParam("file")
	if backupFile == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "file query param is required"})
	}

	backupFilePath := filepath.Join("backups", backupFile)
	if _, err := os.Stat(backupFilePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "backup file not found"})
	}

	// Actualizar docker-compose.yml
	replaceString := fmt.Sprintf(`    SQL_FILE: "%s"`, backupFile)
	err := replaceInFile("docker-compose.yml", `    SQL_FILE: ".*"`, replaceString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Detener y reiniciar el contenedor
	err = restartDockerCompose()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Backup file successfully restored",
	})
}

func replaceInFile(filename, old, new string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	output := strings.Replace(string(input), old, new, 1)

	err = os.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func restartDockerCompose() error {
	cmd := exec.Command("sh", "-c", "docker-compose down && docker-compose up -d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
