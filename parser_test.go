package main

import (
	"reflect"
	"testing"
)

var file1 = `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"`

func TestGetSectionNames(t *testing.T) {
	t.Run("running TestGetSectionNames:1 ", func(t *testing.T) {
		p1 := Parser{}
		p1.LoadFromString(file1)
		got := p1.GetSectionNames()
		want := []string{"[owner]", "[database]"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q  want %q", got, want)
		}
	})
}
func TestGetSections(t *testing.T) {
	p1 := Parser{}
	p1.LoadFromString(file1)
	t.Run("testing get sections:1", func(t *testing.T) {
		want := map[string]map[string]string{
			"[owner]":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"[database]": {"server": "192.0.2.62", "port": "143", "file": `"payroll.dat"`},
		}
		got := p1.GetSections()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q  want %q", got, want)
		}

	})
}
func TestGetValue(t *testing.T) {
	p1 := Parser{}
	p1.LoadFromString(file1)
	t.Run("running testing of value-1", func(t *testing.T) {
		want := "John Doe"
		got := p1.Get("[owner]", "name")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing of value-2", func(t *testing.T) {

		want := `"payroll.dat"`
		got := p1.Get("[database]", "file")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing key on an invalid key", func(t *testing.T) {

		want := ""
		got := p1.Get("database", "invalid")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing key on an invalid section", func(t *testing.T) {
		want := ""
		got := p1.Get("invalid", "port")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
func TestSetKeyValue(t *testing.T) {
	p1 := Parser{}
	p1.LoadFromString(file1)
	//testing this by retriecing the value from the key
	t.Run("running test SetKeyValue (add key:age value :21 to owner) :1", func(t *testing.T) {
		p1.Set("[owner]", "age", "21")
		want := "21"
		got := p1.dict["[owner]"]["age"]
		if want != got {
			t.Errorf("got %q , want %q", got, want)
		}
	})
}

func TestLoadFromFile(t *testing.T) {
	t.Run("running TestloadFromFile on valid file", func(t *testing.T) {
		p1 := Parser{}
		p1.LoadFromString(file1)
		got := p1.LoadFromFile("file1.ini")
		if got != nil {
			t.Errorf("want Null , got %q", got)
		}
	})
	t.Run("running TestloadFromFile on invalid file", func(t *testing.T) {
		p1 := Parser{}
		p1.LoadFromString(file1)
		got := p1.LoadFromFile("invalid.ini")
		if got == nil {
			t.Errorf("want Error , got Null")
		}
	})
}
