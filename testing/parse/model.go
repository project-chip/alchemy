package parse

import (
	"fmt"
	"log/slog"
)

type Test struct {
	Name string   `yaml:"name,omitempty"`
	PICS []string `yaml:"PICS,omitempty"`

	Config TestConfig  `yaml:"config,omitempty"`
	Tests  []*TestStep `yaml:"tests,omitempty"`

	Extras map[string]any
}

func (t *Test) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	slog.Info("Unmarshalling!")
	var yt map[string]any
	if err = unmarshal(&yt); err != nil {
		return err
	}
	t.Name = extractValue[string](yt, "name")
	t.PICS = extractArray[string](yt, "PICS")

	config := extractValue[map[string]any](yt, "config")
	if config != nil {
		err = t.Config.UnmarshalMap(config)
		if err != nil {
			return
		}
	}
	tests := extractValue[[]any](yt, "tests")
	for _, test := range tests {
		switch test := test.(type) {
		case map[string]any:
			var ts TestStep
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
		t.Extras = make(map[string]any)
		for key, val := range yt {
			t.Extras[key] = val
			slog.Info("unmarshalled unknown!", slog.String("key", key), slog.String("val", fmt.Sprintf("%T", val)))
		}
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

	Extras map[string]any
}

func (tc *TestConfig) UnmarshalMap(c map[string]any) error {
	tc.NodeID = extractValue[uint64](c, "nodeId")
	tc.Cluster = extractValue[string](c, "cluster")
	tc.Endpoint = extractValue[uint64](c, "endpoint")
	tc.Timeout = extractValue[uint64](c, "timeout")
	tc.CatalogVendorId = extractValue[TestConfigValue](c, "catalogVendorId")
	tc.ApplicationId = extractValue[TestConfigValue](c, "applicationId")
	tc.Discriminator = extractValue[TestConfigValue](c, "discriminator")
	tc.WaitAfterCommissioning = extractValue[TestConfigValue](c, "waitAfterCommissioning")

	if len(c) > 0 {
		tc.Extras = make(map[string]any)
		for key, val := range c {
			tc.Extras[key] = val
			slog.Info("unmarshalled config unknown!", slog.String(key, key), slog.String("val", fmt.Sprintf("%T", val)))
		}
	}
	return nil
}

type TestConfigValue struct {
	Type         string `yaml:"type,omitempty"`
	DefaultValue any    `yaml:"defaultValue,omitempty"`
}

func (tc *TestConfigValue) UnmarshalMap(c map[string]any) error {
	tc.Type = extractValue[string](c, "type")
	tc.DefaultValue = extractValue[any](c, "defaultValue")
	return nil
}

type TestStep struct {
	Step                      string        `yaml:"-"`
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

	Extras map[string]any
}

func (ts *TestStep) UnmarshalMap(c map[string]any) error {
	ts.Label = extractValue[string](c, "label")
	ts.PICS = extractValue[string](c, "PICS")
	ts.Cluster = extractValue[string](c, "cluster")
	ts.Endpoint = extractValue[int64](c, "endpoint", -1)
	ts.Command = extractValue[string](c, "command")
	ts.Attribute = extractValue[string](c, "attribute")
	ts.Verification = extractValue[string](c, "verification")
	ts.Arguments = extractValue[TestArguments](c, "arguments")
	ts.Disabled = extractValue[bool](c, "disabled")
	ts.FabricFiltered = extractValue[bool](c, "fabricFiltered")
	ts.Response = extractValue[StepResponse](c, "response")
	ts.TimedInteractionTimeoutMs = extractValue[uint64](c, "timedInteractionTimeoutMs")
	ts.Event = extractValue[string](c, "event")
	ts.EventNumber = extractValue[string](c, "eventNumber")
	ts.MaxInterval = extractValue[uint64](c, "maxInterval")
	ts.MinInterval = extractValue[uint64](c, "minInterval")

	if len(c) > 0 {
		ts.Extras = make(map[string]any)
		for key, val := range c {
			ts.Extras[key] = val
			slog.Info("unmarshalled test step unknown!", slog.String("key", key), slog.String("type", fmt.Sprintf("%T", val)), slog.Any("val", val))
		}
	}
	return nil
}

type TestArguments struct {
	Value  any   `yaml:"value,omitempty"`
	Values []any `yaml:"values,omitempty"`
}

func (ta *TestArguments) UnmarshalMap(c map[string]any) error {
	ta.Value = extractValue[any](c, "value")
	ta.Values = extractArray[any](c, "values")
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

	Extras map[string]any
}

func (sr *StepResponse) UnmarshalMap(c map[string]any) error {
	slog.Info("unmarshalling response")
	sr.SaveAs = extractValue[string](c, "saveAs")
	sr.Error = extractValue[string](c, "error")
	sr.Value = extractValue[any](c, "value")
	slog.Info("unmarshalled response value", slog.Any("val", sr.Value))
	sr.Values = extractArray[any](c, "values")

	sr.Constraints = extractObject[StepResponseConstraints](c, "constraints")

	if len(c) > 0 {
		sr.Extras = make(map[string]any)
		for key, val := range c {
			sr.Extras[key] = val
			slog.Info("unmarshalled test step response unknown!", slog.String(key, key), slog.String("val", fmt.Sprintf("%T", val)), slog.Any("val", val))
		}
	}
	return nil
}

type StepResponseConstraints struct {
	//	Type          string   `yaml:"type,omitempty"`
	//	MinLength     any      `yaml:"minLength,omitempty"`
	//	MaxLength     any      `yaml:"maxLength,omitempty"`
	MinValue any `yaml:"minValue,omitempty"`
	MaxValue any `yaml:"maxValue,omitempty"`
	//	NotValue      any      `yaml:"notValue,omitempty"`
	//	HasValue      bool     `yaml:"hasValue,omitempty"`
	//HasMasksSet   []uint64 `yaml:"hasMasksSet,omitempty"`
	//HasMasksClear []uint64 `yaml:"hasMasksClear,omitempty"`
	//Contains      any      `yaml:"contains,omitempty"`
	AnyOf any `yaml:"anyOf,omitempty"`
	//Excludes      []uint64 `yaml:"excludes,omitempty"`

	//Extras map[string]any
}

func (src *StepResponseConstraints) UnmarshalMap(c map[string]any) error {
	//src.Type = extractValue[string](c, "type")
	//src.MinLength = extractValue[any](c, "minLength")
	//src.MaxLength = extractValue[any](c, "maxLength")
	src.MinValue = extractValue[any](c, "minValue")
	src.MaxValue = extractValue[any](c, "maxValue")
	//src.NotValue = extractValue[any](c, "notValue")
	//src.HasValue = extractValue[bool](c, "hasValue")
	//src.HasMasksSet = extractArray[uint64](c, "hasMasksSet")
	//src.HasMasksClear = extractArray[uint64](c, "hasMasksClear")
	//src.Contains = extractValue[any](c, "contains")
	src.AnyOf = extractValue[any](c, "anyOf")
	//src.Excludes = extractArray[uint64](c, "excludes")
	//
	//if len(c) > 0 {
	//	src.Extras = make(map[string]any)
	//	for key, val := range c {
	//		src.Extras[key] = val
	//		slog.Info("unmarshalled test step response constraints unknown!", slog.String(key, key), slog.String("val", fmt.Sprintf("%T", val)), slog.Any("val", val))
	//	}
	//}
	return nil
}
