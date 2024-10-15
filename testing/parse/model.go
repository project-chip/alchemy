package parse

import (
	"fmt"
	"log/slog"

	"github.com/goccy/go-yaml"
)

type Test struct {
	Path string `yaml:"-"`

	Name string   `yaml:"name,omitempty"`
	PICS []string `yaml:"PICS,omitempty"`

	Config TestConfig  `yaml:"config,omitempty"`
	Tests  []*TestStep `yaml:"tests,omitempty"`

	Extras yaml.MapSlice
}

func (t *Test) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	slog.Info("Unmarshalling!")
	var yt yaml.MapSlice
	if err = unmarshal(&yt); err != nil {
		return err
	}
	t.Name, yt = extractValue[string](yt, "name")
	t.PICS, yt = extractArray[string](yt, "PICS")

	var config yaml.MapSlice
	config, yt = extractValue[yaml.MapSlice](yt, "config")
	if config != nil {
		err = t.Config.UnmarshalMap(config)
		if err != nil {
			return
		}
	}
	var tests []any
	tests, yt = extractValue[[]any](yt, "tests")
	for _, test := range tests {
		switch test := test.(type) {
		case yaml.MapSlice:
			ts := TestStep{Parent: t}
			err = ts.UnmarshalMap(test)
			if err != nil {
				return
			}
			t.Tests = append(t.Tests, &ts)
		default:
			slog.Info("unknown test type!", slog.String("val", fmt.Sprintf("%T", test)))
		}
	}
	if len(yt) > 0 {
		t.Extras = make(yaml.MapSlice, len(yt))
		copy(t.Extras, yt)
	}

	return nil
}

type TestConfig struct {
	NodeID                 uint64          `yaml:"nodeId,omitempty"`
	Cluster                string          `yaml:"cluster,omitempty"`
	Endpoint               uint64          `yaml:"endpoint,omitempty"`
	Timeout                uint64          `yaml:"timeout,omitempty"`
	CatalogVendorId        TestConfigValue `yaml:"catalogVendorId,omitempty"`
	ApplicationId          TestConfigValue `yaml:"applicationId,omitempty"`
	Payload                any             `yaml:"payload,omitempty"`
	Discriminator          TestConfigValue `yaml:"discriminator,omitempty"`
	WaitAfterCommissioning TestConfigValue `yaml:"waitAfterCommissioning,omitempty"`
	PakeVerifier           TestConfigValue `yaml:"PakeVerifier,omitempty"`

	Extras yaml.MapSlice
}

func (tc *TestConfig) UnmarshalMap(c yaml.MapSlice) error {
	tc.NodeID, c = extractValue[uint64](c, "nodeId")
	tc.Cluster, c = extractValue[string](c, "cluster")
	tc.Endpoint, c = extractValue[uint64](c, "endpoint")
	tc.Timeout, c = extractValue[uint64](c, "timeout")
	tc.CatalogVendorId, c = extractValue[TestConfigValue](c, "catalogVendorId")
	tc.ApplicationId, c = extractValue[TestConfigValue](c, "applicationId")
	tc.Discriminator, c = extractValue[TestConfigValue](c, "discriminator")
	tc.WaitAfterCommissioning, c = extractValue[TestConfigValue](c, "waitAfterCommissioning")

	if len(c) > 0 {
		tc.Extras = make(yaml.MapSlice, 0, len(c))
		for _, v := range c {
			if ms, ok := v.Value.(yaml.MapSlice); ok {
				var tcv TestConfigValue
				err := tcv.UnmarshalMap(ms)
				if err != nil {
					return err
				}
				tc.Extras = append(tc.Extras, yaml.MapItem{Key: v.Key, Value: tcv})
			}
		}
	}
	return nil
}

type TestConfigValue struct {
	Type         string `yaml:"type,omitempty"`
	DefaultValue any    `yaml:"defaultValue,omitempty"`
}

func (tc *TestConfigValue) UnmarshalMap(c yaml.MapSlice) error {
	tc.Type, c = extractValue[string](c, "type")
	tc.DefaultValue, _ = extractValue[any](c, "defaultValue")
	return nil
}

type TestStep struct {
	Parent *Test `yaml:"-"`

	Label                     string        `yaml:"label,omitempty"`
	Comments                  []string      `yaml:"-"`
	PICS                      string        `yaml:"PICS,omitempty"`
	Cluster                   string        `yaml:"cluster,omitempty"`
	Endpoint                  int64         `yaml:"endpoint,omitempty"`
	Command                   string        `yaml:"command,omitempty"`
	Attribute                 string        `yaml:"attribute,omitempty"`    // handled
	Verification              string        `yaml:"verification,omitempty"` // handled
	Arguments                 TestArguments `yaml:"arguments,omitempty"`
	Disabled                  bool          `yaml:"disabled,omitempty"`       // handled
	FabricFiltered            bool          `yaml:"fabricFiltered,omitempty"` //handled
	Response                  StepResponse  `yaml:"response,omitempty"`
	TimedInteractionTimeoutMs uint64        `yaml:"timedInteractionTimeoutMs,omitempty"`
	Event                     string        `yaml:"event,omitempty"`
	EventNumber               string        `yaml:"eventNumber,omitempty"`
	MaxInterval               uint64        `yaml:"maxInterval,omitempty"`
	MinInterval               uint64        `yaml:"minInterval,omitempty"`

	Extras yaml.MapSlice
}

func (ts *TestStep) UnmarshalMap(c yaml.MapSlice) error {
	ts.Label, c = extractValue[string](c, "label")
	ts.PICS, c = extractValue[string](c, "PICS")
	ts.Cluster, c = extractValue[string](c, "cluster")
	ts.Endpoint, c = extractValue[int64](c, "endpoint", -1)
	ts.Command, c = extractValue[string](c, "command")
	ts.Attribute, c = extractValue[string](c, "attribute")
	ts.Verification, c = extractValue[string](c, "verification")
	ts.Arguments, c = extractValue[TestArguments](c, "arguments")
	ts.Disabled, c = extractValue[bool](c, "disabled")
	ts.FabricFiltered, c = extractValue[bool](c, "fabricFiltered")
	ts.Response, c = extractValue[StepResponse](c, "response")
	ts.TimedInteractionTimeoutMs, c = extractValue[uint64](c, "timedInteractionTimeoutMs")
	ts.Event, c = extractValue[string](c, "event")
	ts.EventNumber, c = extractValue[string](c, "eventNumber")
	ts.MaxInterval, c = extractValue[uint64](c, "maxInterval")
	ts.MinInterval, c = extractValue[uint64](c, "minInterval")

	if len(c) > 0 {
		ts.Extras = make(yaml.MapSlice, len(c))
		copy(ts.Extras, c)
	}
	return nil
}

type TestArguments struct {
	Value  any   `yaml:"value,omitempty"`
	Values []any `yaml:"values,omitempty"`
}

func (ta *TestArguments) UnmarshalMap(c yaml.MapSlice) error {
	ta.Value, c = extractValue[any](c, "value")
	ta.Values, _ = extractArrayAny(c, "values")
	return nil
}

type TestArgumentValue struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type StepResponse struct {
	SaveAs      string                   `yaml:"saveAs,omitempty"`
	Error       string                   `yaml:"error,omitempty"`
	Value       any                      `yaml:"value,omitempty"`
	Values      []any                    `yaml:"values,omitempty"`
	Constraints *StepResponseConstraints `yaml:"constraints,omitempty"`

	Extras yaml.MapSlice
}

func (sr *StepResponse) UnmarshalMap(c yaml.MapSlice) error {
	sr.SaveAs, c = extractValue[string](c, "saveAs")
	sr.Error, c = extractValue[string](c, "error")
	sr.Value, c = extractValue[any](c, "value")
	sr.Values, c = extractArrayAny(c, "values")

	sr.Constraints, c = extractObject[StepResponseConstraints](c, "constraints")

	if len(c) > 0 {
		sr.Extras = make(yaml.MapSlice, len(c))
		copy(sr.Extras, c)
	}
	return nil
}

type StepResponseConstraints struct {
	Type string `yaml:"type,omitempty"`
	//	MinLength     any      `yaml:"minLength,omitempty"`
	//	MaxLength     any      `yaml:"maxLength,omitempty"`
	MinValue any `yaml:"minValue,omitempty"`
	MaxValue any `yaml:"maxValue,omitempty"`
	//	NotValue      any      `yaml:"notValue,omitempty"`
	//	HasValue      bool     `yaml:"hasValue,omitempty"`
	//HasMasksSet   []uint64 `yaml:"hasMasksSet,omitempty"`
	HasMasksClear []uint64 `yaml:"hasMasksClear,omitempty"`
	//Contains      any      `yaml:"contains,omitempty"`
	AnyOf any `yaml:"anyOf,omitempty"`
	//Excludes      []uint64 `yaml:"excludes,omitempty"`

	Extras yaml.MapSlice
}

func (src *StepResponseConstraints) UnmarshalMap(c yaml.MapSlice) error {
	src.Type, c = extractValue[string](c, "type")
	//src.MinLength = extractValue[any](c, "minLength")
	//src.MaxLength = extractValue[any](c, "maxLength")
	src.MinValue, c = extractValue[any](c, "minValue")
	src.MaxValue, c = extractValue[any](c, "maxValue")
	//src.NotValue = extractValue[any](c, "notValue")
	//src.HasValue = extractValue[bool](c, "hasValue")
	//src.HasMasksSet = extractArray[uint64](c, "hasMasksSet")
	src.HasMasksClear, c = extractArray[uint64](c, "hasMasksClear")
	//src.Contains = extractValue[any](c, "contains")
	src.AnyOf, c = extractValue[any](c, "anyOf")
	//src.Excludes = extractArray[uint64](c, "excludes")
	//
	if len(c) > 0 {
		src.Extras = make(yaml.MapSlice, len(c))
		copy(src.Extras, c)
	}

	return nil
}
