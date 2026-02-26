// internal/service/file_generator.go
package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"otus-go/internal/model"
	"otus-go/internal/repository"
)

type FileGenerator struct {
	storage *repository.Storage
	rootDir string
}

func NewFileGenerator(storage *repository.Storage, rootDir string) *FileGenerator {
	return &FileGenerator{
		storage: storage,
		rootDir: rootDir,
	}
}

// StartGenerating сканирует директорию и периодически добавляет файлы в хранилище
func (fg *FileGenerator) StartGenerating() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// первый запуск сразу
	fg.scanAndStoreFiles()

	for range ticker.C {
		fg.scanAndStoreFiles()
	}
}

// сканирует директорию и сохраняет информацию о файлах
func (fg *FileGenerator) scanAndStoreFiles() {
	err := filepath.Walk(fg.rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		staticFile := model.StaticFile{
			Path:     path,
			Size:     info.Size(),
			MimeType: detectMimeType(path),
		}

		fg.storage.StoreItem(staticFile)

		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning files: %v\n", err)
	}
}

// определяет MIME-тип файла по расширению
func detectMimeType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}
