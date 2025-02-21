package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dave/dst"
)

type grammarConfig struct {
	Grammars []grammarOutput `json:"grammars"`
}

type grammarOutput struct {
	Path                 string   `json:"path"`
	Files                []string `json:"files"`
	AlternateEntryPoints []string `json:"alternateEntryPoints"`
}

type ParserPatcher func(file *dst.File) error

func Parser(grammarFile string, debug bool, customPatcher ParserPatcher) (err error) {
	js, err := os.ReadFile(grammarFile)
	if err != nil {
		err = fmt.Errorf("failed reading grammar file %s: %w", grammarFile, err)
		return
	}
	var gc grammarConfig
	err = json.Unmarshal(js, &gc)

	if err != nil {
		err = fmt.Errorf("failed parsing grammar file %s: %w", grammarFile, err)
		return
	}

	root := filepath.Dir(grammarFile)
	for _, grammarOutput := range gc.Grammars {
		var grammar strings.Builder
		slog.Info("Generating parser", "path", grammarOutput.Path, slog.Bool("debug", debug))
		for _, f := range grammarOutput.Files {
			path := filepath.Join(root, f)
			var b []byte
			b, err = os.ReadFile(path)
			if err != nil {
				err = fmt.Errorf("failed reading grammar file %s: %w", path, err)
				return
			}
			slog.Info("Adding file", "path", f)
			grammar.WriteString(string(b))
			grammar.WriteString("\n\n")
		}

		args := []string{}
		if !debug {
			args = append(args, "-optimize-parser")
			//args = append(args, "-optimize-grammar")
		}
		if len(grammarOutput.AlternateEntryPoints) > 0 {
			args = append(args, "-alternate-entrypoints")
			args = append(args, strings.Join(grammarOutput.AlternateEntryPoints, ","))
		}
		pigeon := exec.Command("pigeon", args...)

		var stdin io.WriteCloser
		stdin, err = pigeon.StdinPipe()
		if err != nil {
			err = fmt.Errorf("failed getting pigeon stdin: %w", err)
			return
		}
		defer stdin.Close()

		var out bytes.Buffer

		pigeon.Stdout = &out
		pigeon.Stderr = os.Stderr

		if err = pigeon.Start(); err != nil {
			err = fmt.Errorf("failed starting pigeon: %w", err)
			return
		}

		var g string
		if !debug {
			var pattern = regexp.MustCompile("(// *)?(debug|debugPosition)\\([^\n]*\\)\n")
			g = pattern.ReplaceAllString(grammar.String(), "")
		} else {
			g = grammar.String()
		}

		_, err = io.WriteString(stdin, g)
		if err != nil {
			err = fmt.Errorf("failed writing grammar to pigeon: %w", err)
			return
		}
		err = stdin.Close()
		if err != nil {
			err = fmt.Errorf("failed closing pigeon stdin: %w", err)
			return
		}
		var grammarFile = grammarOutput.Path + ".peg"
		if err = pigeon.Wait(); err != nil {
			err = fmt.Errorf("failed running pigeon: %w", err)
			// Dump the composed peg file for debugging
			_ = os.WriteFile(grammarFile, []byte(g), os.ModeAppend|0644)
			return
		}
		var patched string
		patched, err = optimizeParser(out.String(), customPatcher)
		if err != nil {
			return
		}

		err = os.WriteFile(grammarOutput.Path, []byte(patched), os.ModeAppend|0644)
		if err != nil {
			err = fmt.Errorf("failed saving parser to %s: %w", grammarOutput.Path, err)
			return
		}

		os.Remove(grammarFile)
	}
	return
}
