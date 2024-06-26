//go:generate go run generate.go
//go:build generate

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type grammarConfig struct {
	Grammars []grammarOutput `json:grammars`
}

type grammarOutput struct {
	Path                 string   `json:path`
	Files                []string `json:files`
	AlternateEntryPoints []string `json:alternateEntryPoints`
}

var debugParser = false

func main() {
	js, err := os.ReadFile("grammar/grammar.json")
	if err != nil {
		slog.Error("failed reading grammar.json file", slog.Any("error", err))
		os.Exit(1)
		return
	}
	var gc grammarConfig
	err = json.Unmarshal(js, &gc)

	if err != nil {
		slog.Error("failed parsing grammar.json file", slog.Any("error", err))
		os.Exit(1)
		return
	}
	for _, grammarOutput := range gc.Grammars {
		var grammar strings.Builder
		for _, f := range grammarOutput.Files {
			path := filepath.Join("grammar", f)
			b, err := os.ReadFile(path)
			if err != nil {
				slog.Error("failed reading grammar file", slog.String("path", path), slog.Any("error", err))
				os.Exit(1)
				return
			}
			slog.Info("adding grammar file", "path", path)
			grammar.WriteString(string(b))
			grammar.WriteString("\n\n")
		}

		args := []string{}
		if !debugParser {
			args = append(args, "-optimize-parser")
			//args = append(args, "-optimize-grammar")
		}
		if len(grammarOutput.AlternateEntryPoints) > 0 {
			args = append(args, "-alternate-entrypoints")
			args = append(args, strings.Join(grammarOutput.AlternateEntryPoints, ","))
		}
		pigeon := exec.Command("pigeon", args...)

		stdin, err := pigeon.StdinPipe()
		if err != nil {
			slog.Error("failed getting pigeon stdin", slog.Any("error", err))
			os.Exit(1)
			return
		}
		defer stdin.Close()

		var out bytes.Buffer

		pigeon.Stdout = &out
		pigeon.Stderr = os.Stderr

		if err = pigeon.Start(); err != nil {
			slog.Error("failed starting pigeon", slog.Any("error", err))
			os.Exit(1)
			return
		}

		var g string
		if !debugParser {
			var pattern = regexp.MustCompile("(debug|debugPosition)\\([^\n]*\\)\n")
			g = pattern.ReplaceAllString(grammar.String(), "")
		} else {
			g = grammar.String()
		}

		_, err = io.WriteString(stdin, g)
		if err != nil {
			slog.Error("failed writing grammar to pigeon", slog.Any("error", err))
			os.Exit(1)
			return
		}
		err = stdin.Close()
		if err != nil {
			slog.Error("failed closing stdin", slog.Any("error", err))
			os.Exit(1)
			return
		}
		var grammarFile = grammarOutput.Path + ".peg"
		if err = pigeon.Wait(); err != nil {
			slog.Error("failed running pigeon", slog.Any("error", err))
			os.WriteFile(grammarFile, []byte(grammar.String()), os.ModeAppend|0644)
			os.Exit(1)
			return
		} else {
			parser := out.String()
			parser = strings.ReplaceAll(parser, "globalStore storeDict", "globalStore storeDict\n\tdelimitedBlockState delimitedBlockState\n\tparser *parser\n\ttableColumnsAttribute *asciidoc.TableColumnsAttribute")
			parser = strings.ReplaceAll(parser, "p.setOptions(opts)", "p.setOptions(opts)\n\tp.cur.parser = p")

			parser = strings.ReplaceAll(parser, "globalStore: make(storeDict),", "globalStore: make(storeDict),\n\t\t\tdelimitedBlockState: make(delimitedBlockState),")

			parser = strings.ReplaceAll(parser, "recoveryStack []map[string]any", "recoveryStack []map[string]any\n\toffset position")
			parser = strings.ReplaceAll(parser, "vals := make([]any, 0, len(seq.exprs))", "var vals []any")
			parser = strings.ReplaceAll(parser, "basicLatinChars [128]bool", "//basicLatinChars [128]bool")

			err = os.WriteFile(grammarOutput.Path, []byte(parser), os.ModeAppend|0644)
			if err != nil {
				slog.Error("failed saving asciidoc", slog.Any("error", err))
				os.Exit(1)
				return
			}
		}
		os.Remove(grammarFile)
	}
	os.Exit(0)
}
