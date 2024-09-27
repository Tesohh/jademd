package obsidian

import (
	"bytes"
	"os"
	"path/filepath"
)

type Page struct {
	Name string
	Id   string
	Html string
}

func PageFromFile(path string) (Page, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Page{}, err
	}

	return PageFromMd(filepath.Base(path), b)
}

func PageFromMd(name string, markdown []byte) (Page, error) {
	page := Page{Name: name}

	metadata, content, err := ParseFrontMatter[struct{ Id string }]([]byte(markdown))
	if err != nil {
		return page, err
	}

	page.Id = metadata.Id

	var buf bytes.Buffer
	err = md.Convert(content, &buf)
	if err != nil {
		return page, err
	}

	page.Html = buf.String()
	return page, nil
}
