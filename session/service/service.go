package service

import (
	"net/http"
	"time"

	api_models "github.com/ualter/teachstore-session/gen/models"
	"github.com/ualter/teachstore-session/session/model"
)

var mockData = []model.Session{
	model.Session{ID: 1, Name: "Angular", Date: time.Now(), Enrollment: &api_models.EnrollmentView{ID: 1}},
	model.Session{ID: 2, Name: "Javascript", Date: time.Now(), Enrollment: &api_models.EnrollmentView{ID: 2}},
}

type Service struct {
}

type IServiceRepository interface {
	List() *[]model.Session
	FindById(id int64) *model.Session
}

type IServiceAPI interface {
	GetAll(rw http.ResponseWriter, r *http.Request)
	GetByID(rw http.ResponseWriter, r *http.Request)
	Ping(rw http.ResponseWriter, r *http.Request)
}

func NewService() IServiceAPI {
	return &Service{}
}

func (s *Service) List() *[]model.Session {
	// TODO
	return &mockData
}

func (s *Service) FindById(id int64) *model.Session {
	// TODO
	for idx := range mockData {
		if mockData[idx].ID == id {
			return &mockData[idx]
		}
	}
	return nil
}
