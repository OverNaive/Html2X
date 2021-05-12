package main

import (
	"encoding/json"
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

func htmlToPdf(rw http.ResponseWriter, r *http.Request) {
	var rb requestBody

	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	htmlPath := strings.Join([]string{now, ".html"}, "")
	pdfPath := strings.Join([]string{now, ".pdf"}, "")

	defer os.Remove(htmlPath)
	defer os.Remove(pdfPath)

	err = os.WriteFile(htmlPath, []byte(rb.Html), 0644)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	args := rb.Args
	args = append(args, htmlPath, pdfPath)

	cmd := exec.Command("wkhtmltopdf", args...)
	_, err = cmd.CombinedOutput()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(rw, r, pdfPath)
}

func main() {
	http.HandleFunc("/html-to-pdf", htmlToPdf)
	http.ListenAndServe(":8888", nil)
}
