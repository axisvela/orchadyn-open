package orchadyn

import (
	"errors"
	"net/http"
	"testing"
)

func TestStatusClassification(t *testing.T) {
	unauth := &APIError{StatusCode: http.StatusUnauthorized}
	denied := &APIError{StatusCode: http.StatusForbidden}
	notFound := &APIError{StatusCode: http.StatusNotFound}
	wrapped := errors.New("wrapped")

	if !IsUnauthenticated(unauth) || IsUnauthenticated(denied) {
		t.Error("IsUnauthenticated misclassified")
	}
	if !IsDenied(denied) || IsDenied(unauth) {
		t.Error("IsDenied misclassified")
	}
	if !IsNotFound(notFound) || IsNotFound(unauth) {
		t.Error("IsNotFound misclassified")
	}
	if IsUnauthenticated(wrapped) {
		t.Error("non-APIError misclassified")
	}
}
