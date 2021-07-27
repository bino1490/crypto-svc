package entity

import (
	"errors"
)

var (
	ErrDefault       = errors.New("ErrDefault")       // Default to fallback
	ErrInvalidConfig = errors.New("ErrInvalidConfig") // Invalid configuration in config yml

	ErrInvalidInputAttr1 = errors.New("ErrInvalidInputAttr1") // The user input value for InputAttr1 is empty or invalid
	ErrInvalidInputAttr2 = errors.New("ErrInvalidInputAttr2") // The user input value for InputAttr2 is empty or invalid

	// Business errors
	ErrItemNotFound    = errors.New("ErrItemNotFound")    // The key found no records in Databse
	ErrMaxLimitReached = errors.New("ErrMaxLimitReached") // The user has reached the maximum allowed limit of items
	ErrItemExists      = errors.New("ErrItemExists")      // The key to insert already exists in Database

	// Infrastructure errors
	ErrDatabaseFailure = errors.New("ErrDatabaseFailure") // Critical Database failure
	// ErrReadyzFailure represents that the readiness check of
	// one or more downstream systems has failed.
	ErrReadyzFailure = errors.New("ErrReadyzFailure")

	// ErrHealthzFailure represents that the liveness check of
	// one or more downstream systems has failed.
	ErrHealthzFailure = errors.New("ErrHealthzFailure")

	// ErrInvalidInputItemId represents that the item identifier
	// sent in the request is missing or invalid.
	ErrInvalidInputItemId = errors.New("ErrInvalidInputItemId")

	// ErrInvalidAccessToken represents that the access token sent
	// in the request is missing or invalid or the service is unable
	// to decrypt as expected.
	ErrInvalidAccessToken = errors.New("ErrInvalidAccessToken")

	// ErrCryptoFailure represents that the service is unable to
	// decrypt the access token sent in the request.
	ErrCryptoFailure = errors.New("ErrCryptoFailure")

	// ErrInvalidInputPageNumber represents that the page number value
	// sent in the request is missing or invalid.
	ErrInvalidInputPageNumber = errors.New("ErrInvalidInputPageNumber")

	// ErrInvalidInputPageSize represents that the page size value
	// sent in the request is missing or invalid.
	ErrInvalidInputPageSize = errors.New("ErrInvalidInputPageSize")

	// ErrInvalidInputSortBy represents that sort value
	// sent in the request is missing or invalid.
	ErrInvalidInputSortBy = errors.New("ErrInvalidInputSortBy")

	// ErrInvalidInputSortOrder represents that sort order value
	// sent in the request is missing or invalid.
	ErrInvalidInputSortOrder = errors.New("ErrInvalidInputSortOrder")

	// ErrInvalidInputBody represents that the input request payload
	// is invalid.
	ErrInvalidInputBody = errors.New("ErrInvalidInputBody")

	// ErrInvalidClientId represents that the input client identifier
	// request parameter is invalid or missing.
	ErrInvalidClientId = errors.New("ErrInvalidClientId")
)
