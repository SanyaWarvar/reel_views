package constants

// Headers
const (
	AuthorizationHeader = "Authorization"
	RequestIdHeader     = "X-REQUEST-ID"
	RefreshHeader       = "X-REFRESH-TOKEN"
)

// Roles
const (
	ClientRole  = "CLIENT"
	AdminRole   = "ADMIN"
	SupportRole = "SUPPORT"
)

// Context
const (
	UserIdCtx    = "userId"
	UserRoleCtx  = "userRole"
	RequestIdCtx = "requestId"
	TraceIdCtx   = "traceId"
	SpanIdCtx    = "spanId"
	ApiNameCtx   = "apiName"
)

// Errors
const (
	BindBodyError string = "bind_body"
)
