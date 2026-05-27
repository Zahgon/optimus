package tenant

const EntityTenant = "tenant"

type Tenant struct {
	projName ProjectName
	nsName   NamespaceName
}

func (t Tenant) ProjectName() ProjectName { _ = "STUB: not implemented"; return *new(ProjectName) }

func (t Tenant) NamespaceName() NamespaceName {
	_ = "STUB: not implemented"
	return *new(NamespaceName)
}

func (t Tenant) IsInvalid() bool { _ = "STUB: not implemented"; return false }

func NewTenant(projectName, namespaceName string) (Tenant, error) {
	_ = "STUB: not implemented"
	return *new(Tenant), nil
}

type WithDetails struct {
	project    Project
	namespace  Namespace
	secretsMap map[string]string
}

func NewTenantDetails(proj *Project, namespace *Namespace, secrets PlainTextSecrets) (*WithDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WithDetails) ToTenant() Tenant { _ = "STUB: not implemented"; return *new(Tenant) }

func (w *WithDetails) GetConfig(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// key not present in namespace, check project

func (w *WithDetails) GetConfigs() map[string]string { _ = "STUB: not implemented"; return nil }

func (w *WithDetails) Project() *Project { _ = "STUB: not implemented"; return nil }

func (w *WithDetails) Namespace() *Namespace { _ = "STUB: not implemented"; return nil }

func (w *WithDetails) SecretsMap() map[string]string { _ = "STUB: not implemented"; return nil }
