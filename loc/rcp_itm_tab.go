package loc

// RcpItmTab is the RCP_ITM_TAB definition
type RcpItmTab struct {
	UPCCode              string  `sil:"F01,zeropad"`
	RecipeID             *string `sil:"F2907"`
	RecipeIngredientID   *int    `sil:"F2908"`
	RecipeIngText        *string `sil:"F2920"`
	RecipeIngQuantityUPC *string `sil:"F2921"`
}
