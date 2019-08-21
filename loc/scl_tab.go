package loc

// SclTab is the SCL_TAB definition
type SclTab struct {
	UPCCode                      string  `sil:"F01,zeropad"`
	TargetIDentifier             string  `sil:"F1000" default:"PAL"`
	RecordStatus                 int     `sil:"F1001" default:"1"`
	MaintenanceOperatorLevel     *int    `sil:"F1759"`
	ScaleSpecificFlag            *string `sil:"F1840"`
	ScaleDescriptor2             *string `sil:"F1952"`
	StoreResponsible             *string `sil:"F1964"`
	LastChangeDate               string  `sil:"F253" default:"NOW"`
	ProductDescription           *string `sil:"F256"`
	PriceModifier                *string `sil:"F257"`
	TareA                        *string `sil:"F258"`
	ScaleDescriptor3             *string `sil:"F2581"`
	ScaleDescriptor4             *string `sil:"F2582"`
	ScaleDescription1FontSize    *int    `sil:"F2583"`
	ScaleDescription2FontSize    *int    `sil:"F2584"`
	ScaleDescription3FontSize    *int    `sil:"F2585"`
	ScaleDescription4FontSize    *int    `sil:"F2586"`
	TareB                        *string `sil:"F259"`
	TareC                        *string `sil:"F260"`
	ExceptionPrice               *string `sil:"F261"`
	MessageNumber                *int    `sil:"F262"`
	Class                        *int    `sil:"F263"`
	ProductLife                  *int    `sil:"F264"`
	ActionNumber                 *string `sil:"F265"`
	SortNumber                   *string `sil:"F266"`
	TextNumber                   *int    `sil:"F267"`
	TemplateNumber               *string `sil:"F268"`
	PriceLookUpType              *string `sil:"F269"`
	WeightNet                    *string `sil:"F270"`
	LabelForm1                   *string `sil:"F271"`
	LabelForm2                   *string `sil:"F272"`
	LogoID                       *string `sil:"F273"`
	PrinterCode1                 *string `sil:"F274"`
	PrinterCode2                 *string `sil:"F275"`
	GraphicID                    *string `sil:"F2792"`
	COOLTextNumber               *int    `sil:"F2793"`
	COOLPreTextNumber            *int    `sil:"F2796"`
	COOLShortListNumber          *int    `sil:"F2797"`
	COOLClassNumber              *int    `sil:"F2799"`
	COOLForceSelection           *string `sil:"F2800"`
	COOLTrackingNumber           *string `sil:"F2801"`
	AllergenTextNumber           *int    `sil:"F2939"`
	CookingInstructionTextNumber *int    `sil:"F2940"`
	UserDefined1TextNumber       *int    `sil:"F2941"`
	UserDefined2TextNumber       *int    `sil:"F2942"`
	ForceShelfLifeEntry          *string `sil:"F2944"`
	ForceUsedByEntry             *string `sil:"F2945"`
	LabelRotation1               *int    `sil:"F2946"`
	LabelRotation2               *int    `sil:"F2947"`
	ShelfLifeType                *int    `sil:"F2948"`
	BarcodeFormat                *int    `sil:"F2949"`
	UsedByNumberOfDays           *int    `sil:"F2950"`
	StorageInstructionNumber     *int    `sil:"F2952"`
	BatchIDentifier              *string `sil:"F902"`
}
