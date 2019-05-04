package loc

// DelvTab is the DELV_TAB definition
type DelvTab struct {
	TargetIdentifier          string  `sil:"F1000"`
	OrderTime                 *string `sil:"F2655"`
	DeliveryTime              *string `sil:"F2656"`
	InvisibleOnPDA            *string `sil:"F2658"`
	DocumentSubType           *string `sil:"F2659"`
	DeliveryCreateXDaysFuture *int    `sil:"F2661"`
	DeliveryDescription       *string `sil:"F2662"`
	DeliveryCycle             *string `sil:"F2663"`
	DeliveryDayOfWeek         *int    `sil:"F2664"`
	OrderNoticeDays           *int    `sil:"F2665"`
	VendorId                  *string `sil:"F27"`
}
