package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 20 //50MB
)

// 文件上传示例
func main() {
	r := RegisterHandlers()

	http.ListenAndServe(":8080", r)
}

//RegisterHandlers ...
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/upload", uploadHandler)

	return router
}
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("File is too big")
		return
	}
	file, headers, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		return
	}
	//获取上传文件的类型
	if headers.Header.Get("Content-Type") != "image/png" {
		log.Printf("只允许上传png图片")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		return
	}
	fn := headers.Filename
	err = ioutil.WriteFile("./video/"+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
