package silread

import (
	"context"
	"io"
	"reflect"
	"time"
)

// parser does the actual parsing of bytes into sil type
type parser struct {
	s      *scanner
	ctx    context.Context
	cancel context.CancelFunc
	buf    struct {
		pt part
		n  int // buffer size (max=1)
	}
}

type parsed []part

type part struct {
	tok token
	val string
}

// NewParser returns a new instance of Parser.
func newParser(r io.Reader) *parser {
	return &parser{s: newScanner(r)}
}

func (p *parser) decodeChan(dataChan any) *decoder {
	var d decoder

	readLines := int64(0)

	var i int

	channel := reflect.ValueOf(dataChan)
	channelType := reflect.TypeOf(dataChan).Elem()

	defer func() {
		channel.Close()
	}()

	for {
		pt := p.scan()

		d.p = append(d.p, *pt)
		if pt.tok == CRLF {
			readLines++

			ni := d.identifyLine(i)
			// if the new i matches the old i break out since processing failed
			if ni == i {
				break
			}
			i = 0
			d.p = d.p[:0]

			// if the view has been reached pop off anything from d.data and put on channel
			if d.view {
				for _, data := range d.data {
					if len(d.fieldMap) == 0 {
						ctI := reflect.New(channelType).Interface()
						d.makeFieldMap(ctI)
					}
					dataV := reflect.New(channelType).Elem()
					err := unmarshalValue(data, dataV, d.fieldMap)
					if err != nil {
						d.err = append(d.err, err)
					} else {
						channel.Send(dataV)
					}
				}
				d.data = d.data[:0]
			}
		}
		if pt.tok == EOF {
			break
		}
	}

	return &d
}

func (p *parser) decode() *decoder {
	var d decoder

	readLines := int64(0)

	var i int

	for {

		pt := p.scan()

		d.p = append(d.p, *pt)
		if pt.tok == CRLF {
			readLines++

			ni := d.identifyLine(i)
			// if the new i matches the old i break out since processing failed
			if ni == i {
				break
			}
			i = 0
			d.p = d.p[:0]
		}

		if pt.tok == EOF {
			break
		}

	}

	return &d
}

func (p *parser) parse() *parsed {
	var prsd parsed

	for {
		pt := p.scan()
		prsd = append(prsd, *pt)

		if pt.tok == EOF {
			break
		}
	}

	return &prsd
}

func (p *parser) scan() (pt *part) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return &p.buf.pt
	}

	// Otherwise read the next token from the scanner.
	pt = p.s.scan()

	// Save it to the buffer in case we unscan later.
	p.buf.pt = *pt

	return
}

func (p *parser) scanChan() chan *part {
	ch := make(chan *part, 1000)
	go func() {
		ticker := time.NewTicker(time.Microsecond)
		defer ticker.Stop()
		for {
			select {
			case <-p.ctx.Done():
				return
			case <-ticker.C:
				if p.buf.n != 0 {
					p.buf.n = 0
					ch <- &p.buf.pt
				}

				// Otherwise read the next token from the scanner.
				pt := p.s.scan()

				// Save it to the buffer in case we unscan later.
				p.buf.pt = *pt
				ch <- pt
			}
		}
	}()

	return ch
}
