package parser

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

var file2 = `; last modified 1 April 2001 by John Doe
[owner]
this is invalid line
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
		got, _ := p1.Get("[owner]", "name")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing of value-2", func(t *testing.T) {

		want := `"payroll.dat"`
		got, _ := p1.Get("[database]", "file")
		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("running testing key on an invalid key", func(t *testing.T) {

		_, err := p1.Get("database", "invalid")
		if err == nil {
			t.Errorf("error there should be error her")
		}
	})

	t.Run("running testing key on an invalid section", func(t *testing.T) {
		_, err := p1.Get("invalid", "port")
		if err == nil {
			t.Errorf("there should be error here")
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
	t.Run("running test Set with new section name", func(t *testing.T) {
		p1.Set("[new_section]", "age", "21")
		want := "21"
		got := p1.dict["[new_section]"]["age"]
		if want != got {
			t.Errorf("got %q , want %q", got, want)
		}
	})

}
func TestLoadFromString(t *testing.T) {
	t.Run("running TestloadFromString on valid syntax", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(file1)
		if got != nil {
			t.Errorf("there should not be an error")
		}
	})
	t.Run("running TestloadFromString on invalid syntax ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(file2)
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
}
func TestLoadFromFile(t *testing.T) {
	t.Run("running TestloadFromFile on valid file and syntax", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromFile("file1.ini")
		if got != nil {
			t.Errorf("there should not be an error")
		}
	})
	t.Run("running TestloadFromFile on invalid file ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString("invalid_file_name.ini")
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
	t.Run("running TestloadFromFile on valid file but invalid syntax ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString("invalid_syntax_file.ini")
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
	t.Run("running TestloadFromFile on invalid file name syntax ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString("invalid_syntax_file.html")
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
}
