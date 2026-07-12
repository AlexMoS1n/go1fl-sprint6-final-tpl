package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(res http.ResponseWriter, req *http.Request) {
	content, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(res, "file not found", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Write(content)
}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "the request method must be POST", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(res, fmt.Sprintf("failed to parse form: %v", err), http.StatusBadRequest)
		return
	}

	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, fmt.Sprintf("error when receiving the file: %v", err), http.StatusBadRequest)
		return
	}

	defer file.Close()

	reader, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, fmt.Sprintf("error reading the file: %v", err), http.StatusInternalServerError)
		return
	}

	dataMorse, err := service.DecoderMorseCode(string(reader))
	if err != nil {
		http.Error(res, fmt.Sprintf("error in Morse data conversion: %v", err), http.StatusInternalServerError)
		return
	}

	fileMorseName := time.Now().Format("02_01_2006_15__04_05") + filepath.Ext(header.Filename)

	dirFileMorse := "./upload/morse/"

	err = os.MkdirAll(dirFileMorse, 0755)
	if err != nil {
		http.Error(res, fmt.Sprintf("error when creating a directory: %v", err), http.StatusInternalServerError)
		return
	}

	fileMorse, err := os.Create(dirFileMorse + fileMorseName)
	if err != nil {
		http.Error(res, fmt.Sprintf("error when creating the Morse file: %v", err), http.StatusInternalServerError)
		return
	}
	defer fileMorse.Close()

	_, err = fileMorse.WriteString(dataMorse)
	if err != nil {
		http.Error(res, fmt.Sprintf("error when writing data to the Morse file: %v", err), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(dataMorse))
}
