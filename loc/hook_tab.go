package loc

// HookTab is the HOOK_TAB definition
type HookTab struct {
	HookHostProgram *string `sil:"F1700"`
	HookTask        *string `sil:"F1701"`
	HookDevice      *string `sil:"F1702"`
	HookTrigger     *string `sil:"F1703"`
	HookWindow      *string `sil:"F1704"`
	HookCommand     *string `sil:"F1705"`
}
