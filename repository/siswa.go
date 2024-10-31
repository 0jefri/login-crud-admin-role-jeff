package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-embed-go-web/model"
)

type SiswaRepository interface {
	Create(siswa model.Siswa) error
	Update(siswa model.Siswa) error
	List() ([]model.Siswa, error)
	// Get(id int) (model.Siswa, error)
	Delete(id int) error
}

type siswaRepository struct {
	db *sql.DB
}

// Delete implements SiswaRepository.
func (s *siswaRepository) Delete(id int) error {
	query := `DELETE FROM siswa WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete student: %v", err)
	}
	return nil
}

// List implements SiswaRepository.
func (s *siswaRepository) List() ([]model.Siswa, error) {
	query := `SELECT id, name, class FROM siswa`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed view data siswa")
	}
	defer rows.Close()

	var students []model.Siswa
	for rows.Next() {
		var student model.Siswa
		err := rows.Scan(&student.ID, &student.Name, &student.Class)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

// Update implements SiswaRepository.
func (s *siswaRepository) Update(siswa model.Siswa) error {
	query := `UPDATE siswa SET name = $1, class = $2 WHERE id = $3`
	_, err := s.db.Exec(query, siswa.Name, siswa.Class, siswa.ID)
	if err != nil {
		fmt.Println("err:", err)
		return fmt.Errorf("failed update data siswa")
	}
	return nil
}

func NewSiswaRepository(db *sql.DB) SiswaRepository {
	return &siswaRepository{db: db}
}

func (s *siswaRepository) Create(siswa model.Siswa) error {
	query := `INSERT INTO Siswa(name, class) VALUES($1, $2)`
	_, err := s.db.Exec(query, siswa.Name, siswa.Class)
	if err != nil {
		return fmt.Errorf("failed to add siswa: %v", err)
	}
	return nil
}
