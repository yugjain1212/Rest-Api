package storage

import "github.com/yugjain1212/student-api/internal/types"

type Storage interface {
	Createstudent(name string, email string, age int) (int64, error)
	Getstudentbyid(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
