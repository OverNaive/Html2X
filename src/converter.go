package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type requestBody struct {
	Html string `json:"html"`
	Args map[string]string `json:"args"`
	Ext string `json:"ext"`
	htmlPath string
	extPath string
	cmdArgs []string
}

func (rb *requestBody) ToPdf() (string, error) {
	rb.Ext = "pdf"

	return rb.extPath, rb.toX("wkhtmltopdf")
}

func (rb *requestBody) ToImg() (string , error) {
	if rb.Ext == "" {
		rb.Ext = "jpg"
	}

	return rb.extPath, rb.toX("wkhtmltoimage")
}

func (rb *requestBody) Remove() {
	_ = os.Remove(rb.htmlPath)
	_ = os.Remove(rb.extPath)
}

func (rb *requestBody) toX(commandName string) (err error) {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	rb.htmlPath = strings.Join([]string{now, "html"}, ".")
	rb.extPath = strings.Join([]string{now, rb.Ext}, ".")

	err = os.WriteFile(rb.htmlPath, []byte(rb.Html), 0644)
	if err != nil {
		return
	}

	for k, v := range rb.Args {
		if v == "" {
			continue
		}

		rb.cmdArgs = append(rb.cmdArgs, "--" + k, v)
	}

	rb.cmdArgs = append(rb.cmdArgs, rb.htmlPath, rb.extPath)

	cmd := exec.Command(commandName, rb.cmdArgs...)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return
	}

	return
}
