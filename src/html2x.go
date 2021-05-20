package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func htmlToPdf(w http.ResponseWriter, r *http.Request) {
	var rb requestBody

	err := parseRequest(r, &rb)
	if err != nil {
		var mr *malformedRequest

		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	var pdfPath string
	pdfPath, err = rb.ToPdf()
	defer rb.Remove()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, pdfPath)
}

func htmlToImg(w http.ResponseWriter, r *http.Request) {
	var rb requestBody

	err := parseRequest(r, &rb)
	if err != nil {
		var mr *malformedRequest

		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	var imgPath string
	imgPath, err = rb.ToImg()
	defer rb.Remove()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, imgPath)
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
