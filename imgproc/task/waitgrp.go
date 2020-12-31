package task

import (
	"fmt"
	"imgproc/filter"
	"path"
	"path/filepath"
	"sync"
)

type WaitGrpTask struct {
	dirCtx
	Filter filter.Filter
}

func NewWaitGrpTask(srcDir, dstDir string, filter filter.Filter) Tasker {
	return &WaitGrpTask{
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
		Filter: filter,
	}
}

func (w WaitGrpTask) Process() error {
	var wg sync.WaitGroup
	size := len(w.files)
	for i, f := range w.files {
		filename := filepath.Base(f)
		dst := path.Join(w.DstDir, filename)
		wg.Add(1)
		go w.applyFilter(f, dst, &wg, i+1, size)
	}
	wg.Wait()
	fmt.Println("Done processing files!")
	return nil
}

func (w *WaitGrpTask) applyFilter(src, dst string, wg *sync.WaitGroup, i, size int) {
	_ = w.Filter.Process(src, dst)
	fmt.Printf("Processes [%d/%d] %v => %v\n", i, size, src, dst)
	wg.Done()
}
