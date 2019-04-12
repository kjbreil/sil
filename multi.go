package sil

import "github.com/kjbreil/sil/loc"

// Multi holds an array of SIL's and methods to work with them
type Multi map[string]*SIL

// Marshal creates the SIL structure from the information in the Multi
func (m Multi) Marshal() (data []byte, err error) {

	for _, s := range m {
		b, err := s.Marshal()
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
	m[name] = Make(name, loc.OBJ{})
}

// Append data to the view
func (m Multi) AppendView(name string, data interface{}) {
	// should check if the data is the correct type
	m[name].View.Data = append(m[name].View.Data, data)
}
