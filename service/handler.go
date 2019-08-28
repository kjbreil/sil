package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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

	_ = p.logger.Infof("got request\n")

	//var obj loc.ObjTab

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

	err = s.Write("out.sil")
	if err != nil {
		_ = p.logger.Errorf("sil file could not be written: %v", err)
	}
}
