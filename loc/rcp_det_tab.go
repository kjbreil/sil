package loc

// RcpDetTab is the RCP_DET_TAB definition
type RcpDetTab struct {
	UPCCode               string  `sil:"F01,zeropad"`
	RecipeID              *string `sil:"F2907"`
	RecipeTitle           *string `sil:"F2909"`
	RecipeDescription     *string `sil:"F2910"`
	RecipeServings        *string `sil:"F2911"`
	RecipePreparation     *string `sil:"F2912"`
	RecipeNotes           *string `sil:"F2913"`
	RecipeNutrition       *string `sil:"F2914"`
	RecipeDisclaimer      *string `sil:"F2915"`
	RecipeImageLink       *string `sil:"F2916"`
	RecipeVideoLink       *string `sil:"F2917"`
	RecipeClassification  *string `sil:"F2918"`
	RecipeStatus          *string `sil:"F2919"`
	RecipeTimePreparation *string `sil:"F2922"`
	RecipeTimeTotal       *string `sil:"F2923"`
	RecipeTypeDish        *string `sil:"F2924"`
	RecipeDifficultyLevel *string `sil:"F2925"`
}
