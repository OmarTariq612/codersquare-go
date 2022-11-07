package types

type APIError string

var (
	ErrTokenExpired      = APIError("Token expired")
	ErrBadToken          = APIError("Bad token")
	ErrUserNotFound      = APIError("User not found")
	ErrPostNotFound      = APIError("Post not found")
	ErrLikeNotFound      = APIError("Like not found")
	ErrDuplicateEmail    = APIError("An account with this email already exists")
	ErrDuplicateUsername = APIError("An account with this username already exists")
	ErrDuplicatePostURL  = APIError("A post with this URL already exists")
	ErrDuplicateLike     = APIError("Duplicate like")
)
