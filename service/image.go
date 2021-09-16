package service

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
)

// GetThumbnail 生成缩略图
func GetThumbnail(path string) {
	imgData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	image = imaging.Resize(image, 200, 200, imaging.NearestNeighbor)
	err = imaging.Save(image, "3.jpg")
	if err != nil {
		fmt.Println(err)
	}
}
