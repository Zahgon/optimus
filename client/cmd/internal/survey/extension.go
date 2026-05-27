package survey

// ExtensionSurvey defines survey for extension
type ExtensionSurvey struct{}

// NewExtensionSurvey initializes extension survey
func NewExtensionSurvey() *ExtensionSurvey {
	_ = "STUB: not implemented"

	// AskConfirmClean asks the user to confirm clean
	return nil
}

func (*ExtensionSurvey) AskConfirmClean() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AskConfirmUninstall asks the user to confirm uninstallation
func (*ExtensionSurvey) AskConfirmUninstall(commandName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
