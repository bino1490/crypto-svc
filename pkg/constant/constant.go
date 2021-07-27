// Package constant contains the common literals used by
// the service.
package constant

const (

	// ApplicationJSON is the name of the request header that contains
	// the content type value.
	ApplicationJSON string = "application/json"

	// SuccessCode represents successful processing of the
	// incoming request.
	SuccessCode int = 0

	// FailureCode represents failure processing of the
	// incoming request.
	FailureCode int = -1

	// SuccessMessage represents successful processing
	// of the incoming request.
	SuccessMessage string = "Success"

	// FailureMessage represents failure processing
	// of the incoming request.
	FailureMessage string = "Failure"
)
