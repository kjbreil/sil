package silread

import (
	"errors"
	"fmt"
	"github.com/kjbreil/loc-macro/pkg/macro"
	"github.com/kjbreil/loc-macro/pkg/script"
	"os"
	"strings"
)

type Stats struct {
	HasHeader bool
	Header    []string
	HasView   bool
	Table     string
	DataLines int
	FCodes    []string
	HasCreate bool
}

func GetStats(filename string) (Stats, error) {
	s := Stats{}

	f, err := os.Open(filename)
	if err != nil {
		return s, err
	}

	p := newParser(f)
	var d decoder

	var i int
	for {
		pt := p.scan()
		d.p = append(d.p, *pt)
		if pt.tok == CRLF {
			if !s.HasHeader && d.headerInsert {
				s.HasHeader = true
			}
			if len(s.Header) == 0 && len(d.header) > 0 {
				s.Header = d.header
			}
			if !s.HasView && d.view {
				s.HasView = true
				s.Table = d.tableName
				s.FCodes = d.fcodes
			}

			ni := d.identifyLine(i)
			if ni == i {
				break
			}
			i = 0
			d.p = d.p[:0]

			if d.view && len(d.data) > 0 {
				s.DataLines++
				d.data = d.data[:0]
			}
		}
		if pt.tok == EOF {
			break
		}
	}

	return s, nil
}

var ErrInvalidTable = errors.New("invalid table name")

func (r *Reader) newStatsFromReaderOnly(validTables []string) error {
	// initialize a decoder for getting the stats
	var err error
	macrosSB := &strings.Builder{}
	var prsd parsed

mainLoop:
	for {
		// clear prsd
		prsd = prsd[:0]
		// capture a line into prsd
		for {

			pt := r.p.scan()
			prsd = append(prsd, *pt)
			if pt.tok == CRLF {
				break
			}
			if pt.tok == EOF {
				break mainLoop
			}
		}

		// check the first part to see what to do with this line
		switch prsd[0].val {
		case "CREATE":
			var action, table string
			var fcodes []string
			action, table, fcodes, err = prsd.getCreate(0)
			if err != nil {
				return err
			}
			if action != "DCT" {
				r.Table = table
				r.FCodes = fcodes
				if !isValidTable(table, validTables) {
					return fmt.Errorf("%w: %s", ErrInvalidTable, table)
				}

			}
		case "INSERT":
			name := prsd.getTable(0)
			if prsd.isInsert(0, name) {
				r.HasView = true
			}
		case "(":
			r.DataLines++
		}

		if prsd[0].val[0] == '@' {
			prsd.stringBuilder(macrosSB, 0, len(prsd)-1)
		}

	}
	// TODO: Do checks to make sure stats are valid

	// check for macros
	r.macros, err = script.Read(strings.NewReader(macrosSB.String()))
	if err != nil {
		return err
	}

	for m := range r.macros.Macros().Range {
		if mm, ok := m.(*macro.Macro); ok {
			switch mm.Def.Name {
			case "CREATE":
				tableName := mm.ParametersParts.Index(0).String()
				if !strings.Contains(strings.ToUpper(tableName), strings.ToUpper(r.Stats.Table)) {
					return fmt.Errorf("CREATE table does not match table name: %s", tableName)
				}
				r.HasCreate = true
				// case "UPDATE_BATCH":
				// 	_, _ = mm.ParametersParts.FindSetter("JOB")
			}
		}

	}

	return nil
}

func isValidTable(tableName string, validTables []string) bool {

	for _, rd := range validTables {
		if strings.EqualFold(rd, fmt.Sprintf("%s_LOAD", tableName)) {
			return true
		}
	}
	return false
}

func (r *Reader) newStatsFromReader() error {
	// initialize a decoder for getting the stats
	var err error
	macrosSB := &strings.Builder{}
	var prsd parsed

mainLoop:
	for {
		// clear prsd
		prsd = prsd[:0]
		// capture a line into prsd
		for {

			pt := r.p.scan()
			prsd = append(prsd, *pt)
			if pt.tok == CRLF {
				break
			}
			if pt.tok == EOF {
				break mainLoop
			}
		}

		// check the first part to see what to do with this line
		switch prsd[0].val {
		case "CREATE":
			var action, table string
			var fcodes []string
			action, table, fcodes, err = prsd.getCreate(0)
			if err != nil {
				return err
			}
			if action != "DCT" {
				r.Table = table
				r.FCodes = fcodes
			}
		case "INSERT":
			name := prsd.getTable(0)
			if prsd.isInsert(0, name) {
				r.HasView = true
			}
		case "(":
			r.DataLines++
		}

		if prsd[0].val[0] == '@' {
			prsd.stringBuilder(macrosSB, 0, len(prsd)-1)
		}

	}
	// TODO: Do checks to make sure stats are valid

	// check for macros
	r.macros, err = script.Read(strings.NewReader(macrosSB.String()))
	if err != nil {
		return err
	}

	for m := range r.macros.Macros().Range {
		if mm, ok := m.(*macro.Macro); ok {
			switch mm.Def.Name {
			case "CREATE":
				tableName := mm.ParametersParts.Index(0).String()
				if !strings.Contains(strings.ToUpper(tableName), strings.ToUpper(r.Stats.Table)) {
					return fmt.Errorf("CREATE table does not match table name: %s", tableName)
				}
				r.HasCreate = true
				// case "UPDATE_BATCH":
				// 	_, _ = mm.ParametersParts.FindSetter("JOB")
			}
		}

	}

	return nil
}
