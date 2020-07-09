package tribool

import (
	"encoding/json"
	"encoding/xml"
)

// Tribool acts like the built-in bool type, but for 3-state boolean logic. The three states are True, False, and Indeterminate.
type Tribool int8

// Tribool
const (
	Indeterminate Tribool = iota
	True
	False
)

// String impelements Stringer interface
func (b *Tribool) String() string {
	if b != nil {
		switch *b {
		case True:
			return "true"
		case False:
			return "false"
		}
	}
	return "indeterminate"
}

// Test returns true only when the tribool is "True"
func (b *Tribool) Test() bool {
	return *b == True
}

// Set sets the tribool to "True" or "False"
func (b *Tribool) Set(flag bool) {
	if flag == true {
		*b = True
	} else {
		*b = False
	}
}

// Unset set the tribool to "Indeterminate"
func (b *Tribool) Unset() {
	*b = Indeterminate
}

// MarshalJSON implements json.Marshaler interface
func (b *Tribool) MarshalJSON() ([]byte, error) {
	switch *b {
	case True:
		return json.Marshal(true)
	case False:
		return json.Marshal(false)
	default:
		return []byte("null"), nil
	}
}

// UnmarshalJSON implements json.Unmarshaler interface
func (b *Tribool) UnmarshalJSON(dat []byte) error {
	switch string(dat) {
	case "true":
		*b = True
	case "false":
		*b = False
	default:
		*b = Indeterminate
	}
	return nil
}

// MarshalXML implements xml.Marshaler interface
func (b *Tribool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(xml.CharData(b.String()), start)
}

// MarshalXMLAttr implements xml.MarshalerAttr interface
func (b *Tribool) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: b.String()}, nil
}

// UnmarshalXML implements xml.Unmarshaler interface
func (b *Tribool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var s string
	err = d.DecodeElement(&s, &start)

	if err == nil {
		switch s {
		case "true":
			*b = True
		case "false":
			*b = False
		default:
			*b = Indeterminate
		}
	}
	return
}

// UnmarshalXMLAttr implements xml.UnmarshalerAttr interface
func (b *Tribool) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "true":
		*b = True
	case "false":
		*b = False
	default:
		*b = Indeterminate
	}
	return nil
}
