// Package orchadyn is the Apache-2.0 Go client for the public ORCHADYN
// Planning API.
//
// # Response decoding
//
// Methods take a caller-supplied `response any` target rather than returning a
// concrete type, so a caller decodes into exactly the shape it needs (the
// planning surface has many large, evolving response envelopes). Pass `nil` to
// discard the body.
//
// # Errors
//
// A non-2xx response is returned as *APIError (StatusCode + Body). Use the
// classification helpers IsUnauthenticated / IsDenied / IsNotFound / IsConflict
// / IsUnavailable to branch on the failure class.
package orchadyn
