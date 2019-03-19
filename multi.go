package sil

import "fmt"

// Optional returns an array of SIL's with only required and present elements
func (s *SIL) Optional() (sils []SIL, err error) {
	colmap := make(map[int][]interface{})

	for _, e := range s.View.Data {
		// get the fields and values
		fields, values := fieldValue(e)
		members, err := forFields(fields, values, true)
		if err != nil {
			return sils, err
		}

		colmap[len(members)] = append(colmap[len(members)], e)

	}

	for k, v := range colmap {
		newSil := *s
		newSil.Header.F902 = fmt.Sprintf("%08d", k)
		newSil.View.Data = v
		sils = append(sils, newSil)
	}

	return
}
