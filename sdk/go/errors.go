package orchadyn

import (
	"errors"
	"net/http"
)

// Status-classification helpers over *APIError, so callers can branch on the
// failure class without inspecting StatusCode directly.

// IsUnauthenticated reports whether err is an *APIError with status 401.
func IsUnauthenticated(err error) bool { return hasStatus(err, http.StatusUnauthorized) }

// IsDenied reports whether err is an *APIError with status 403.
func IsDenied(err error) bool { return hasStatus(err, http.StatusForbidden) }

// IsNotFound reports whether err is an *APIError with status 404.
func IsNotFound(err error) bool { return hasStatus(err, http.StatusNotFound) }

// IsConflict reports whether err is an *APIError with status 409.
func IsConflict(err error) bool { return hasStatus(err, http.StatusConflict) }

// IsUnavailable reports whether err is an *APIError with status 503.
func IsUnavailable(err error) bool { return hasStatus(err, http.StatusServiceUnavailable) }

func hasStatus(err error, code int) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == code
}
