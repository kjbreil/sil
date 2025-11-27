package silread

import (
	"context"
	"io"
	"reflect"
	"strings"
)

// parser does the actual parsing of bytes into sil type
type parser struct {
	s   *scanner
	buf struct {
		pt part
		n  int // buffer size (max=1)
	}
}

type parsed []part

func (prsd parsed) stringBuilder(sb *strings.Builder, s int, end int) {
	for i := s; i < end; i++ {
		sb.WriteString(prsd[i].val)
	}
	return
}

type part struct {
	tok token
	val string
}

// NewParser returns a new instance of Parser.
func newParser(r io.Reader) *parser {
	return &parser{s: newScanner(r)}
}

func (p *parser) parse() parsed {
	var prsd parsed
	for {
		pt := p.scan()

		prsd = append(prsd, *pt)
		if pt.tok == EOF {
			break
		}
	}
	return prsd
}

// decodeChan decodes data from the parser and sends it to a channel
// Uses context.Background() for backwards compatibility
func (p *parser) decodeChan(dataChan any) *decoder {
	return p.decodeChanWithContext(context.Background(), dataChan)
}

// decodeChanWithContext decodes data from the parser and sends it to a channel with context support
// The context can be used to cancel the operation
func (p *parser) decodeChanWithContext(ctx context.Context, dataChan any) *decoder {
	var d decoder

	readLines := int64(0)

	var i int

	channel := reflect.ValueOf(dataChan)
	channelType := reflect.TypeOf(dataChan).Elem()

	defer func() {
		channel.Close()
	}()

	for {
		// Check context at the start of each iteration
		select {
		case <-ctx.Done():
			d.err = append(d.err, ctx.Err())
			return &d
		default:
		}

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
					// Check context before processing each data item
					select {
					case <-ctx.Done():
						d.err = append(d.err, ctx.Err())
						return &d
					default:
					}

					if len(d.fieldMap) == 0 {
						ctI := reflect.New(channelType).Interface()
						d.makeFieldMap(ctI)
					}
					dataV := reflect.New(channelType).Elem()

					err := unmarshalValues(data, dataV, d.fieldMap)
					if err != nil {
						d.err = append(d.err, err)
					} else {
						// Use select for sending to allow context cancellation
						select {
						case <-ctx.Done():
							d.err = append(d.err, ctx.Err())
							return &d
						default:
							channel.Send(dataV)
						}
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

func (p *parser) advanceTo(tok token) {
	for p.scan().tok != tok {
	}
}
