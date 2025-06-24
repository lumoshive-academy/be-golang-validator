// handler/lecturer_handler.go
package handler

import (
	"encoding/json"
	"fmt"
	"go-23/dto"
	"go-23/service"
	"go-23/utils"
	"net/http"
	"strconv"
)

type SubmissionHandler struct {
	Service service.Service
}

func NewSubmissionHandler(service service.Service) SubmissionHandler {
	return SubmissionHandler{
		Service: service,
	}
}

func (h *SubmissionHandler) Home(w http.ResponseWriter, r *http.Request) {
	submissions, err := h.Service.SubmissionService.GetAllSubmissions()
	if err != nil {
		http.Error(w, "Gagal mengambil data submission", http.StatusInternalServerError)
		return
	}
	fmt.Println(submissions)
}

func (h *SubmissionHandler) ShowGradeForm(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.URL.Query().Get("student_id")
	assignmentIDStr := r.URL.Query().Get("assignment_id")

	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		http.Error(w, "Invalid student_id", http.StatusBadRequest)
		return
	}

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		http.Error(w, "Invalid assignment_id", http.StatusBadRequest)
		return
	}

	// Ambil data untuk ditampilkan di form
	student, err := h.Service.UserService.GetUserByID(studentID)
	if err != nil {
		http.Error(w, "Student not found", http.StatusInternalServerError)
		return
	}

	assignment, err := h.Service.AssignmentService.GetAssignmentByID(assignmentID)
	if err != nil {
		http.Error(w, "Assignment not found", http.StatusInternalServerError)
		return
	}

	data := struct {
		StudentID       int
		AssignmentID    int
		StudentName     string
		AssignmentTitle string
	}{
		StudentID:       student.ID,
		AssignmentID:    assignment.ID,
		StudentName:     student.Name,
		AssignmentTitle: assignment.Title,
	}

	fmt.Println(data)
}

func (h *SubmissionHandler) GradeSubmission(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Gagal parsing form", http.StatusBadRequest)
	// 	return
	// }

	// studentID, err := strconv.Atoi(r.FormValue("student_id"))
	// if err != nil {
	// 	http.Error(w, "Invalid student_id", http.StatusBadRequest)
	// 	return
	// }

	// assignmentID, err := strconv.Atoi(r.FormValue("assignment_id"))
	// if err != nil {
	// 	http.Error(w, "Invalid assignment_id", http.StatusBadRequest)
	// 	return
	// }

	// gradeStr := r.FormValue("grade")
	// grade, err := strconv.ParseFloat(gradeStr, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid grade", http.StatusBadRequest)
	// 	return
	// }
	var gradeSubmission dto.GradeSubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&gradeSubmission); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "invalid request body")
		return
	}

	message, err := utils.ValidateData(gradeSubmission)
	if err != nil {
		utils.ResponseBadRequest2(w, http.StatusBadRequest, message)
		return
	}

	err = h.Service.SubmissionService.GradeSubmission(gradeSubmission.UserID, gradeSubmission.AssignmentID, gradeSubmission.Grade)
	if err != nil {
		http.Error(w, "Gagal memberi nilai: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success submitted", nil)
}
