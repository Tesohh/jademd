# jade.md

`jade.md` is a solution to create a learning platform out of a [Obsidian.md](https://obsidian.md) vault.
Think of it like Obsidian Publish for courses, but self hosted and free.
It is designed so the writer can do everything inside obsidian and for it to feel natural.

It encompasses 3 things:

- Go Web server + a HTMX frontend to go with it
  - Docker compose configs to run the server, DB and (maybe) a Git server to track the vault in teams
- 99% Preconfigured obsidian vault
- Custom Obsidian plugin for `jade.md`

## Vault structure

```
Vault
    Course
        Chapter
            01 Welcome.md // every file must have a number in their name and
            02 Foobar.md  // a id in their frontmatter or it will be ignored.
            03 Barfoo.md
            metadata.md // data about the chapter
        metadata.md // data about the course

    Chapterless Course
        01 Welcome.md
        02 Foobar.md
        metadata.md

metadata.md // data about the vault
```
