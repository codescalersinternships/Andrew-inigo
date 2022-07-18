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

var dict1 = make(map[string]map[string]string)
var p1 = Parser{file1, dict1}

//test functions

func TestGetSectionNames(t *testing.T) {
	t.Run("running TestGetSectionNames:1 ", func(t *testing.T) {
		got := p1.GetSectionNames()
		want := []string{"owner", "database"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q  want %q", got, want)
		}
	})
}
func TestGetSections(t *testing.T) {
	t.Run("testing get sections:1", func(t *testing.T) {
		want := map[string]map[string]string{
			"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": `"payroll.dat"`},
		}
		got := p1.GetSections()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q  want %q", got, want)
		}

	})
}
func TestMakeDict(t *testing.T) {
	t.Run("trying first", func(t *testing.T) {
		want := map[string]string{"server": "192.0.2.62",
			"port": "143", "file": `"payroll.dat"`}
		got := make_dictionary(file1, "database")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q \n want %q", got, want)
		}
	})
	t.Run("trying second", func(t *testing.T) {
		want := map[string]string{"name": "John Doe",
			"organization": "Acme Widgets Inc."}
		got := make_dictionary(file1, "owner")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q \n want %q", got, want)
		}
	})
}
func TestGetValue(t *testing.T) {
	t.Run("running testing of value-1", func(t *testing.T) {

		want := "John Doe"
		got := p1.Getvalue("owner", "name")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing of value-2", func(t *testing.T) {

		want := `"payroll.dat"`
		got := p1.Getvalue("database", "file")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing key on an invalid key", func(t *testing.T) {

		want := ""
		got := p1.Getvalue("database", "invalid")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing key on an invalid section", func(t *testing.T) {
		want := ""
		got := p1.Getvalue("invalid", "port")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
func TestSetKeyValue(t *testing.T) {
	//testing this by retriecing the value from the key
	t.Run("running test SetKeyValue (add key:age value :21 to owner) :1", func(t *testing.T) {
		p1.SetKeyValue("owner", "age", "21")
		want := "21"
		got := p1.dict["owner"]["age"]
		if want != got {
			t.Errorf("got %q , want %q", got, want)
		}
	})
}
