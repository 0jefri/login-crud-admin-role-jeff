package handler

import (
	"fmt"

	"github.com/go-embed-go-web/model"
	"github.com/go-embed-go-web/service"
)

type SiswaHandler struct {
	service service.SiswaService
}

func NewSiswaHandler(service service.SiswaService) *SiswaHandler {
	return &SiswaHandler{service: service}
}

func (h *SiswaHandler) AddStudent() {
	var name string
	var class string
	fmt.Print("Enter student name: ")
	fmt.Scan(&name)
	fmt.Print("Enter student class: ")
	fmt.Scan(&class)

	siswa := model.Siswa{Name: name, Class: class}
	err := h.service.RegisterNewSiswa(siswa)
	if err != nil {
		fmt.Println("Error adding student:", err)
	} else {
		fmt.Println("Student added successfully!")
	}
}

func (h *SiswaHandler) UpdateStudent() {
	var id int
	var name string
	var class string
	fmt.Print("Enter student ID to update: ")
	fmt.Scan(&id)
	fmt.Print("Enter new student name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new student class: ")
	fmt.Scan(&class)

	student := model.Siswa{ID: id, Name: name, Class: class}
	err := h.service.UpdateDataSiswa(student)
	if err != nil {
		fmt.Println("Error updating student:", err)
	} else {
		fmt.Println("Student updated successfully!")
	}
}

func (h *SiswaHandler) ViewStudents() {
	students, err := h.service.GetAllSiswa()
	if err != nil {
		fmt.Println("Error retrieving students:", err)
		return
	}

	fmt.Println("List of Students:")
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Class: %s\n", student.ID, student.Name, student.Class)
	}
}

func (h *SiswaHandler) DeleteStudent() {
	var id int
	fmt.Print("Enter student ID to delete: ")
	fmt.Scan(&id)

	err := h.service.DeleteSiswa(id)
	if err != nil {
		fmt.Println("Error deleting student:", err)
	} else {
		fmt.Println("Student deleted successfully!")
	}
}
