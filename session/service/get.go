package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"

	tracing "github.com/ualter/teachstore-session/tracing"
	"github.com/ualter/teachstore-session/utils"

	log "github.com/sirupsen/logrus"

	//api_models "github.com/ualter/teachstore-session/gen/models"

	api_client "github.com/ualter/teachstore-session/gen/client"
	api_enrrollment "github.com/ualter/teachstore-session/gen/client/enrollment"
)

// swagger:route GET /sessions sessions listSessions
// Return a list of sessions
// responses:
//	200: sessionsResponse
// GetAll handles GET requests and returns all current sessions
func (s *Service) GetAll(rw http.ResponseWriter, r *http.Request) {
	log.Info("GetAll")
	rw.Header().Add("Content-Type", "application/json")
	tracing.TraceRequest("GetAll", rw, r)
	err := utils.ToJSON(s.List(), rw)
	if err != nil {
		log.Error(err.Error())
	}
	rw.WriteHeader(http.StatusOK)
}

// swagger:route GET /sessions/{id} sessions Session
// Return a session found by its Id
// responses:
//	200: Session
// GetById handles GET requests and returns a Session found bit its Id
func (s *Service) GetByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Error(err.Error())
	} else {
		log.Infof("GetById id=%d", id)
		rw.Header().Add("Content-Type", "application/json")
		tracing.TraceRequest("GetById", rw, r)
		session := s.FindById(id)
		if session != nil {
			///
			var clientEnrollment *api_client.EnrollmentAPI
			transportConfig := api_client.DefaultTransportConfig().WithHost("192.168.1.49")
			clientEnrollment = api_client.NewHTTPClientWithConfig(strfmt.Default, transportConfig)

			params := api_enrrollment.NewListUsingGETParams()
			// ----> AQUI!!   Change to FindByIDUsingGET
			// search the Enrollment by the Session Enrollment ID
			results, err := clientEnrollment.Enrollment.ListUsingGET(params)
			if err != nil {
				log.Error(err.Error())
			}
			payload := results.GetPayload()
			ok := payload != nil
			if !ok {
				msg := "Expected to receive something on the Payload"
				log.Errorf(msg)
				panic(errors.New(msg))
			}
			for idx := range payload {
				//enrollment := &models.EnrollmentView{}
				enrollment := payload[idx]
				fmt.Printf("Enrollment:\n Id:%d, Register Date:%s\n Course Title:%#v\n Student Name:%#v \n---------\n",
					enrollment.ID, enrollment.RegisterDate, enrollment.Course.Title, enrollment.Student.Name)
				session.Enrollments = append(session.Enrollments, enrollment)
			}
			///
			rw.WriteHeader(http.StatusOK)
			err = utils.ToJSON(session, rw)
			if err != nil {
				log.Error(err.Error())
			}

		} else {
			rw.WriteHeader(http.StatusNotFound)
			log.Infof("GetById id %d not found", id)
		}
	}
}

// swagger:route GET /session sessions listSessions
// Return a signal that the service is UP (for Kubernetes Probes use)
// responses:
//	200: "OK"
// Ping handles GET requests and returns a OK signal, if the service is alive
func (s *Service) Ping(rw http.ResponseWriter, r *http.Request) {
	log.Info("Ping")
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["result"] = "OK!"
	err := utils.ToJSON(res, rw)
	if err != nil {
		log.Error(err.Error())
	}

}
