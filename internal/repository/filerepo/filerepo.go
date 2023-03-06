package filerepo

import (
	"github.com/izzettinozbektas/golang-api/internal/repository"
)

type fileRepo struct {
}

func NewFileRepo() repository.FileRepo {
	return &fileRepo{}
}

func (m *fileRepo) UploadFile() {

}
