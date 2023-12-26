package cmd

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"

	"github.com/hasty/alchemy/cmd/compare"
	"github.com/hasty/alchemy/cmd/database"
	"github.com/hasty/alchemy/cmd/disco"
	"github.com/hasty/alchemy/cmd/dm"
	"github.com/hasty/alchemy/cmd/dump"
	"github.com/hasty/alchemy/cmd/format"
	"github.com/hasty/alchemy/cmd/zap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "alchemy",
	Short:        "",
	Long:         ``,
	SilenceUsage: true,
}

func Execute() {
	verbose, _ := rootCmd.Flags().GetBool("verbose")
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("dryrun", "d", false, "whether or not to actually output files")
	rootCmd.PersistentFlags().Bool("serial", false, "process files one-by-one")
	rootCmd.PersistentFlags().Bool("verbose", false, "display verbose information")
	rootCmd.PersistentFlags().StringSliceP("attribute", "a", []string{}, "attribute for pre-processing asciidoc; this flag can be provided more than once")

	rootCmd.AddCommand(format.Command)
	rootCmd.AddCommand(disco.Command)
	rootCmd.AddCommand(zap.Command)
	rootCmd.AddCommand(database.Command)
	rootCmd.AddCommand(compare.Command)
	rootCmd.AddCommand(conformanceCommand)
	rootCmd.AddCommand(dump.Command)
	rootCmd.AddCommand(dm.Command)
	rootCmd.AddCommand(versionCommand)
	rootCmd.AddCommand(xmlCommand)
}

var selfClosingTags = regexp.MustCompile("></[^>]+>")

var xmlCommand = &cobra.Command{
	Use:   "xml",
	Short: "test conformance values",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return cmd.Usage()
		}
		xb, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		var out bytes.Buffer
		xd := xml.NewDecoder(bytes.NewReader(xb))
		xe := xml.NewEncoder(&out)
		xe.Indent("", "    ")
		for {
			tok, _ := xd.Token()
			if tok == nil {
				break
			}
			switch tok := tok.(type) {
			case xml.StartElement:
				slog.Info("start", "start", tok)
			case xml.EndElement:
				slog.Info("end", "end", tok)
			case xml.Comment:
				slog.Info("commend", "comment", tok)
			case xml.ProcInst:
				slog.Info("proc", "proc", tok)
			case xml.Directive:
				slog.Info("directive", "direct", tok)
			case xml.CharData:
				slog.Info("cdata", "val", tok)
			default:
				slog.Info("unknown token", "token", tok)
			}
			xe.EncodeToken(tok)
		}
		xe.Flush()

		b := selfClosingTags.ReplaceAll(out.Bytes(), []byte("/>"))

		os.WriteFile(args[0], b, os.ModeAppend|0644)
		return nil
	},
}

type xmlDecoder interface {
	Token() (xml.Token, error)
}

type loggingDecoder struct {
	d xmlDecoder
}

func (le *loggingDecoder) Token() (xml.Token, error) {
	tok, err := le.d.Token()
	if err != nil {
		return tok, err
	}
	switch t := tok.(type) {
	case xml.StartElement:
		fmt.Fprintf(os.Stderr, "decoding start element %s\n ", t.Name.Local)
	case xml.EndElement:
		fmt.Fprintf(os.Stderr, "decoding end element %s\n ", t.Name.Local)
	case xml.CharData:
		fmt.Fprintf(os.Stderr, "decoding char data element %s\n ", string(t))
	case xml.Comment:
		fmt.Fprintf(os.Stderr, "decoding comment %s\n ", string(t))
	case xml.ProcInst:
		fmt.Fprintf(os.Stderr, "decoding proc inst\n")
	case xml.Directive:
		fmt.Fprintf(os.Stderr, "decoding directive\n")
	default:

	}
	return tok, err
}

type loggingEncoder struct {
	w io.Writer
	e *xml.Encoder
}

func (le *loggingEncoder) EncodeToken(t xml.Token) error {
	switch t := t.(type) {
	case xml.StartElement:
		fmt.Fprintf(os.Stderr, "encoding start element %s\n ", t.Name.Local)
	case xml.EndElement:
		fmt.Fprintf(os.Stderr, "encoding end element %s\n ", t.Name.Local)
	case xml.CharData:
		fmt.Fprintf(os.Stderr, "encoding char data element %s\n ", string(t))
	case xml.Comment:
		fmt.Fprintf(os.Stderr, "encoding comment %s\n ", string(t))
	case xml.ProcInst:
		fmt.Fprintf(os.Stderr, "encoding proc inst\n")
	case xml.Directive:
		fmt.Fprintf(os.Stderr, "encoding directive\n")
	default:

	}
	return le.e.EncodeToken(t)
}

func (le *loggingEncoder) Flush() error {
	return le.e.Flush()
}
