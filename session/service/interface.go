package service

import (
	"net/http"
	"time"

	"github.com/ualter/teachstore-session/session/model"
)

type Service struct {
	List []model.Session
}

type ServiceAPI interface {
	ListAll(rw http.ResponseWriter, r *http.Request)
	Ping(rw http.ResponseWriter, r *http.Request)
}

func NewService() ServiceAPI {
	return &Service{
		List: []model.Session{
			model.Session{ID: 1, Name: "Angular", Date: time.Now()},
			model.Session{ID: 2, Name: "Javascript", Date: time.Now()},
		},
	}
}
