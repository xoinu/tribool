package tribool

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestTribool(t *testing.T) {
	var b Tribool

	if b != Indeterminate {
		t.Fatal(`b != Indeterminate`)
	}

	s := b.String()
	if s != "indeterminate" {
		t.Fatal(`s != "indeterminate"`)
	}

	b.Set(true)
	if b.Test() != true {
		t.Fatal(`b.Test() != true`)
	}

	s = b.String()
	if s != "true" {
		t.Fatal(`s != "true"`)
	}

	b.Set(false)
	if b.Test() != false {
		t.Fatal(`b.Test() != false`)
	}

	s = b.String()
	if s != "false" {
		t.Fatal(`s != "false"`)
	}

	b.Unset()
	if b.Test() != false {
		t.Fatal(`b.Test() != false`)
	}
}

func TestJSON(t *testing.T) {
	type test struct {
		Bool Tribool `json:"bool"`
	}
	var val test
	val.Bool = True
	dat, err := json.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}
	if string(dat) != `{"bool":true}` {
		t.Fatal("Failed to json.Marshal:", string(dat))
	}
	var val2 test
	err = json.Unmarshal(dat, &val2)
	if err != nil {
		t.Fatal(err)
	}
	if val2.Bool != True {
		t.Fatal(`val2.Bool != True`)
	}

	val.Bool = False
	dat, err = json.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}
	if string(dat) != `{"bool":false}` {
		t.Fatal("Failed to json.Marshal:", string(dat))
	}
	err = json.Unmarshal(dat, &val2)
	if err != nil {
		t.Fatal(err)
	}
	if val2.Bool != False {
		t.Fatal(`val2.Bool != True`)
	}

	val.Bool = Indeterminate
	dat, err = json.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}
	if string(dat) != `{"bool":null}` {
		t.Fatal("Failed to json.Marshal:", string(dat))
	}
	err = json.Unmarshal(dat, &val2)
	if err != nil {
		t.Fatal(err)
	}
	if val2.Bool != Indeterminate {
		t.Fatal(`val2.Bool != Indeterminate`)
	}
}

func TestXML(t *testing.T) {
	type test struct {
		Bool  Tribool
		Bool1 Tribool `xml:"bool1,attr"`
	}
	var val test
	val.Bool = True
	val.Bool1 = True
	dat, err := xml.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}

	var val2 test
	err = xml.Unmarshal(dat, &val2)
	if err != nil {
		t.Fatal(err)
	}
	if val2.Bool != True {
		t.Fatal(`val2.Bool != True`)
	}
	if val2.Bool1 != True {
		t.Fatal(`val2.Bool1 != True`)
	}

	val.Bool = False
	val.Bool1 = False
	dat, err = xml.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}

	err = xml.Unmarshal(dat, &val2)
	if err != nil {
		t.Fatal(err)
	}
	if val2.Bool != False {
		t.Fatal(`val2.Bool != False`)
	}
	if val2.Bool1 != False {
		t.Fatal(`val2.Bool1 != False`)
	}

	val.Bool.Unset()
	val.Bool1.Unset()
	dat, err = xml.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}
	err = xml.Unmarshal(dat, &val2)
	if err != nil {
		t.Fatal(err)
	}
	if val2.Bool != Indeterminate {
		t.Fatal(`val2.Bool != Indeterminate`)
	}
	if val2.Bool1 != Indeterminate {
		t.Fatal(`val2.Bool1 != Indeterminate`)
	}
}
