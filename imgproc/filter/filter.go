package filter

import "github.com/disintegration/imaging"

type Filter interface {
	Process(srcPath, dstPath string) error
}

type GrayScale struct{}

func (g GrayScale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Grayscale(src)
	err = imaging.Save(dst, dstPath)
	if err != nil {
		return err
	}
	return nil
}

type Blur struct{}

func (b Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Blur(src, 3.5)
	err = imaging.Save(dst, dstPath)
	if err != nil {
		return err
	}
	return nil
}
