package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateClientConfig(conf *ClientConfig) error {
	_ = "STUB: not implemented"
	// implement this
	return nil
}

// ... etc

func ValidateServerConfig(_ *ServerConfig) error {
	_ = "STUB: not implemented"
	// implement this
	return nil
}

func validateNamespaces(value interface{}) error { _ = "STUB: not implemented"; return nil }

// ozzo-validation helper for nested validation struct
// https://github.com/go-ozzo/ozzo-validation/issues/136
func nestedFields(target interface{}, fieldRules ...*validation.FieldRules) *validation.FieldRules {
	_ = "STUB: not implemented"
	return nil
}
