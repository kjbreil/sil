package loc

// DelvTab is the DELV_TAB definition
type DelvTab struct {
	TargetIDentifier          string  `sil:"F1000" default:"PAL"`
	OrderTime                 *string `sil:"F2655"`
	DeliveryTime              *string `sil:"F2656"`
	InvisibleOnPDA            *string `sil:"F2658"`
	DocumentSubType           *string `sil:"F2659"`
	DeliveryCreateXDaysFuture *int    `sil:"F2661"`
	DeliveryDescription       *string `sil:"F2662"`
	DeliveryCycle             *string `sil:"F2663"`
	DeliveryDayOfWeek         *int    `sil:"F2664"`
	OrderNoticeDays           *int    `sil:"F2665"`
	VendorID                  *string `sil:"F27"`
}
