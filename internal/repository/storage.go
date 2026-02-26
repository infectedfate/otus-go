// internal/repository/storage.go
package repository

import (
	"fmt"
	"strings"

	"otus-go/internal/model"
)

type Storage struct {
	htmlFiles  []model.StaticFile
	cssFiles   []model.StaticFile
	jsFiles    []model.StaticFile
	imageFiles []model.StaticFile
}

func NewStorage() *Storage {
	return &Storage{
		htmlFiles:  []model.StaticFile{},
		cssFiles:   []model.StaticFile{},
		jsFiles:    []model.StaticFile{},
		imageFiles: []model.StaticFile{},
	}
}

func (s *Storage) StoreItem(item model.Storable) {
	switch v := item.(type) {
	case model.StaticFile:
		switch {
		case strings.HasSuffix(v.Path, ".html"):
			s.htmlFiles = append(s.htmlFiles, v)
			fmt.Printf("Stored HTML file: %+v\n", v)
		case strings.HasSuffix(v.Path, ".css"):
			s.cssFiles = append(s.cssFiles, v)
			fmt.Printf("Stored CSS file: %+v\n", v)
		case strings.HasSuffix(v.Path, ".js"):
			s.jsFiles = append(s.jsFiles, v)
			fmt.Printf("Stored JS file: %+v\n", v)
		case strings.HasSuffix(v.Path, ".png") ||
			strings.HasSuffix(v.Path, ".jpg") ||
			strings.HasSuffix(v.Path, ".gif"):
			s.imageFiles = append(s.imageFiles, v)
			fmt.Printf("Stored image file: %+v\n", v)
		default:
			fmt.Printf("Unrecognized file type: %+v\n", v)
		}
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func (s *Storage) GetHTMLFiles() []model.StaticFile {
	return s.htmlFiles
}

func (s *Storage) GetCSSFiles() []model.StaticFile {
	return s.cssFiles
}

func (s *Storage) GetJSFiles() []model.StaticFile {
	return s.jsFiles
}

func (s *Storage) GetImageFiles() []model.StaticFile {
	return s.imageFiles
}
