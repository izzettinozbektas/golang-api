package handlers

import (
	"github.com/izzettinozbektas/golang-api/internal/response"
	"io"
	"log"
	"net/http"
	"os"
)

func (m *Repository) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
		response.Write(w, response.Error(".errors.upload_image.cannot_read_file", nil), response.Code(http.StatusInternalServerError))

	}
	log.Println(handler.Filename)
	defer file.Close()

	// Create file locally
	dst, err := os.Create("./build/temp-images/" + handler.Filename)
	if err != nil {
		log.Fatal(err)
		response.Write(w, response.Error(".errors.upload_image.cannot_create_local_file", nil), response.Code(http.StatusInternalServerError))
	}
	defer dst.Close()

	// Copy the uploaded file data to the newly created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		log.Fatal(err)
		response.Write(w, response.Error(".errors.upload_image.cannot_copy_to_file", nil), response.Code(http.StatusInternalServerError))

		return
	}
	response.Write(w, response.Success("success", nil), response.Code(http.StatusOK))
}
