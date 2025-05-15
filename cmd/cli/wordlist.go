package cli

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/matter/spec"
)

type Wordlist struct {
	spec.ParserOptions `embed:""`
}

func (f *Wordlist) Run(cc *Context) (err error) {

	_, pe := exec.LookPath("pyspelling")
	if pe != nil {
		slog.Info("Please install pyspelling before running wordlist")
		return
	}

	_, pe = exec.LookPath("aspell")
	if pe != nil {
		slog.Info("Please install aspell before running wordlist")
		return
	}

	slog.Info("Checking spelling...")
	cmd := exec.Command("pyspelling", "--config", ".spellcheck.yml")
	cmd.Dir = f.ParserOptions.Root

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()

	result := strings.TrimSpace(string(out))

	pattern := regexp.MustCompile(`Misspelled words:\n<context> ([^\n]+)\n------+\n((?:[^\n]+\n)+)------+`)

	matches := pattern.FindAllSubmatch(out, -1)
	if matches == nil {
		slog.Info("Result", "result", result)
		return
	}

	var words []string
	for _, m := range matches {
		wordlist := strings.Split(string(m[2]), "\n")
		for _, w := range wordlist {
			w = strings.TrimSpace(w)
			if w == "" {
				continue
			}
			words = append(words, w)
		}
	}

	slices.Sort(words)
	for _, w := range words {
		slog.Info("Adding to wordlist", "word", w)
	}

	wordlistPath := filepath.Join(f.ParserOptions.Root, ".github/.wordlist.txt")

	var wordlistFile []byte
	wordlistFile, err = os.ReadFile(wordlistPath)
	if err != nil {
		return
	}

	wordlist := string(wordlistFile)
	lines := strings.Split(wordlist, "\n")
	for _, word := range words {
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
