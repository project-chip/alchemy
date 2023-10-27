package cmd

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/matterfmt/disco"
)

type Option func(o interface{}) error

func DryRun(dry bool) Option {
	return func(o interface{}) error {
		if fp, ok := o.(fileProcessor); ok {
			fp.setDryRun(dry)
			return nil
		}
		return fmt.Errorf("invalid option for %T", o)
	}
}

func Serial(serial bool) Option {
	return func(o interface{}) error {
		if fp, ok := o.(fileProcessor); ok {
			fp.setSerial(serial)
			return nil
		}
		return fmt.Errorf("invalid option for %T", o)
	}
}

func Disco(opt disco.Option) Option {
	return func(o interface{}) error {
		switch v := o.(type) {
		case *discoBall:
			v.options = append(v.options, opt)
			return nil
		default:
			return fmt.Errorf("invalid option for %T", o)
		}
	}
}

func AsciiAttributes(attributes []string) Option {
	return func(o interface{}) error {
		as, ok := o.(asciiSettings)
		if !ok {
			return fmt.Errorf("invalid option for %T", o)
		}
		if len(attributes) == 0 {
			return nil
		}
		for _, a := range attributes {
			if len(a) == 0 {
				continue
			}
			for _, set := range strings.Split(a, ",") {
				as.addSetting(configuration.WithAttribute(strings.TrimSpace(set), true))
			}
		}

		return nil
	}
}

func DumpAscii(d bool) Option {
	return func(o interface{}) error {
		switch v := o.(type) {
		case *dumper:
			v.dumpAscii = d
			return nil
		default:
			return fmt.Errorf("invalid option for %T", o)
		}
	}
}
