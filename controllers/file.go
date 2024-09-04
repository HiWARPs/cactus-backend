package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HiWARPs/cactus-backend/models"
	"github.com/unrolled/render"
)

const uploadDir = "uploads"

type ElectronDocument struct {
	Data []map[string]interface{} `json:"data"`
}

func NewElectronDocument() *ElectronDocument {
	data := []map[string]interface{}{}

	return &ElectronDocument{
		Data: data,
	}
}

// Creates a copy of the uploaded CSV file
func createFile(file multipart.File, fh *multipart.FileHeader) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to parse file %s", err)
	}

	filePath := createFileName(fh.Filename)
	newFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file %s", err)
	}
	defer newFile.Close()

	_, err = newFile.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to write file %s", err)
	}

	return filePath, nil
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

// Converts CSV file into JSON
func csvToDocument(filePath string) (*ElectronDocument, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csvHeaders, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read data %s", err)
	}

	doc := NewElectronDocument()

	for {
		electrons, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to read CSV data: %s", err)
		}

		electron := make(map[string]interface{})

		for i, header := range csvHeaders {
			header = clearString(header)
			value := electrons[i]

			// Convert the value to a float64 (you can use other numeric types if needed)
			numValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				// Handle the error (e.g., log it or use a default value)
				numValue = 0.0 // Default value
			}

			electron[header] = numValue
		}

		doc.Data = append(doc.Data, electron)
	}

	return doc, nil
}

// Names the uploaded CSV file
func createFileName(fh string) string {
	ts := time.Now().Unix()
	newFileName := strings.Replace(fh, " ", "_", -1)
	return fmt.Sprintf("./%s/%d_%s", uploadDir, ts, newFileName)
}

// Upload request
func UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("csvfile")
		if err != nil {
			http.Error(w, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		filePath, err := createFile(file, fileHeader)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to create file", http.StatusBadRequest)
			return
		}

		doc, err := csvToDocument(filePath)
		if err != nil {
			http.Error(w, "Failed to convert file", http.StatusBadRequest)
			return
		}

		result, err := mh.Upload(doc)
		if err != nil {
			log.Println("MongoDB upload error:", err)
			return
		}

		renderOutput := render.New()
		renderOutput.JSON(w, 200, map[string]interface{}{"message": "file uploaded successfully",
			"fileID": result.InsertedID})
	}
}

func QueryElectrons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		query := models.Request{}
		err := json.NewDecoder(r.Body).Decode(&query)
		if err != nil {
			http.Error(w, "error decoding payload", http.StatusBadRequest)
			log.Print(err)
			return
		}

		if err := query.IsValid(query); err != nil {
			http.Error(w, "multiple X queries detected", http.StatusBadRequest)
			return
		}

		electron, err := mh.QueryElectron(query)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}
		renderOutput := render.New()
		renderOutput.JSON(w, 200, electron)
	}
}
