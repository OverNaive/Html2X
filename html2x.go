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
	json.NewDecoder(r.Body).Decode(&rb)

	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	htmlPath := strings.Join([]string{now, ".html"}, "")
	pdfPath := strings.Join([]string{now, ".pdf"}, "")

	defer os.Remove(htmlPath)
	defer os.Remove(pdfPath)

	os.WriteFile(htmlPath, []byte(rb.Html), 0644)

	args := rb.Args
	args = append(args, htmlPath, pdfPath)

	cmd := exec.Command("wkhtmltopdf", args...)
	cmd.CombinedOutput()

	http.ServeFile(rw, r, pdfPath)
}

func main() {
	http.HandleFunc("/html-to-pdf", htmlToPdf)
	http.ListenAndServe(":8888", nil)
}
