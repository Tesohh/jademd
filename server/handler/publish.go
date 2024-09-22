package handler

import (
	"archive/zip"
	"jademd/data"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func Publish(c *fiber.Ctx) error {
	// Try to retrieve the publisher from the key that we can get from Headers
	key := c.Get("PublisherKey", "")
	if key == "" {
		return &ErrPublisherKeyNotSet
	}

	db := data.FromCtx(c)
	publisher := data.Publisher{}
	db.First(&publisher, "key = ?", key)

	if (publisher == data.Publisher{}) {
		return &ErrPublisherKeyNotFound
	}

	// Take the vault.zip
	header, err := c.FormFile("vault")
	if err != nil {
		return err
	}

	formFile, err := header.Open()
	if err != nil {
		return err
	}

	folder, err := zip.NewReader(formFile, header.Size)
	if err != nil {
		return err
	}

	if os.Getenv("JADE_PUBLISH_PATH") == "" {
		return &ErrPublishPathNotSet
	}

	for _, zf := range folder.File {
		// TODO: save files to disk...
		path := filepath.Join(os.Getenv("JADE_PUBLISH_PATH"), zf.Name)
		// mkdirall...
		// write file...
	}

	return nil
}
