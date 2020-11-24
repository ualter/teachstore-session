package service

/*
import (
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ualter/teachstore-session/session/model"
)

type mockService struct {
	List []model.Session
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

var contextLogger = log.WithFields(log.Fields{
	"mocktest": "true",
})

func NewMockService() ServiceAPI {
	contextLogger.Info("New MockService")
	return &mockService{
		List: []model.Session{
			model.Session{ID: 90, Name: "MOCK **** Angular", Date: time.Now()},
			model.Session{ID: 91, Name: "MOCK **** Javascript", Date: time.Now()},
		},
	}
}

func (s *mockService) ListAll(rw http.ResponseWriter, r *http.Request) ([]model.Session, error) {
	contextLogger.Info("Listing all sessions")
	return s.List, nil
}
*/
