package loc

// ObjTab is the OBJ_TAB definition
type ObjTab struct {
	UPCCode                                string  `sil:"F01,zeropad" json:"upc_code,omitempty"`
	UPCCodeFormat                          *int    `sil:"F07" json:"upc_code_format,omitempty"`
	TargetIDentifier                       string  `sil:"F1000" default:"PAL" json:"target_i_dentifier,omitempty"`
	RecordStatus                           int     `sil:"F1001" default:"1" json:"record_status,omitempty"`
	SizeCubic                              *string `sil:"F1002" json:"size_cubic,omitempty"`
	ContainerType                          *int    `sil:"F1004" json:"container_type,omitempty"`
	MeasurementSystem                      *int    `sil:"F11" json:"measurement_system,omitempty"`
	ManufacturerID                         *string `sil:"F1118" json:"manufacturer_id,omitempty"`
	GraphicFile                            *string `sil:"F1119" json:"graphic_file,omitempty"`
	OperatorResponsible                    *int    `sil:"F1168" json:"operator_responsible,omitempty"`
	SizeHeight                             *string `sil:"F12" json:"size_height,omitempty"`
	SizeWidth                              *string `sil:"F13" json:"size_width,omitempty"`
	SizeDepth                              *string `sil:"F14" json:"size_depth,omitempty"`
	BrandDescription                       *string `sil:"F155" json:"brand_description,omitempty"`
	FamilyCode                             *int    `sil:"F16" json:"family_code,omitempty"`
	ShippingPieceCount                     *string `sil:"F1699" json:"shipping_piece_count,omitempty"`
	CategoryCode                           *int    `sil:"F17" json:"category_code,omitempty"`
	BestBeforeDays                         *int    `sil:"F1736" json:"best_before_days,omitempty"`
	MedicalProductCode                     *string `sil:"F1737" json:"medical_product_code,omitempty"`
	NDCDINNumber                           *string `sil:"F1738" json:"ndcdin_number,omitempty"`
	MeasureWeightOrVolume                  *string `sil:"F1744" json:"measure_weight_or_volume,omitempty"`
	MaintenanceOperatorLevel               *int    `sil:"F1759" json:"maintenance_operator_level,omitempty"`
	SpareFieldOBJ1                         *string `sil:"F1781" json:"spare_field_obj_1,omitempty"`
	SpareFieldOBJ2                         *string `sil:"F1782" json:"spare_field_obj_2,omitempty"`
	NutritionIndex                         *int    `sil:"F1783" json:"nutrition_index,omitempty"`
	SpareFieldOBJ4                         *int    `sil:"F1784" json:"spare_field_obj_4,omitempty"`
	ReportCode                             *int    `sil:"F18" json:"report_code,omitempty"`
	ManufactureCode                        *string `sil:"F180" json:"manufacture_code,omitempty"`
	ReportingDepartment                    *string `sil:"F193" json:"reporting_department,omitempty"`
	AltBrandDesc                           *string `sil:"F1939" json:"alt_brand_desc,omitempty"`
	AltExpandedDesc                        *string `sil:"F1940" json:"alt_expanded_desc,omitempty"`
	AltSizeDesc                            *string `sil:"F1941" json:"alt_size_desc,omitempty"`
	AltLongDesc                            *string `sil:"F1942" json:"alt_long_desc,omitempty"`
	Classification                         *string `sil:"F1957" json:"classification,omitempty"`
	TargetCustomerType                     *string `sil:"F1958" json:"target_customer_type,omitempty"`
	TargetStoreType                        *string `sil:"F1959" json:"target_store_type,omitempty"`
	HandlingType                           *string `sil:"F1960" json:"handling_type,omitempty"`
	MarketingJustification                 *string `sil:"F1962" json:"marketing_justification,omitempty"`
	StoreResponsible                       *string `sil:"F1964" json:"store_responsible,omitempty"`
	MeasureSellPack                        *string `sil:"F21" json:"measure_sell_pack,omitempty"`
	PrintLabelOnNextBatch                  *string `sil:"F2119" json:"print_label_on_next_batch,omitempty"`
	DUNSNumberPlusSuffix                   *string `sil:"F213" json:"duns_number_plus_suffix,omitempty"`
	AliasCode                              *string `sil:"F214" json:"alias_code,omitempty"`
	AliasCodeFormat                        *int    `sil:"F215" json:"alias_code_format,omitempty"`
	ComparableSizeUnitOfMeasureDescription *string `sil:"F218" json:"comparable_size_unit_of_measure_description,omitempty"`
	SizeDescription                        *string `sil:"F22" json:"size_description,omitempty"`
	MeasureDescription                     *string `sil:"F23" json:"measure_description,omitempty"`
	LastChangeDate                         string  `sil:"F253" default:"NOW" json:"last_change_date,omitempty"`
	LongDescription                        *string `sil:"F255" json:"long_description,omitempty"`
	ItemSubstitutionPolicy                 *string `sil:"F2600" json:"item_substitution_policy,omitempty"`
	CompetitiveCode                        *string `sil:"F2693" json:"competitive_code,omitempty"`
	WeightNet                              *string `sil:"F270" json:"weight_net,omitempty"`
	InterDeptCode                          *string `sil:"F2789" json:"inter_dept_code,omitempty"`
	ExpandedDescription                    *string `sil:"F29" json:"expanded_description,omitempty"`
	NACSCode                               *int    `sil:"F2931" json:"nacs_code,omitempty"`
	BatchIDentifier                        *string `sil:"F902" json:"batch_i_dentifier,omitempty"`
	AccountCode                            *string `sil:"F93" json:"account_code,omitempty"`
	CreatedByUser                          *int    `sil:"F940" json:"created_by_user,omitempty"`
	ModifiedByUser                         *int    `sil:"F941" json:"modified_by_user,omitempty"`
}
