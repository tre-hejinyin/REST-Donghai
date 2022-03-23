package constants

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

const (
	CACHE_ACCOUNT_GROUP = "account:"
	EMAIL               = "email"
	HASH_TOKEN_KEY      = "token"
)
