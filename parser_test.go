package parser

import (
	"reflect"
	"testing"
)

var correct_text = `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"`
var incorrect_line_text = `; last modified 1 April 2001 by John Doe
[owner]
this is invalid line
name = John Doe
organization = Acme Widgets Inc.
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"`
var missing_section_bracket_text = `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[owner2
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"`
var empty_section_text = `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[]
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"`
var empty_key_text = `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
 = Acme Widgets Inc.

; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"`

func TestGetSectionNames(t *testing.T) {
	assertError := func(t testing.TB, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q  want %q", got, want)
		}
	}
	t.Run("running TestGetSectionNames:1 ", func(t *testing.T) {
		p1 := Parser{}
		p1.LoadFromString(correct_text)
		got := p1.GetSectionNames()
		want := []string{"[owner]", "[database]"}
		assertError(t, got, want)
	})
}
func TestGetSections(t *testing.T) {
	assertError := func(t testing.TB, got, want map[string]map[string]string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q  want %q", got, want)
		}
	}
	p1 := Parser{}
	p1.LoadFromString(correct_text)
	t.Run("testing get sections:1", func(t *testing.T) {
		want := map[string]map[string]string{
			"[owner]":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"[database]": {"server": "192.0.2.62", "port": "143", "file": `"payroll.dat"`},
		}
		got := p1.GetSections()
		assertError(t, got, want)
	})
}
func TestGetValue(t *testing.T) {
	assertError := func(t testing.TB, got, want string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q  want %q", got, want)
		}
	}
	p1 := Parser{}
	p1.LoadFromString(correct_text)
	t.Run("running testing of value-1", func(t *testing.T) {
		want := "John Doe"
		got, _ := p1.Get("[owner]", "name")
		assertError(t, got, want)
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
	assertError := func(t testing.TB, got, want string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q  want %q", got, want)
		}
	}
	p1 := Parser{}
	p1.LoadFromString(correct_text)
	//testing this by retriecing the value from the key
	t.Run("running test SetKeyValue (add key:age value :21 to owner) :1", func(t *testing.T) {
		p1.Set("[owner]", "age", "21")
		want := "21"
		got := p1.dict["[owner]"]["age"]
		assertError(t, got, want)
	})
	t.Run("running test Set with new section name", func(t *testing.T) {
		p1.Set("[new_section]", "age", "21")
		want := "21"
		got := p1.dict["[new_section]"]["age"]
		assertError(t, got, want)
	})
	t.Run("running test Set to add nil to value", func(t *testing.T) {
		p1.Set("[owner]", "name", "")
		want := ""
		got := p1.dict["[owner]"]["name"]
		assertError(t, got, want)
	})
	t.Run("running test Set with empty key", func(t *testing.T) {
		got := p1.Set("[owner]", "", "value")

		if got == nil {
			t.Errorf("expected error here")
		}
	})
	t.Run("running test Set with empty section", func(t *testing.T) {
		got := p1.Set("", "owner", "value")

		if got == nil {
			t.Errorf("expected error here")
		}
	})
}
func TestLoadFromString(t *testing.T) {
	t.Run("running TestloadFromString on valid syntax", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(correct_text)
		if got != nil {
			t.Errorf("there should not be an error")
		}
	})
	t.Run("running TestloadFromString on invalid syntax ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(incorrect_line_text)
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
	t.Run("running TestloadFromString for missing bracket ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(missing_section_bracket_text)
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
	t.Run("running TestloadFromString for empty section name ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(empty_section_text)
		if got == nil {
			t.Errorf("there should be error here")
		}
	})
	t.Run("running TestloadFromString for empty key value ", func(t *testing.T) {
		p1 := Parser{}
		got := p1.LoadFromString(empty_key_text)
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

func TestSaveToFile(t *testing.T) {
	t.Run("testing SaveToFIile with invalid format not .ini", func(t *testing.T) {
		p1 := Parser{}
		got := p1.SaveToFile("file1.html")
		if got == nil {
			t.Errorf("there should not be an error")
		}
	})
}
