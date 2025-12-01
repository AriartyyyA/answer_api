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
		httpHandler: transport.NewHTTPHandlers(),
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/health").Methods("GET").HandlerFunc(s.httpHandler.Health)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
