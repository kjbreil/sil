package loc

// FgpriceBat is the FGPRICE_BAT definition
type FgpriceBat struct {
	UPCCode                  string  `sil:"F01,zeropad"`
	UPCCodeFormat            *int    `sil:"F07"`
	TargetIDentifier         string  `sil:"F1000" default:"PAL"`
	RecordStatus             int     `sil:"F1001" default:"1"`
	DiscountPrintdetailtrack *string `sil:"F109"`
	AutoCalculateStorePrice  *string `sil:"F1135"`
	HoldRetailPrice          *string `sil:"F119"`
	LikeCode                 *string `sil:"F122"`
	PriceLevel               *int    `sil:"F126"`
	ItemLinkCode             string  `sil:"F164"`
	AlternatePrice           *string `sil:"F166"`
	AlternatePriceQuantity   *string `sil:"F167"`
	ProhibitItemMarkdown     *string `sil:"F175"`
	MaintenanceOperatorLevel *int    `sil:"F1759"`
	NotInAdmissibleSpending  *string `sil:"F177"`
	SpareFieldPRI1           *string `sil:"F1799"`
	SpareFieldPRI2           *string `sil:"F1800"`
	AddingTotalizerfunction  *int    `sil:"F1802"`
	TimeIncluded             *string `sil:"F1803"`
	SpareFieldPRI6           *string `sil:"F1804"`
	CommissionDollars        *string `sil:"F1885"`
	MeasureSellPack          *string `sil:"F21"`
	Price1Description        *string `sil:"F2694"`
	Price2Description        *string `sil:"F2695"`
	Price3Description        *string `sil:"F2696"`
	PriceAudit               *string `sil:"F41"`
	ApplyItemDiscount        *string `sil:"F43"`
	BatchIDentifier          *string `sil:"F902"`
	BatchCreator             *string `sil:"F903"`
}
