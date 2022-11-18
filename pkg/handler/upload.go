package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (h *Handler) upload(c *gin.Context) {

	r := c.Request
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Printf("Er1: %s", err.Error())
	}
	n := r.FormValue("fileName")
	//n := func() string {
	//	return "test"
	//}()

	f, head, err := r.FormFile("pdfFile")
	if err != nil {
		fmt.Printf("Er2: %s", err.Error())
	}
	defer f.Close()
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + n + filepath.Ext(head.Filename)
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Er3: %s", err.Error())
	}
	defer file.Close()
	_, err = io.Copy(file, f)
	if err != nil {
		fmt.Printf("Er4: %s", err.Error())
	}
}
func (h *Handler) uploadForm(c *gin.Context) {
	c.HTML(http.StatusOK, "upload_pdf.tmpl.html", gin.H{
		"Title": "Upload",
	})
}
