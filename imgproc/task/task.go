package task

import (
	"fmt"
	"os"
	"path/filepath"
)

type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFileList(srcDir string) []string {
	var files []string
	fmt.Println("Generating file list...")
	_ = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".jpg" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}
