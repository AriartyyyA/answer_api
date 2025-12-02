package transport

import "github/Ariartyyy/answer_api/internal/service"

type HTTPHandlers struct {
	service *service.Service
}

func NewHTTPHandlers(service *service.Service) *HTTPHandlers {
	return &HTTPHandlers{
		service: service,
	}
}
