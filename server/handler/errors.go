package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrPublishPathNotSet = fiber.Error{
		Code:    http.StatusInternalServerError,
		Message: "JADE_PUBLISH_PATH not set by server",
	}
	ErrPublisherKeyNotSet = fiber.Error{
		Code:    http.StatusNotAcceptable,
		Message: "PublisherKey header not set or empty",
	}
	ErrPublisherKeyNotFound = fiber.Error{
		Code:    http.StatusUnauthorized,
		Message: "PublisherKey not found in Publishers, which means you are unauthorized",
	}
)
