package obsidian

import (
	"os"
	"path/filepath"
)

type Chapter struct {
	Name string

	Pages []Page
}

func ChapterFromDir(path string, expandPages bool) (Chapter, error) {
	chapter := Chapter{Name: filepath.Base(path)}

	if expandPages {
		entries, err := os.ReadDir(path)
		if err != nil {
			return chapter, err
		}

		for _, e := range entries {
			if e.IsDir() {
				continue
			}

			path := filepath.Join(path, e.Name())

			page, err := PageFromFile(path)
			if err != nil {
				return chapter, err
			}
			chapter.Pages = append(chapter.Pages, page)
		}
	}

	return chapter, nil
}
