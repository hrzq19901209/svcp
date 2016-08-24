package main

import (
	"os"
)

type Dockerfile struct {
	baseImage string
	execPath  string
	cmd       string
}

func (df *Dockerfile) GetPath() string {
	fout, err := os.Create("Dockerfile")
	if err != nil {
		panic(err)
	}

	defer fout.Close()
	fout.WriteString("FROM " + df.baseImage + "\r\n")
	return "Dockerfile"
}

func NewDockerfile(baseImage_ string, execPath_ string, cmd_ string) *Dockerfile {
	return &Dockerfile{
		baseImage: baseImage_,
		execPath:  execPath_,
		cmd:       cmd_,
	}
}

func main() {
	df := NewDockerfile("ubuntu", "path", "main")
	df.GetPath()
}
