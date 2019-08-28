package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/locug/sil"
	"github.com/locug/sil/loc"
)

func (p *program) makeOBJ(w http.ResponseWriter, r *http.Request) {
	// make types to unmarshal obj table into
	// have to directly call the data type
	// will convert into a SIL down the line
	type View struct {
		Name     string
		Required bool
		Data     []loc.ObjTab
	}
	type SIL struct {
		Header sil.Header
		View   View
		Footer sil.Footer
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("error reading body: %v", err)

	}
	if err = r.Body.Close(); err != nil {
		log.Panicf("error closing the body: %v", err)
	}

	var ts SIL

	err = json.Unmarshal(body, &ts)
	if err != nil {
		_ = p.logger.Info(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Panicf("error encoding JSON: %v", err)
		}
	}

	var v sil.View

	v.Name = "OBJ"

	for _, eachData := range ts.View.Data {
		v.Data = append(v.Data, eachData)
	}

	s := sil.SIL{
		Header:    ts.Header,
		View:      v,
		Footer:    ts.Footer,
		TableType: loc.ObjTab{},
	}

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()))
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}

func (p *program) makeCLT(w http.ResponseWriter, r *http.Request) {
	// make types to unmarshal obj table into
	// have to directly call the data type
	// will convert into a SIL down the line
	type View struct {
		Name     string
		Required bool
		Data     []loc.CltTab
	}
	type SIL struct {
		Header sil.Header
		View   View
		Footer sil.Footer
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("error reading body: %v", err)

	}
	if err = r.Body.Close(); err != nil {
		log.Panicf("error closing the body: %v", err)
	}

	var ts SIL

	err = json.Unmarshal(body, &ts)
	if err != nil {
		_ = p.logger.Info(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Panicf("error encoding JSON: %v", err)
		}
	}

	var v sil.View

	v.Name = "CLT"

	for _, eachData := range ts.View.Data {
		v.Data = append(v.Data, eachData)
	}

	s := sil.SIL{
		Header:    ts.Header,
		View:      v,
		Footer:    ts.Footer,
		TableType: loc.CltTab{},
	}

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()))
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}

func (p *program) makeCLL(w http.ResponseWriter, r *http.Request) {
	// make types to unmarshal obj table into
	// have to directly call the data type
	// will convert into a SIL down the line
	type View struct {
		Name     string
		Required bool
		Data     []loc.CllTab
	}
	type SIL struct {
		Header sil.Header
		View   View
		Footer sil.Footer
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("error reading body: %v", err)

	}
	if err = r.Body.Close(); err != nil {
		log.Panicf("error closing the body: %v", err)
	}

	var ts SIL

	err = json.Unmarshal(body, &ts)
	if err != nil {
		_ = p.logger.Info(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Panicf("error encoding JSON: %v", err)
		}
	}

	var v sil.View

	v.Name = "CLL"

	for _, eachData := range ts.View.Data {
		v.Data = append(v.Data, eachData)
	}

	s := sil.SIL{
		Header:    ts.Header,
		View:      v,
		Footer:    ts.Footer,
		TableType: loc.CllTab{},
	}

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()))
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}

func (p *program) makeOFR(w http.ResponseWriter, r *http.Request) {
	// make types to unmarshal obj table into
	// have to directly call the data type
	// will convert into a SIL down the line
	type View struct {
		Name     string
		Required bool
		Data     []loc.OfrTab
	}
	type SIL struct {
		Header sil.Header
		View   View
		Footer sil.Footer
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicf("error reading body: %v", err)

	}
	if err = r.Body.Close(); err != nil {
		log.Panicf("error closing the body: %v", err)
	}

	var ts SIL

	err = json.Unmarshal(body, &ts)
	if err != nil {
		_ = p.logger.Info(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Panicf("error encoding JSON: %v", err)
		}
	}

	var v sil.View

	v.Name = "OFR"

	for _, eachData := range ts.View.Data {
		v.Data = append(v.Data, eachData)
	}

	s := sil.SIL{
		Header:    ts.Header,
		View:      v,
		Footer:    ts.Footer,
		TableType: loc.OfrTab{},
	}

	err = s.Write(fmt.Sprintf("%s_%d.sil", s.View.Name, time.Now().Nanosecond()))
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}
