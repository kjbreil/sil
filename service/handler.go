package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kjbreil/go-loc/loc"
	"github.com/kjbreil/sil"
)

type errReturn struct {
	Error errReturnBody
}
type errReturnBody struct {
	ID      int
	Message string
}

type sucReturn struct {
	Status string
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("error reading body: %v", err)
	}
	if err = r.Body.Close(); err != nil {
		log.Panicf("error closing the body: %v", err)
	}
	return body
}

func (p *program) returnError(w http.ResponseWriter, err error) {
	// put the error into error type and convert to JSON
	e := errReturn{
		Error: errReturnBody{
			ID:      422,
			Message: err.Error(),
		},
	}
	_ = p.logger.Info(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	// send the error string to the client
	// this wont error because we are just encoding the base error string
	_ = json.NewEncoder(w).Encode(e)
}

func returnSuccess(w http.ResponseWriter) {
	s := sucReturn{Status: "success"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(s)
}

// post header puts the header information for a post
// most of these are defaults in the sil package
func postHeader(s *sil.SIL) {
	s.Header.ActionType = sil.ADDRPL
	s.Header.Description = fmt.Sprintf("%s update from API", s.View.Name)
}

func (p *program) postOBJ(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var objs []loc.ObjTab

	err := json.Unmarshal(body, &objs)
	if err != nil {
		p.returnError(w, err)
		return
	}

	var s sil.SIL
	s.View.Name = "OBJ"
	for _, each := range objs {
		s.View.Data = append(s.View.Data, each)
	}

	postHeader(&s)

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()), true, true)
	if err != nil {
		p.returnError(w, err)
		return
	}
	returnSuccess(w)
}

func (p *program) postCLT(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var clts []loc.CltTab

	err := json.Unmarshal(body, &clts)
	if err != nil {
		p.returnError(w, err)
		return
	}

	var s sil.SIL
	s.View.Name = "OBJ"
	for _, each := range clts {
		s.View.Data = append(s.View.Data, each)
	}

	postHeader(&s)

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()), true, true)
	if err != nil {
		p.returnError(w, err)
		return
	}
	returnSuccess(w)
}

func (p *program) postCLL(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var clls []loc.CllTab

	err := json.Unmarshal(body, &clls)
	if err != nil {
		p.returnError(w, err)
		return
	}

	var s sil.SIL
	s.View.Name = "CLL"
	for _, each := range clls {
		s.View.Data = append(s.View.Data, each)
	}

	postHeader(&s)

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()), true, true)
	if err != nil {
		p.returnError(w, err)
		return
	}
	returnSuccess(w)
}

func (p *program) postOFR(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var ofrs []loc.OfrTab

	err := json.Unmarshal(body, &ofrs)
	if err != nil {
		p.returnError(w, err)
		return
	}

	var s sil.SIL
	s.View.Name = "OFR"
	for _, each := range ofrs {
		s.View.Data = append(s.View.Data, each)
	}

	postHeader(&s)

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()), true, true)
	if err != nil {
		p.returnError(w, err)
		return
	}
	returnSuccess(w)
}
