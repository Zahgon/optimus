package connection

import (
	"context"
)

type bearerAuthentication struct {
	Token string
}

func (a *bearerAuthentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*bearerAuthentication) RequireTransportSecurity() bool {
	_ = "STUB: not implemented"
	return false
}
