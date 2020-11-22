package main

import (
	"fmt"
	"testing"

	"github.com/ualter/teachstore/session/service"
)

func TestMockListAll(t *testing.T) {
	sessionService := service.NewMockService()
	//sessionService := service.NewService()
	listSession, err := sessionService.ListAll(nil, nil, nil)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if len(listSession) == 0 {
		t.Errorf("Expected to found something on the list of Session")
	}
	for idx := range listSession {
		fmt.Printf("%d --> %+v\n", idx, listSession[idx])
	}
}

func setup() {

}
