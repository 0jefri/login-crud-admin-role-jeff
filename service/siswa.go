package service

import (
	"fmt"

	"github.com/go-embed-go-web/model"
	"github.com/go-embed-go-web/repository"
)

type SiswaService interface {
	RegisterNewSiswa(payload model.Siswa) error
	UpdateDataSiswa(siswa model.Siswa) error
	GetAllSiswa() ([]model.Siswa, error)
	DeleteSiswa(id int) error
}

type siswaService struct {
	repo repository.SiswaRepository
}

// DeleteSiswa implements SiswaService.
func (s *siswaService) DeleteSiswa(id int) error {
	return s.repo.Delete(id)
}

// UpdateSiswa implements SiswaService.
func (s *siswaService) GetAllSiswa() ([]model.Siswa, error) {
	return s.repo.List()
}

// UpdateDataSiswa implements SiswaService.
func (s *siswaService) UpdateDataSiswa(siswa model.Siswa) error {
	return s.repo.Update(siswa)
}

func NewSiswaService(swaRepo repository.SiswaRepository) SiswaService {
	return &siswaService{repo: swaRepo}
}

func (s *siswaService) RegisterNewSiswa(payload model.Siswa) error {
	if payload.Name == "" || payload.Class == "" {
		return fmt.Errorf("payload is required")
	}
	err := s.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create siswa: %s", err)
	}
	return nil
}
