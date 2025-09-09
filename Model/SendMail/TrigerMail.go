package TrigerMailModel

type SendMailRequest struct {
	Subject string `json:"subject" binding : "required"`
	Content string `json:"content" binding : "required"`
}

type SendMailReponce struct {
	Status  bool   `json:"status" binding : "required"`
	Message string `json:"message" binding :"required"`
}
