package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	// 将文件转化成二进制流传到client端
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//文件上传到server端
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(int64(MAX_UPLOAD_SIZE)); err != nil {
		log.Printf("Try to open file error: %v", err)
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload file successfully")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}
