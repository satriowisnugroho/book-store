package config

const (
	// ServiceFee is a service fee
	ServiceFee = 1000
	// UniqueConstraintViolationCode is the pgError code for unique constraint violation error
	UniqueConstraintViolationCode = "23505"
	// MinPasswordLen is the the minimum length of password
	MinPasswordLen = 5
	// AuthorizationHeader is a header for authorization
	AuthorizationHeader = "Authorization"
	// AuthorizationHeaderBearer is an authorization header format
	AuthorizationHeaderBearer = "Bearer"
)
