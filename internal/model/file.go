package model

type File struct {
    name  string `json:"name"`
    path  string `json:"path"`
    type_ string `json:"type"`
}

func NewFile(name, path, type_ string) *File {
    return &File{name: name, path: path, type_: type_}
}

func (f *File) GetName() string {
    return f.name
}

func (f *File) GetPath() string {
    return f.path
}

func (f *File) GetType() string {
    return f.type_
}

func (f *File) UpdatePath(newPath string) {
    f.path = newPath
}
