package main

import (
	"flag"
	"fmt"
	"imgproc/filter"
	"imgproc/task"
	"time"
)

func main() {

	srcDir := flag.String("src", "", "Input directory")
	dstDir := flag.String("dst", "", "Output directory")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "waitgrp", "waitgrp/channel")
	poolSize := flag.Int("poolsize", 4, "Workers pool size of the channel task")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.GrayScale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	start := time.Now()
	_ = t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Images processing took %v\n", elapsed)

}
