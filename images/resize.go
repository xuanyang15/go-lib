package images

import (
	"os"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"fmt"
	"errors"
)

type LengthType uint

type ImageSize struct {
	Width 	LengthType	`json:"width"`
	Height 	LengthType	`json:"height"`
}

type ImageResizer struct {
	FilePtr 		*os.File
	TargetFolder	string
	TargetName 		string
	TargetSizes		[]ImageSize
}

func NewImageResizer(file *os.File, folder, name string, szs []ImageSize) (ir *ImageResizer){
	for folder[len(folder)-1] == '/' {
		folder = folder[:len(folder)-1]
	}

	ir = &ImageResizer{
		FilePtr:  		file,
		TargetFolder: 	folder,
		TargetName: 	name,
		TargetSizes: 	make([]ImageSize, 0, 0),
	}

	if szs != nil {
		for _, sz := range szs {
			ir.AddSize(sz.Width, sz.Height)
		}
	}

	return ir
}

func (this* ImageResizer) AddSize(width, height LengthType) {
	this.TargetSizes = append(this.TargetSizes, ImageSize{Width: width, Height: height})
}

func (this *ImageResizer) Resize() (paths []string, err error) {
	paths = make([]string, 0, 0)
	if this.TargetSizes == nil {
		return
	}

	img, fmtName, err := image.Decode(this.FilePtr)
	if err != nil {
		return
	}

	for _, sz := range this.TargetSizes{
		m := resize.Resize(uint(sz.Width), uint(sz.Height), img, resize.Lanczos3)
		outPath := fmt.Sprintf("%s/%d-%d__%s", this.TargetFolder, sz.Width, sz.Height, this.TargetName)
		outFile, theErr := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if theErr != nil {
			err = theErr
			return
		}
		if fmtName == `jpeg` {
			jpeg.Encode(outFile, m, nil)
		} else if fmtName == `png` {
			png.Encode(outFile, m)
		} else {
			err = errors.New(fmt.Sprintf("%s is not a supported type", fmtName))
			return
		}

		outFile.Close()
		paths = append(paths, outPath)
	}

	return
}