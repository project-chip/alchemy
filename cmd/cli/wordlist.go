package cli

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

type Wordlist struct {
	spec.ParserOptions `embed:""`

	Words []string `arg:"" help:"Paths of AsciiDoc files to format" required:""`
}

func (f *Wordlist) Run(cc *Context) (err error) {

	wordlistPath := filepath.Join(f.ParserOptions.Root, ".github/.wordlist.txt")

	var wordlistFile []byte
	wordlistFile, err = os.ReadFile(wordlistPath)
	if err != nil {
		return
	}

	wordlist := string(wordlistFile)
	lines := strings.Split(wordlist, "\n")
	for _, word := range f.Words {
		lines = insertWord(lines, word)
	}
	wordlist = strings.Join(lines, "\n")

	err = os.WriteFile(wordlistPath, []byte(wordlist), 0644)

	return
}

func insertWord(lines []string, word string) []string {
	lcWord := strings.ToLower(word)
	for i, l := range lines {
		l = strings.ToLower(strings.TrimSpace(l))
		switch strings.Compare(lcWord, l) {
		case -1:
			lines = slices.Insert(lines, i, word)
			return lines
		case 0:
			return lines
		case 1:
			continue
		}
	}
	lines = append(lines, word)
	return lines
}
