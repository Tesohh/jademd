package obsidian

import (
	"bytes"
	"os"
	"path/filepath"
)

type Course struct {
	Name string
	Html string
	CourseMetadata

	Chapters []Chapter
}

type CourseMetadata struct {
	Color           string
	Authors         string
	ContactEmail    string
	DifficultyStars int
}

func CourseFromDir(path string, expandChapters bool, expandPages bool) (Course, error) {
	course := Course{Name: filepath.Base(path)}

	if expandChapters {
		entries, err := os.ReadDir(path)
		if err != nil {
			return course, err
		}

		for _, e := range entries {
			if !e.IsDir() {
				continue
			}

			path := filepath.Join(path, e.Name())

			chapter, err := ChapterFromDir(path, expandPages)
			if err != nil {
				return course, err
			}
			course.Chapters = append(course.Chapters, chapter)
		}
	}

	// parse metadata
	metadataFile, err := os.ReadFile(filepath.Join(path, "metadata.md"))
	if err != nil {
		return course, err
	}
	metadata, content, err := ParseFrontMatter[CourseMetadata](metadataFile)
	if err != nil {
		return course, err
	}

	course.CourseMetadata = metadata

	var buf bytes.Buffer
	err = md.Convert(content, &buf)
	if err != nil {
		return course, err
	}
	course.Html = buf.String()

	return course, nil
}
