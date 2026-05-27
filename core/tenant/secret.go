package tenant

const (
	EntitySecret = "secret"

	SecretStorageKey    = "STORAGE"
	SecretSchedulerAuth = "SCHEDULER_AUTH"
	SecretNotifySlack   = "NOTIFY_SLACK"
)

type SecretName string

func SecretNameFrom(name string) (SecretName, error) {
	_ = "STUB: not implemented"
	return *new(SecretName), nil
}

func (sn SecretName) String() string { _ = "STUB: not implemented"; return "" }

type PlainTextSecret struct {
	name  SecretName
	value string
}

func NewPlainTextSecret(name, value string) (*PlainTextSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PlainTextSecret) Value() string { _ = "STUB: not implemented"; return "" }

func (p *PlainTextSecret) Name() SecretName { _ = "STUB: not implemented"; return *new(SecretName) }

type PlainTextSecrets []*PlainTextSecret

func (p PlainTextSecrets) ToMap() map[string]string { _ = "STUB: not implemented"; return nil }

type Secret struct {
	name         SecretName
	encodedValue string

	projName      ProjectName
	namespaceName string
}

func (s *Secret) Name() SecretName { _ = "STUB: not implemented"; return *new(SecretName) }

func (s *Secret) EncodedValue() string { _ = "STUB: not implemented"; return "" }

func (s *Secret) ProjectName() ProjectName { _ = "STUB: not implemented"; return *new(ProjectName) }

func (s *Secret) NamespaceName() string { _ = "STUB: not implemented"; return "" }

func NewSecret(name, encodedValue string, projName ProjectName, nsName string) (*Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
