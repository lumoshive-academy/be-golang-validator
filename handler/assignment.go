package handler

import (
	"encoding/json"
	"fmt"
	"go-23/model"
	"go-23/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AssignmentHandler struct {
	Service service.Service
}

func NewAssignmentHandler(server service.Service) AssignmentHandler {
	return AssignmentHandler{
		Service: server,
	}
}

func (assignmentHandler *AssignmentHandler) ListAssignments(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Ambil data assignments dari service
	assignments, err := assignmentHandler.Service.AssignmentService.GetAllAssignments()
	if err != nil {
		http.Error(w, "Failed to fetch assignments: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ambil data user (student) untuk menampilkan namanya
	student, err := assignmentHandler.Service.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch student data", http.StatusInternalServerError)
		return
	}

	// Kirim ke template
	data := struct {
		StudentName string
		Assignments []model.Assignment
	}{
		StudentName: student.Name,
		Assignments: assignments,
	}

	var response = struct {
		Status  bool
		Message string
		Data    any
	}{
		Status:  false,
		Message: "invalid Email or Password",
		Data:    data,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (assignmentHandler *AssignmentHandler) SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form for file upload (max 10MB)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Gagal membaca form: "+err.Error(), http.StatusBadRequest)
			return
		}

		assignmentID, err := strconv.Atoi(r.FormValue("assignment_id"))
		if err != nil {
			http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
			return
		}

		studentID, err := strconv.Atoi(r.FormValue("student_id"))
		if err != nil {
			http.Error(w, "Invalid student ID", http.StatusBadRequest)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "File tidak valid: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		status, err := assignmentHandler.Service.AssignmentService.SubmitAssignment(studentID, assignmentID, file, fileHeader)
		if err != nil {
			http.Error(w, "Gagal submit: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Berhasil submit dengan status: " + status))
	}
}

func (h *AssignmentHandler) ShowSubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/student/home", http.StatusSeeOther)
		return
	}

	assignmentIDStr := r.URL.Query().Get("assignment_id")
	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	assignment, err := h.Service.AssignmentService.GetAssignmentByID(assignmentID)
	if err != nil {
		http.Error(w, "Assignment not found", http.StatusNotFound)
		return
	}

	// Ambil user_id dari cookie
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	data := struct {
		Assignment  model.Assignment
		StudentID   int
		StudentName string
	}{
		Assignment:  *assignment,
		StudentID:   user.ID,
		StudentName: user.Name,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println(data)
}
