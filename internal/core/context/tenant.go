package context_keys

type contextKey string

func (k contextKey) String() string {
	return "context key " + string(k)
}

const (
	TenantIDKey contextKey = "tenant_id"
	ServiceKey  contextKey = "service"
)
