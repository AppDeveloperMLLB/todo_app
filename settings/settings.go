package settings

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Initialize() {
	loadEnv()
}

func loadEnv() {
	rootPath := ProjectRoot()
	err := godotenv.Load(filepath.Join(rootPath, ".env"))

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}

func ProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}
