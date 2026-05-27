package service

import (
	"context"

	"github.com/raystack/salt/log"

	"github.com/raystack/optimus/core/tenant"
	"github.com/raystack/optimus/core/tenant/dto"
)

const keyLength = 32

type SecretRepository interface {
	Save(ctx context.Context, secret *tenant.Secret) error
	Update(ctx context.Context, secret *tenant.Secret) error
	Get(ctx context.Context, projName tenant.ProjectName, nsName string, name tenant.SecretName) (*tenant.Secret, error)
	GetAll(ctx context.Context, projName tenant.ProjectName, nsName string) ([]*tenant.Secret, error)
	Delete(ctx context.Context, projName tenant.ProjectName, nsName string, name tenant.SecretName) error
	GetSecretsInfo(ctx context.Context, projName tenant.ProjectName) ([]*dto.SecretInfo, error)
}

type SecretService struct {
	appKey *[keyLength]byte
	repo   SecretRepository

	logger log.Logger
}

func (s SecretService) Save(ctx context.Context, projName tenant.ProjectName, nsName string, secret *tenant.PlainTextSecret) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SecretService) Update(ctx context.Context, projName tenant.ProjectName, nsName string, secret *tenant.PlainTextSecret) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SecretService) Get(ctx context.Context, projName tenant.ProjectName, namespaceName, name string) (*tenant.PlainTextSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s SecretService) GetAll(ctx context.Context, projName tenant.ProjectName, namespaceName string) ([]*tenant.PlainTextSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s SecretService) Delete(ctx context.Context, projName tenant.ProjectName, nsName string, name tenant.SecretName) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SecretService) GetSecretsInfo(ctx context.Context, projName tenant.ProjectName) ([]*dto.SecretInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSecretService(appKey *[32]byte, repo SecretRepository, logger log.Logger) *SecretService {
	_ = "STUB: not implemented"
	return nil
}
