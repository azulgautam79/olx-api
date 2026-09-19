package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string

const (
	CodeInvalidID        Code = "invalid_id"
	CodeMalformedJSON    Code = "malformed_json"
	CodeUnAuthenticated  Code = "unauthenticated"
	CodeForbidden        Code = "forbidden"
	CodeNotFound         Code = "not_found"
	CodeMethodNotAllowed Code = "method_not_allowed"
	CodeConflict         Code = "conflict"
	CodePayloadTooLarge  Code = "payload_too_large"
	CodeValidationFailed Code = "validation_failed"
	CodeRateLimited      Code = "rate_limited"
	CodeInternalError    Code = "internal_error"
)

type errorEnvelope struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Field   string `json: "field", omitempty`
}

func Error(w http.ResponseWriter, status int, message string, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(errorEnvelope{Error: errorPayload{
		Code:    code,
		Message: message,
	}})
}

func ValidationError(w http.ResponseWriter, status int, message string, code Code, field string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(errorEnvelope{Error: errorPayload{
		Code:    code,
		Message: message,
		Field:   field,
	}})
}
