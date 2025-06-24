package router

import (
	"go-23/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", handler.AuthHandler.Login)

	r.Route("/student", func(r chi.Router) {
		r.Get("/listassignment", handler.AssignmentHandler.ListAssignments)
		r.Post("/submit", handler.AssignmentHandler.SubmitAssignment)
	})

	r.Route("/lecturer", func(r chi.Router) {
		r.Post("/grade", handler.SubmissionHandler.GradeSubmission)
	})

	// r.Get("/student/submit", assignmentHandler.ShowSubmitForm)
	// r.Get("/lecturer/home", submissionHandler.Home)
	// r.Get("/lecturer/grade-form", submissionHandler.ShowGradeForm)
	return r
}
