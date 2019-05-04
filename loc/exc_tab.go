package loc

// ExcTab is the EXC_TAB definition
type ExcTab struct {
	ExceptCode       *int    `sil:"F1511"`
	ExceptDescriptor *string `sil:"F1512"`
	ExceptType       *string `sil:"F1513"`
	ShowInEJ         *string `sil:"F1897"`
}
