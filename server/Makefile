dev:
	@templ generate
	@go run .

templ:
	@templ generate -watch -proxy=http://localhost:8080 -cmd="./.airtempl"
