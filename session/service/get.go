package service

import (
	"net/http"

	tracing "github.com/ualter/teachstore/tracing"
	"github.com/ualter/teachstore/utils"

	log "github.com/sirupsen/logrus"
)

// swagger:route GET /session sessions listSessions
// Return a list of sessions
// responses:
//	200: sessionsResponse
// ListAll handles GET requests and returns all current sessions
func (s *Service) ListAll(rw http.ResponseWriter, r *http.Request) {
	log.Info("ListAll")
	rw.Header().Add("Content-Type", "application/json")

	tracing.TraceRequest("ListAll", rw, r)

	err := utils.ToJSON(s.List, rw)
	if err != nil {
		log.Error(err.Error())
	}
}
