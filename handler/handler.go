package handler

import "go-23/service"

type Handler struct {
	AssignmentHandler AssignmentHandler
	AuthHandler       AuthHandler
	SubmissionHandler SubmissionHandler
}

func NewHandler(service service.Service) Handler {
	return Handler{
		AssignmentHandler: NewAssignmentHandler(service),
		AuthHandler:       NewAuthHandler(service),
		SubmissionHandler: NewSubmissionHandler(service),
	}
}
