package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func htmlToPdf(w http.ResponseWriter, req *http.Request) {
	var b requestBody

	err := parseRequest(req, &b)
	if err != nil {
		var mr *malformedRequest

		if errors.Is(err, mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	var pdfPath string
	pdfPath, err = b.ToPdf()
	defer b.Remove()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, req, pdfPath)
}

func htmlToImg(w http.ResponseWriter, req *http.Request) {
	var b requestBody

	err := parseRequest(req, &b)
	if err != nil {
		var mr *malformedRequest

		if errors.Is(err, mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	var imgPath string
	imgPath, err = b.ToImg()
	defer b.Remove()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, req, imgPath)
}

func main() {
	s := &http.Server{
		Addr: ":8888",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/wk/html-to-pdf", htmlToPdf)
	http.HandleFunc("/wk/html-to-img", htmlToImg)

	log.Fatal(s.ListenAndServe())
}
