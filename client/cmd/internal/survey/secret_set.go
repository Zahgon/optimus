package survey

// SecretSetSurvey defines survey for setting secret
type SecretSetSurvey struct{}

// NewSecretSetSurvey initializes survey to set secret
func NewSecretSetSurvey() *SecretSetSurvey { _ = "STUB: not implemented"; return nil }

// AskToConfirmUpdate asks the user to confirm updating secret
func (*SecretSetSurvey) AskToConfirmUpdate() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
