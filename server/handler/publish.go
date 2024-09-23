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

// "http://localhost:8080 TUBARAO"

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

	dateStr := time.Now().Format(time.RFC3339)
	dateStr = strings.ReplaceAll(dateStr, "/", "-")
	dateStr = strings.ReplaceAll(dateStr, ":", "-")

	vaultPath := filepath.Join(os.Getenv("JADE_PUBLISH_PATH"), dateStr)
	fmt.Printf("received %s files\n", len(folder.File))
	var unzipErr error
	for _, zf := range folder.File {
		path := filepath.Join(vaultPath, zf.Name)

		if strings.Contains(path, "__MACOSX") || strings.Contains(path, ".DS_Store") {
			continue
		}

		fmt.Println(path)

		// is directory
		if zf.FileInfo().IsDir() {
			unzipErr = os.MkdirAll(path, os.ModePerm)
			if unzipErr != nil {
				break
			}
			continue
		}

		dstFile, unzipErr := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if unzipErr != nil {
			break
		}

		unzippedArchive, unzipErr := zf.Open()
		if unzipErr != nil {
			break
		}

		_, unzipErr = io.Copy(dstFile, unzippedArchive)
		if unzipErr != nil {
			break
		}
	}

	if unzipErr != nil {
		fmt.Println("Aborting operation due to %s", err.Error())
		err := os.RemoveAll(vaultPath)
		if err != nil {
			return err
		}

		return fmt.Errorf("%w (operation aborted)", unzipErr)
	}

	return nil
}
