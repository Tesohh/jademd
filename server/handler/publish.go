package handler

import (
	"archive/zip"
	"fmt"
	"io"
	"jademd/data"
	"os"
	"path/filepath"
	"strings"
	"time"

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
	defer formFile.Close()

	folder, err := zip.NewReader(formFile, header.Size)
	if err != nil {
		return err
	}

	if os.Getenv("JADE_PUBLISH_PATH") == "" {
		return &ErrPublishPathNotSet
	}

	// TODO: improe error handling, by canceling the whole folder being created to avoid half baked vaults being published
	dateStr := time.Now().Format(time.RFC3339)
	dateStr = strings.ReplaceAll(dateStr, "/", "-")
	dateStr = strings.ReplaceAll(dateStr, ":", "-")

	for _, zf := range folder.File {
		path := filepath.Join(os.Getenv("JADE_PUBLISH_PATH"), dateStr, zf.Name)

		if strings.Contains(path, "__MACOSX") || strings.Contains(path, ".DS_Store") {
			continue
		}

		fmt.Println(path)

		// is directory
		if zf.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			return err
		}

		unzippedArchive, err := zf.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(dstFile, unzippedArchive)
		if err != nil {
			return err
		}
	}

	return nil
}
