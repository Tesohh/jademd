package obsidian

import (
	"bytes"
	"os"
	"path/filepath"
)

type Vault struct {
	VaultMetadata
	Html string

	Courses []Course
}

type VaultMetadata struct {
	Name         string
	Authors      string
	ContactEmail string
}

func VaultFromDir(path string, expandCourses bool, expandChapters bool, expandPages bool) (Vault, error) {
	vault := Vault{} // name comes from metadata

	if expandChapters {
		entries, err := os.ReadDir(path)
		if err != nil {
			return vault, err
		}

		for _, e := range entries {
			if !e.IsDir() {
				continue
			}

			path := filepath.Join(path, e.Name())

			course, err := CourseFromDir(path, expandChapters, expandPages)
			if err != nil {
				return vault, err
			}
			vault.Courses = append(vault.Courses, course)
		}
	}

	// parse metadata
	metadataFile, err := os.ReadFile(filepath.Join(path, "metadata.md"))
	if err != nil {
		return vault, err
	}
	metadata, content, err := ParseFrontMatter[VaultMetadata](metadataFile)
	if err != nil {
		return vault, err
	}

	vault.VaultMetadata = metadata
	if vault.Name == "" {
		vault.Name = "Unnamed vault"
	}

	var buf bytes.Buffer
	err = md.Convert(content, &buf)
	if err != nil {
		return vault, err
	}
	vault.Html = buf.String()

	return vault, nil
}
