package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/yugjain1212/student-api/internal/storage"
	"github.com/yugjain1212/student-api/internal/types"
	"github.com/yugjain1212/student-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var student types.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if errors.Is(err, io.EOF) {
		response.Writejson(w, http.StatusBadRequest, response.Generalerror(errors.New("request body is empty")))
		return
	}
	if err != nil {
		response.Writejson(w, http.StatusBadRequest, response.Generalerror(errors.New("invalid JSON format")))
		return
	}
		//request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.Writejson(w, http.StatusBadRequest, response.Validationerror(validateErrs))
			return
		}

		id, err := storage.Createstudent(
			student.Name,
			student.Email,
			student.Age,
		)
		if err != nil {
			response.Writejson(w, http.StatusInternalServerError, response.Generalerror(err))
			return
		}
		slog.Info("user created", slog.String("userid", fmt.Sprintf("%d", id)))
		response.Writejson(w, http.StatusCreated, map[string]int64{"id": id})

	}
}
func Getbyid(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("get by id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.Writejson(w, http.StatusBadRequest, response.Generalerror(err))
			return
		}

		student, err := storage.Getstudentbyid(intId)
		if err != nil {
			slog.Error("Error getting student", slog.String("id", id), slog.String("error", err.Error()))
			response.Writejson(w, http.StatusInternalServerError, response.Generalerror(err))
			return
		}

		response.Writejson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			slog.Error("Error getting students", slog.String("error", err.Error()))
			response.Writejson(w, http.StatusInternalServerError, response.Generalerror(err))
			return
		}

		response.Writejson(w, http.StatusOK, students)
	}
}
