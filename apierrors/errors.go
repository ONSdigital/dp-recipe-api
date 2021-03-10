package apierrors

import (
	"errors"
	"fmt"
	"strconv"
)

// A list of error messages for Recipes API
var (
	ErrAddRecipeAlreadyExists            = errors.New("forbidden - recipe already exists")
	ErrAddUpdateRecipeBadRequest         = errors.New("failed to parse json body")
	ErrAuditActionAttemptedFailure       = errors.New("internal server error")
	ErrInvalidPositiveInteger            = errors.New("value is not a positive integer")
	ErrIndexOutOfRange                   = errors.New("index out of range")
	ErrInternalServer                    = errors.New("internal error")
	ErrInsertedObservationsInvalidSyntax = errors.New("inserted observation request parameter not an integer")
	ErrMissingJobProperties              = errors.New("missing job properties")
	ErrMissingParameters                 = errors.New("missing properties in JSON")
	ErrNoAuthHeader                      = errors.New("no authentication header provided")
	ErrNotFound                          = errors.New("not found")
	ErrRecipeNotFound                    = errors.New("recipe not found")
	ErrTooManyWildcards                  = errors.New("only one wildcard (*) is allowed as a value in selected query parameters")
	ErrUnableToParseJSON                 = errors.New("failed to parse json body")
	ErrUnableToReadMessage               = errors.New("failed to read message body")
	ErrUnauthorised                      = errors.New("unauthorised access to API")

	NotFoundMap = map[error]bool{
		ErrRecipeNotFound: true,
	}

	BadRequestMap = map[error]bool{
		ErrInsertedObservationsInvalidSyntax: true,
		ErrMissingJobProperties:              true,
		ErrMissingParameters:                 true,
		ErrInvalidPositiveInteger:            true,
		ErrUnableToParseJSON:                 true,
		ErrUnableToReadMessage:               true,
	}
)

// ErrorMaximumLimitReached creates an for the given limit
func ErrorMaximumLimitReached(m int) error {
	return fmt.Errorf("the maximum limit has been reached, the limit cannot be more than %s", strconv.Itoa(m))
}
