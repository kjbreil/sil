package sil

// View is the data of the SIL file
// Name is the table name
// Data is an array of interfaces
type View struct {
	Name     string
	Required bool
	Action   string
	Data     []interface{}
}

func (v *View) action() string {
	// 	if action is not set default to chg
	if v.Action == "" {
		return v.Name + "_CHG"
	}

	return v.Name + "_" + v.Action
}
