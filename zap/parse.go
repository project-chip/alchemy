package zap

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
	"github.com/hasty/matterfmt/render/zcl"
)

func Parse(r io.Reader) (models []any, err error) {
	d := xml.NewDecoder(r)
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = nil
			return
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.ProcInst:
		case xml.CharData:
		case xml.StartElement:
			switch t.Name.Local {
			case "configurator":
				var cm []any
				cm, err = readConfigurator(d)
				if err == nil {
					models = append(models, cm...)
				}
			default:
				err = fmt.Errorf("unexpected top level element: %s", t.Name.Local)
			}

		default:
			err = fmt.Errorf("unexpected top level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readSimpleElement(d *xml.Decoder, name string) (val string, err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", name)
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData:
			val = string(t)
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func ignore(d *xml.Decoder, name string) (err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of %s", name)
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return nil
			default:
			}
		case xml.CharData:
		default:
		}
		if err != nil {
			return
		}
	}
}

func readConfigurator(d *xml.Decoder) (models []any, err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of configurator")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "cluster":
				var cluster *matter.Cluster
				cluster, err = readCluster(d, t)
				if err == nil {
					models = append(models, cluster)
				}
			case "domain":
				_, err = readSimpleElement(d, t.Name.Local)
			case "enum":
				var en *matter.Enum
				en, err = readEnum(d, t)
				if err == nil {
					models = append(models, en)
				}
			case "struct":
				var s *matter.Struct
				s, err = readStruct(d, t)
				if err == nil {
					models = append(models, s)
				}
			case "bitmap":
				var bitmap *matter.Bitmap
				bitmap, err = readBitmap(d, t)
				if err == nil {
					models = append(models, bitmap)
				}
			case "accessControl", "atomic", "clusterExtension", "global", "deviceType":
				err = ignore(d, t.Name.Local)
			default:
				err = fmt.Errorf("unexpected configurator level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "configurator":
				return
			default:
				err = fmt.Errorf("unexpected configurator end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected configurator level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readCluster(d *xml.Decoder, e xml.StartElement) (cluster *matter.Cluster, err error) {
	cluster = &matter.Cluster{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "singleton":
		case "apiMaturity":
		default:
			return nil, fmt.Errorf("unexpected cluster attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return nil, fmt.Errorf("EOF before end of cluster")
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "attribute":
				var attribute *matter.Field
				attribute, err = readAttribute(d, t)
				if err == nil {
					cluster.Attributes = append(cluster.Attributes, attribute)
				}
			case "client":
				err = readClient(d, t)
			case "server":
				err = readServer(d, t)
			case "code":
				cluster.ID, err = readSimpleElement(d, t.Name.Local)
			case "define":
				_, err = readSimpleElement(d, t.Name.Local)
			case "description":
				cluster.Description, err = readSimpleElement(d, t.Name.Local)
			case "domain":
				_, err = readSimpleElement(d, t.Name.Local)
			case "event":
				var event *matter.Event
				event, err = readEvent(d, t)
				if err == nil {
					cluster.Events = append(cluster.Events, event)
				}
			case "name":
				cluster.Name, err = readSimpleElement(d, t.Name.Local)
			case "command":
				err = readCommand(d, t)
			case "globalAttribute":
				err = ignore(d, t.Name.Local)
			default:
				err = fmt.Errorf("unexpected cluster level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "cluster":
				return
			default:
				err = fmt.Errorf("unexpected cluster end element: %s", t.Name.Local)
			}
		case xml.Comment:
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected cluster level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readAttribute(d *xml.Decoder, e xml.StartElement) (attr *matter.Field, err error) {
	attr = &matter.Field{Type: &matter.DataType{}}
	err = readFieldAttributes(e, attr, "attribute")
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of attribute")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "access":
				a := &matter.Access{}
				err = readAccess(d, t, a)
			case "description":
				_, err = readSimpleElement(d, t.Name.Local)
			default:
				err = fmt.Errorf("unexpected attribute level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "attribute":
				return
			default:
				err = fmt.Errorf("unexpected attribute end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected attribute level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readAccess(d *xml.Decoder, e xml.StartElement, access *matter.Access) (err error) {
	var op string
	var privilege string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "op":
			op = a.Value
		case "role", "privilege":
			privilege = a.Value
		default:
			return fmt.Errorf("unexpected access attribute: %s", a.Name.Local)
		}
	}
	p := parsePrivilege(strings.ToLower(privilege))
	if p == matter.PrivilegeUnknown {
		return fmt.Errorf("unknown privilege value: %s", privilege)
	}
	switch strings.ToLower(op) {
	case "read":
		access.Read = p
	case "write":
		access.Write = p
	case "invoke":
		access.Invoke = p
	default:
		return fmt.Errorf("unknown privilege value: %s", privilege)
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of access")

		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "access":
				return nil
			default:
				return fmt.Errorf("unexpected access end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected access level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readClient(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "init":
		case "tick":
		default:
			return fmt.Errorf("unexpected client attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of client")

		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "client":
				return nil
			default:
				return fmt.Errorf("unexpected client end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected client level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readServer(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "init":
		case "tick":
		case "tickFrequency":
		default:
			return fmt.Errorf("unexpected server attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of server")
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "server":
				return nil
			default:
				return fmt.Errorf("unexpected server end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected server level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readEvent(d *xml.Decoder, e xml.StartElement) (event *matter.Event, err error) {
	event = &matter.Event{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "side":
		case "code":
			event.ID = a.Value
		case "priority":
			event.Priority = a.Value
		case "name":
			event.Name = a.Value
		case "isFabricSensitive":
			event.FabricSensitive = (a.Value == "true")
		case "optional":
			if a.Value == "true" {
				event.Conformance = "M"
			}
		default:
			return nil, fmt.Errorf("unexpected event attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of event")

		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "access":
				if event.Access == nil {
					event.Access = &matter.Access{}
				}
				err = readAccess(d, t, event.Access)
			case "description":
				event.Description, err = readSimpleElement(d, t.Name.Local)
			case "field":
				var field *matter.Field
				field, err = readField(d, t, "field")
				if err == nil {
					event.Fields = append(event.Fields, field)
				}
			default:
				err = fmt.Errorf("unexpected event level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "event":
				return
			default:
				err = fmt.Errorf("unexpected event end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected event level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readField(d *xml.Decoder, e xml.StartElement, name string) (field *matter.Field, err error) {
	field = &matter.Field{Type: &matter.DataType{}}
	err = readFieldAttributes(e, field, name)
	if err != nil {
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func readFieldAttributes(e xml.StartElement, field *matter.Field, name string) error {
	var max, min, length string
	var fieldType, entryType string
	var isArray bool
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "id", "fieldId", "code": // Pick a lane, jeez
			field.ID = a.Value
		case "name":
			field.Name = a.Value
		case "isFabricSensitive":
			if field.Access == nil {
				field.Access = &matter.Access{}
			}
			field.Access.FabricSensitive = a.Value == "true"
		case "isNullable":
			field.Quality |= matter.QualityNullable
		case "type":
			fieldType = a.Value
		case "max":
			max = a.Value
		case "min":
			min = a.Value
		case "entryType":
			entryType = a.Value
		case "array":
			isArray = (a.Value == "true")
		case "length":
			length = a.Value
		case "optional":
			if a.Value != "true" {
				field.Conformance = "M"
			}
		case "writable": // Ugh

		default:
			return fmt.Errorf("unexpected %s attribute: %s", name, a.Name.Local)
		}
	}
	fieldType = zcl.ConvertZapToDataType(fieldType)
	entryType = zcl.ConvertZapToDataType(entryType)
	if isArray {
		if len(entryType) > 0 {
			field.Type.Name = entryType
		} else {
			field.Type.Name = fieldType
		}
		field.Type.IsArray = true
	} else if fieldType == "ARRAY" {
		field.Type.Name = entryType
		field.Type.IsArray = true
	} else {
		field.Type.Name = fieldType
	}
	if len(max) > 0 && len(min) > 0 {
		field.Constraint = ascii.ParseConstraint(field.Type, fmt.Sprintf("%s to %s", min, max))
	} else if len(max) > 0 {
		field.Constraint = ascii.ParseConstraint(field.Type, fmt.Sprintf("max %s", max))
	} else if len(min) > 0 {
		field.Constraint = ascii.ParseConstraint(field.Type, fmt.Sprintf("min %s", min))
	} else if len(length) > 0 {
		field.Constraint = ascii.ParseConstraint(field.Type, fmt.Sprintf("max %s", length))
	}
	return nil
}

func readEnum(d *xml.Decoder, e xml.StartElement) (en *matter.Enum, err error) {
	en = &matter.Enum{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			en.Name = a.Value
		case "type":
			en.Type = zcl.ConvertZapToDataType(a.Value)
		default:
			return nil, fmt.Errorf("unexpected enum attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of enum")

		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			//case "access":
			//	err = readAccess(d, t, a)
			case "description":
				en.Description, err = readSimpleElement(d, t.Name.Local)
			case "item":
				var ev *matter.EnumValue
				ev, err = readEnumItem(d, t)
				if err == nil {
					en.Values = append(en.Values, ev)
				}
			case "cluster":
				err = readClusterCode(d, t)
			default:
				err = fmt.Errorf("unexpected enum level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "enum":
				return
			default:
				err = fmt.Errorf("unexpected enum end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected enum level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readEnumItem(d *xml.Decoder, e xml.StartElement) (ev *matter.EnumValue, err error) {
	ev = &matter.EnumValue{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			ev.Name = a.Value
		case "value":
			ev.Value = a.Value
		default:
			return nil, fmt.Errorf("unexpected enum item attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of enum item")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "item":
				return
			default:
				err = fmt.Errorf("unexpected enum item end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected enum item level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readClusterCode(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "code":
		default:
			return fmt.Errorf("unexpected cluster code attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of cluster code")

		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "cluster":
				return nil
			default:
				return fmt.Errorf("unexpected cluster code end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected cluster code level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readStruct(d *xml.Decoder, e xml.StartElement) (s *matter.Struct, err error) {
	s = &matter.Struct{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			s.Name = a.Value
		case "isFabricScoped":
		default:
			return nil, fmt.Errorf("unexpected struct attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of struct")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "cluster":
				err = readClusterCode(d, t)
			/*case "access":
			a := &matter.Access{}
			err = readAccess(d, t, a)*/
			case "description":
				s.Description, err = readSimpleElement(d, t.Name.Local)
			case "item":
				var f *matter.Field
				f, err = readField(d, t, "item")
				if err != nil {
					s.Fields = append(s.Fields, f)
				}
			default:
				err = fmt.Errorf("unexpected struct level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "struct":
				return
			default:
				err = fmt.Errorf("unexpected struct end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected struct level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readCommand(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "source":
		case "code":
		case "name":
		case "isFabricScoped":
		case "optional":
		case "response":
		case "mustUseTimedInvoke":
		case "cli":
		case "disableDefaultResponse":
		case "apiMaturity":
		default:
			return fmt.Errorf("unexpected command attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of command")
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "access":
				a := &matter.Access{}
				err = readAccess(d, t, a)
			case "description":
				_, err = readSimpleElement(d, t.Name.Local)
			case "arg":
				err = readArg(d, t)
			default:
				return fmt.Errorf("unexpected command level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "command":
				return nil
			default:
				return fmt.Errorf("unexpected command end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected command level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readArg(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
		case "length", "lenght":
		case "minLength":
		case "type":
		case "entryType":
		case "min":
		case "max":
		case "optional":
		case "array":
		case "default":
		case "isNullable":
		default:
			return fmt.Errorf("unexpected arg attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of arg")
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "arg":
				return nil
			default:
				return fmt.Errorf("unexpected arg end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected arg level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readBitmap(d *xml.Decoder, e xml.StartElement) (bitmap *matter.Bitmap, err error) {
	bitmap = &matter.Bitmap{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			bitmap.Name = a.Value
		case "type":
			bitmap.Type = zcl.ConvertZapToDataType(a.Value)
		case "apiMaturity":
		default:
			return nil, fmt.Errorf("unexpected bitmap attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of bitmap")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "cluster":
				err = readClusterCode(d, t)
			case "description":
				bitmap.Description, err = readSimpleElement(d, t.Name.Local)
			case "field":
				var bit *matter.BitmapValue
				bit, err = readBitmapField(d, t)
				if err == nil {
					bitmap.Bits = append(bitmap.Bits, bit)
				}
			default:
				err = fmt.Errorf("unexpected bitmap level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "bitmap":
				return
			default:
				err = fmt.Errorf("unexpected bitmap end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected bitmap level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func readBitmapField(d *xml.Decoder, e xml.StartElement) (bv *matter.BitmapValue, err error) {
	bv = &matter.BitmapValue{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			bv.Name = a.Value
		case "mask":
			var mask uint64
			mask, err = parse.HexOrDec(a.Value)
			if err != nil {
				return
			}
			bit := 0
			for mask > 0 {
				if (mask & 1) == 1 {
					if mask > 1 {
						err = fmt.Errorf("non-power of 2 mask: %s", a.Value)
						return
					}
					break
				}
				bit++
				mask >>= 1
			}
			bv.Bit = strconv.Itoa(bit)
		case "optional":
			if a.Value != "true" {
				bv.Conformance = "M"
			}
		default:
			return nil, fmt.Errorf("unexpected bitmap field attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "field":
				return
			default:
				err = fmt.Errorf("unexpected field end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected field level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
