package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type requestBody struct {
	Html string `json:"html"`
	Args []string `json:"args"`
}

func htmlToPdf(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if ! strings.HasPrefix(req.Header.Get("Content-Type"), "application/json") {
		http.Error(w, "Invalid content type", http.StatusUnsupportedMediaType)
		return
	}

	var b requestBody

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	htmlPath := strings.Join([]string{now, ".html"}, "")
	pdfPath := strings.Join([]string{now, ".pdf"}, "")

	defer os.Remove(htmlPath)
	defer os.Remove(pdfPath)

	err = os.WriteFile(htmlPath, []byte(b.Html), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	args := b.Args
	args = append(args, htmlPath, pdfPath)

	cmd := exec.Command("wkhtmltopdf", args...)
	_, err = cmd.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, req, pdfPath)
}

func htmlToImg(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if ! strings.HasPrefix(req.Header.Get("Content-Type"), "application/json") {
		http.Error(w, "Invalid content type", http.StatusUnsupportedMediaType)
		return
	}

	var b requestBody

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	htmlPath := strings.Join([]string{now, ".html"}, "")
	imgPath := strings.Join([]string{now, ".jpg"}, "")

	defer os.Remove(htmlPath)
	defer os.Remove(imgPath)

	err = os.WriteFile(htmlPath, []byte(b.Html), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	args := b.Args
	args = append(args, htmlPath, imgPath)

	cmd := exec.Command("wkhtmltoimage", args...)
	_, err = cmd.CombinedOutput()
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

	http.HandleFunc("/html-to-pdf", htmlToPdf)
	http.HandleFunc("/html-to-img", htmlToImg)

	log.Fatal(s.ListenAndServe())
}
