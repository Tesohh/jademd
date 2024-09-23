package handler

import (
	"archive/zip"
	"fmt"
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
		fmt.Println(path)

		// is directory
		// if zf.FileInfo().IsDir() {
		// 	os.MkdirAll(path, os.ModePerm)
		// 	fmt.Println("IsDir making directory", path)
		// 	continue
		// }

		//at this point we are only dealing with files
		err = os.MkdirAll(path, os.ModePerm)
		fmt.Println("making directory", path)
		if err != nil {
			return err
		}

		unzippedFile, err := zf.Open()
		if err != nil {
			return err
		}
		defer unzippedFile.Close()

		b := make([]byte, 0)
		_, err = unzippedFile.Read(b)
		if err != nil {
			return err
		}
		defer unzippedFile.Close()

		fmt.Println("writing FILE", path)
		err = os.WriteFile(path, b, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
