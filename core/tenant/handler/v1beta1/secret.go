package v1beta1

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/core/tenant/dto"
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

const (
	metricSecretEvents               = "secret_events_total"
	secretEventsStatusRegistered     = "registered"
	secretEventsStatusUpdated        = "updated"
	secretEventsStatusDeleted        = "deleted"
	secretEventsStatusRegisterFailed = "register_failed"
	secretEventsStatusUpdateFailed   = "update_failed"
	secretEventsStatusDeleteFailed   = "delete_failed"
)

type SecretService interface {
	Save(ctx context.Context, projName tenant.ProjectName, nsName string, pts *tenant.PlainTextSecret) error
	Update(ctx context.Context, projName tenant.ProjectName, nsName string, pts *tenant.PlainTextSecret) error
	Delete(ctx context.Context, projName tenant.ProjectName, nsName string, secretName tenant.SecretName) error
	GetSecretsInfo(ctx context.Context, projName tenant.ProjectName) ([]*dto.SecretInfo, error)
}

type SecretHandler struct {
	l             log.Logger
	secretService SecretService

	pb.UnimplementedSecretServiceServer
}

func (sv *SecretHandler) RegisterSecret(ctx context.Context, req *pb.RegisterSecretRequest) (*pb.RegisterSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sv *SecretHandler) UpdateSecret(ctx context.Context, req *pb.UpdateSecretRequest) (*pb.UpdateSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sv *SecretHandler) ListSecrets(ctx context.Context, req *pb.ListSecretsRequest) (*pb.ListSecretsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sv *SecretHandler) DeleteSecret(ctx context.Context, req *pb.DeleteSecretRequest) (*pb.DeleteSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDecodedSecret(encodedString string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func NewSecretsHandler(l log.Logger, secretService SecretService) *SecretHandler {
	_ = "STUB: not implemented"
	return nil
}

func raiseSecretEventsMetric(projectName, namespaceName, state string) {
	_ = "STUB: not implemented"
	return
}
