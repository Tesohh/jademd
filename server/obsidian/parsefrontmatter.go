package obsidian

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

func ParseFrontMatter[T any](markdown []byte) (T, []byte, error) {
	lines := bytes.Split(markdown, []byte("\n"))

	frontMatterEnd := -1
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if bytes.Compare(line, []byte("---")) == 0 {
			frontMatterEnd = i
			break
		}
	}

	var metadata T
	var content []byte

	hasFrontMatter := bytes.Compare(lines[0], []byte("---")) == 0 && frontMatterEnd != -1

	if hasFrontMatter {
		frontmatter := bytes.Join(lines[1:frontMatterEnd], []byte("\n"))
		fmt.Println(string(frontmatter), "END")

		err := yaml.Unmarshal(frontmatter, &metadata)
		if err != nil {
			return metadata, nil, err
		}
	}

	if hasFrontMatter {
		content = bytes.Join(lines[frontMatterEnd+1:], []byte("\n"))
	} else {
		content = markdown
	}

	return metadata, content, nil
}
