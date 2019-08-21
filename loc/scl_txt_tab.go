package loc

// SclTxtTab is the SCL_TXT_TAB definition
type SclTxtTab struct {
	SubDepartmentCode        *int    `sil:"F04"`
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	Text9                    *string `sil:"F1517"`
	Text10                   *string `sil:"F1518"`
	Text11                   *string `sil:"F1519"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	IngredientDescriptor     *string `sil:"F1836"`
	Text2                    *string `sil:"F1837"`
	Text3                    *string `sil:"F1838"`
	Text4                    *string `sil:"F1839"`
	Text5                    *string `sil:"F1853"`
	Text6                    *string `sil:"F1854"`
	Text7                    *string `sil:"F1855"`
	Text8                    *string `sil:"F1856"`
	StoreResponsible         *string `sil:"F1964"`
	Text12                   *string `sil:"F1968"`
	Text13                   *string `sil:"F1969"`
	Text14                   *string `sil:"F1970"`
	Text15                   *string `sil:"F1971"`
	Text16                   *string `sil:"F1972"`
	LastChangeDate           string  `sil:"F253" default:"NOW"`
	TextNumber               *int    `sil:"F267"`
	IngredientTextFontSize   *int    `sil:"F2943"`
	Text1                    *string `sil:"F297"`
	BatchIDentifier          *string `sil:"F902"`
}
