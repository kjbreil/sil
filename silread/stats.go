package silread

import "os"

type Stats struct {
	HasHeader bool
	Header    []string
	HasView   bool
	Table     string
	DataLines int
	FCodes    []string
}

func GetStats(filename string) (Stats, error) {
	var s Stats

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

			if pt.tok == EOF {
				break
			}
		}
	}

	return s, nil
}
