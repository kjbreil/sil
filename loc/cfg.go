package loc

// CFG is the CFG_TAB definition
// F1000 CHAR(5),F1056 CHAR(4),F2845 CHAR(100),F2846 CHAR(16),F253 DATE(7),F940 INTEGER,F941 INTEGER,F1001 INTEGER,F1264 DATE(7),F2847 CHAR(2000)
type CFG struct {
	F1000 string `sil:"CHAR(5)" default:"001901"`
	F1056 string `sil:"CHAR(4)" default:"001"`
	F2845 string `sil:"CHAR(100)"`
	F2846 string `sil:"CHAR(16)" default:"GROC_LANE"`
	F253  string `sil:"DATE(7)" default:"NOW"`
	F940  int    `sil:"INTEGER" default:"999"`
	F941  int    `sil:"INTEGER" default:"999"`
	F1001 int    `sil:"INTEGER" default:"1"`
	F1264 string `sil:"DATE(7)" default:"NOW"`
	F2847 string `sil:"CHAR(2000)"`
}
