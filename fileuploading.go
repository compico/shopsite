package main

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/compico/shopsite/internal/utils"
)

func fileupload(name string, id int, mfh multipart.FileHeader) (string, error) {
	nameid := utils.Transcript(strconv.Itoa(id) + "_" + name)
	pathtofile := "public/" + "image/" + "products/" + nameid + "/" + mfh.Filename
	if _, err := os.Stat(pathtofile); os.IsNotExist(err) {
		err = os.MkdirAll("public/"+"image/"+"products/"+nameid+"/", 0644)
		if err != nil {
			return "", err
		}
	}
	file, err := os.Create(pathtofile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	x, err := mfh.Open()
	if err != nil {
		return "", err
	}
	defer x.Close()
	_, err = io.Copy(file, x)
	if err != nil {
		return "", err
	}
	return pathtofile, nil
}
