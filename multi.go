package sil

import (
	"math/rand"
	"time"
)

// Multi holds an array of SIL's and methods to work with them
type Multi map[string]*SIL

// Marshal creates the SIL structure from the information in the Multi
func (m Multi) Marshal() (data []byte, err error) {
	// assign a local prefix to the Multi to group all the batches
	rand.Seed(time.Now().UnixNano())
	prefix := rand.Intn(100)

	for _, s := range m {
		// assign that prefix to the local SIL
		s.prefix = prefix

		b, err := s.Marshal(false)
		if err != nil {
			return data, err
		}
		data = append(data, b...)
	}
	return data, nil
}

// Make creates a SIL object within the Multi type
func (m Multi) Make(name string, definition interface{}) {
	// should check if it exists yet because this will overwrite a previously
	// made SIL
	m[name] = Make(name, definition)
}

// AppendView appends data to the view
func (m Multi) AppendView(name string, data interface{}) {
	// should check if the data is the correct type
	m[name].View.Data = append(m[name].View.Data, data)
}

// SetHeaders sets all the headers in a sil file to the same name
func (m Multi) SetHeaders(name string) {
	for _, s := range m {
		s.Header.Description = name
	}
}
