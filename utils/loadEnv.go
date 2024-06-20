package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

/*
This function loads the environment variables from .env file
 1. Get path of current working dir
 2. Trim "specs" from pwd to get the root dir
 3. Get path of .env by joining root dir and .env file name
 4. Load environment variables from .env
*/
func LoadEnv() error {
	pwd, err := os.Getwd()
	rootDir := strings.TrimSuffix(pwd, "specs")
	envPath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("error loading environment variables: %w", err)
	}
	return nil
}
