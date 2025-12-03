package server

import (
	"errors"
	"github/Ariartyyy/answer_api/internal/transport"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandler *transport.HTTPHandlers
}

func NewHTTPServer(httpHandlers *transport.HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandler: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/questions", s.httpHandler.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", s.httpHandler.GetQuestionByIDWithAnswers).Methods("GET")
	router.HandleFunc("/questions/{id}", s.httpHandler.DeleteQuestionByID).Methods("DELETE")
	router.HandleFunc("/questions", s.httpHandler.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{question_id}/answers", s.httpHandler.AddAnswerToQuestionHandler).Methods("POST")
	router.HandleFunc("/answers/{id}", s.httpHandler.GetAnswerByIDHandler).Methods("GET")
	router.HandleFunc("/answers/{id}", s.httpHandler.DeleteAnswerByIDHandler).Methods("DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
