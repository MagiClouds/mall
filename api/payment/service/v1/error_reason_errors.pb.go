// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsContentAlreadyExits(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_ALREADY_EXITS.String() && e.Code == 205
}

func ErrorContentAlreadyExits(format string, args ...interface{}) *errors.Error {
	return errors.New(205, ErrorReason_CONTENT_ALREADY_EXITS.String(), fmt.Sprintf(format, args...))
}
