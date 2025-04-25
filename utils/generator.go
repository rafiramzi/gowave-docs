package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
func Generate(folderName, fileName string) {
	var content string

	// Generate based on folder
	switch strings.ToLower(folderName) {
	case "controllers":
		content = `
package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ` + toPascalCase(fileName) + `Controller struct {
	DB *gorm.DB
}

func (pc *` + toPascalCase(fileName) + `Controller) GetHello(c echo.Context) error {

	var hello_world = "Hello World!"

	return c.String(http.StatusOK, hello_world)
}
`

	case "models":
		content = `
package models

type ` + toPascalCase(fileName) + ` struct {
	
}
`

	case "middlewares":
		content = 
`package middlewares

import (
	"github.com/labstack/echo/v4"
)

func ` + toPascalCase(fileName) + `echo.MidddlewareFunc {
	return func(c echo.Context) error {
		// Do something before
		err := next(c)
		// Do something after
		return err
	}
}
`

	default:
		fmt.Println("❌ Invalid folder name. Please use: controllers, models, or middlewares.")
		return
	}

	// Append Controller suffix if needed
	if strings.ToLower(folderName) == "controllers" && !strings.HasSuffix(fileName, "Controller") {
		fileName += "Controller"
	}

	// Create folder
	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	fmt.Println("📁 Folder ready:", folderName)

	// Create file
	filePath := filepath.Join(folderName, fileName+".go")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write content
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("✅ "+folderName+" created:", filePath)
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}