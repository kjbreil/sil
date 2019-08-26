package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/locug/sil"
	"github.com/locug/sil/loc"
)

func (p *program) makeOBJ(w http.ResponseWriter, r *http.Request) {

	_ = p.logger.Infof("got request\n")

	var obj loc.ObjTab

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		_ = p.logger.Info(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			panic(err)
		}
	}

	var s sil.SIL

	s.View.Name = "OBJ"

	s.View.Data = append(s.View.Data, obj)

	err = s.Write("out.sil")
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}

// makeSil is used to create the base sil type
func (p *program) makeSIL(w http.ResponseWriter, r *http.Request) {

	_ = p.logger.Infof("Got POST for new SIL file creation\n")

	// create the SIL file
	var s sil.SIL

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	// the JSON received should contain only the SIL header
	// Un marshalling the header into the sil
	err = json.Unmarshal(body, &s.Header)
	if err != nil {
		_ = p.logger.Infof("unmarshal failed with: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			_ = p.logger.Infof("error marshal failed with: %v", err)
		}
	}
	vars := mux.Vars(r)
	silType := vars["silType"]

	// hmmm this seems like it should be better or at least automatically
	// created based on the types in the loc package
	// there is no validation of anything right now, this will create problems
	// and hopefully errors but I haven't been great at making sure the right
	// errors are caught and passed in the sil package
	switch silType {
	case "OBJ":
		s.View.Name = "OBJ"
	}

	n := p.next()

	// error if n = -1
	if n == -1 {
		_ = p.logger.Infof("no more active sets in service, something is wrong, restart service")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode("no more active sets in service, something is wrong, restart service")
		if err != nil {
			_ = p.logger.Infof("error marshal failed with: %v", err)
		}
	}

	p.active[n] = &s

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(n)
	if err != nil {
		_ = p.logger.Infof("error marshal failed with: %v", err)
	}
}

// add adds data to the sil file
// WARNING NO VALIDATION RIGHT NOW
func (p *program) add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Post for addition to sil")
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["silID"])

	s, ok := p.active[id]
	if !ok {
		//todo: handle error
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	switch s.View.Name {
	case "OBJ":
		// assign v the object type
		var v loc.ObjTab
		err = json.Unmarshal(body, &v)
		p.returnErr(w, err)
		s.View.Data = append(s.View.Data, v)
	}

}

func (p *program) returnErr(w http.ResponseWriter, err error) {
	if err != nil {
		_ = p.logger.Infof("unmarshal failed with: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			_ = p.logger.Infof("error marshal failed with: %v", err)
		}
	}
}

// add adds data to the sil file
// WARNING NO VALIDATION RIGHT NOW
func (p *program) write(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["silID"])

	s, ok := p.active[id]
	if !ok {
		//todo: handle error
		return
	}

	err := s.Write(fmt.Sprintf("%08d.sil", id))
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}
