package tables

// OBJ is a sample of a main item
type OBJ struct {
	UPCCode             string  `sil:"F01,zeropad" json:"upc_code,omitempty"`
	UPCCodeFormat       *int    `sil:"F07" json:"upc_code_format,omitempty"`
	TargetIDentifier    string  `sil:"F1000" default:"PAL" json:"target_i_dentifier,omitempty"`
	RecordStatus        int     `sil:"F1001" default:"1" json:"record_status,omitempty"`
	LastChangeDate      string  `sil:"F253" default:"NOW" json:"last_change_date,omitempty"`
	ExpandedDescription *string `sil:"F29" json:"expanded_description,omitempty"`
	LongDescription     *string `sil:"F255" json:"long_description,omitempty"`
	CreatedByUser       *int    `sil:"F940" json:"created_by_user,omitempty"`
	ModifiedByUser      *int    `sil:"F941" json:"modified_by_user,omitempty"`
}
