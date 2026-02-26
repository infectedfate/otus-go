package model

// интерфейс для объектов, которые можно сохранять
type Storable interface {
	GetPath() string
}

// представляет статический файл
type StaticFile struct {
	Path     string
	Size     int64
	MimeType string
}

func (sf StaticFile) GetPath() string {
	return sf.Path
}
