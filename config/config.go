package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func GetProjectRoot() string {
	// Adjust this to match your project structure
	// Assumes this file is always located within the root directory
	root, _ := os.Getwd()
	return root
}

type FilePaths struct {
	ProductsFilePath  string
	CustomersFilePath string
}

// NewFilePaths initializes file paths based on project root.
func NewFilePaths() *FilePaths {
	root := GetProjectRoot()
	return &FilePaths{
		ProductsFilePath:  filepath.Join(root, "data", "products.json"),
		CustomersFilePath: filepath.Join(root, "data", "customers.json"),
	}
}
