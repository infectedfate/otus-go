package model

type ServerConfig struct {
	Port int `json:"port"`
	StaticDir string `json:"static_dir"`
}
