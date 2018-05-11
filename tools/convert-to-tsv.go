package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
)

type bookRecord struct {
	Title   string
	Author  string
	Year    string
	Place   string
	Country string
	URL     string
}

func main() {

	var (
		scanner        = bufio.NewScanner(os.Stdin)
		curRecord      *bookRecord
		currentCountry = ""
		writer         = csv.NewWriter(os.Stdout)
	)

	writer.Comma = '\t'

	row := []string{
		"Author",
		"Title",
		"Year",
		"Language",
		"Publisher",
		"Place",
		"Country",
		"URL",
	}
	writer.Write(row)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "# ") {
			continue
		} else if strings.HasPrefix(line, "## ") {
			currentCountry = line[3:]
		} else if strings.HasPrefix(line, "### ") {
			curRecord.write(writer)
			curRecord = &bookRecord{
				Title:   line[4:],
				Country: currentCountry,
			}
		} else if strings.HasPrefix(line, "* ") {
			curRecord.parseField(line[2:])
		}
	}

	curRecord.write(writer)
	writer.Flush()
}

func (r *bookRecord) write(writer *csv.Writer) {
	if r == nil {
		return
	}

	row := []string{
		r.Author,
		r.Title,
		r.Year,
		"TODO: Language?",
		"TODO: Publisher?",
		r.Place,
		r.Country,
		r.URL,
	}
	writer.Write(row)
}

func (r *bookRecord) parseField(line string) {
	fields := strings.SplitN(line, ": ", 2)
	if len(fields) < 2 {
		return
	}
	switch fields[0] {
	case "Author":
		r.Author = fields[1]
	case "Place":
		r.Place = fields[1]
	case "Year":
		r.Year = fields[1]
	case "URL":
		r.URL = fields[1]
	}
}
