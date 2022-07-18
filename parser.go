package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	text string
	dict map[string]map[string]string
}

// Load from file function just reads the file context and add it to the the text attribute
func (p *Parser) LoadFromFile(file string) string {
	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	Text := string(dat)
	return Text
}

// LoadFromString function parse the text and make the dictionary
func (p *Parser) LoadFromString(Text string) {
	p.text = Text
	p.dict = make(map[string]map[string]string)
	p.dict = p.Parsing()
}

// Get section names function returns a slice "array" of section names
func (p Parser) GetSectionNames() []string {
	r, _ := regexp.Compile(` *\[.*`)
	sections := get_certain_lines(p.text, r)
	return sections
}

//parsing function returns dictionary containg [section , [key , value]]
func (p Parser) Parsing() map[string]map[string]string {
	var res_dict = make(map[string]map[string]string)
	var dict2 = make(map[string]string)
	sections := p.GetSectionNames() //array of section name instance and return it
	for _, section := range sections {
		dict2 = make_key_value_dict(p.text, section)
		res_dict[section] = dict2
	}
	return res_dict
}

func (p Parser) GetSections() map[string]map[string]string {
	return p.dict
}

//Getvalue function return value for certain section and key
func (p Parser) Getvalue(section, key string) string {
	return p.dict[section][key]
}

//setkeyvalue functions add another key and value to a certain section
func (p *Parser) SetKeyValue(section, key, value string) {
	fmt.Println("------------------")
	fmt.Println(p.dict)
	p.dict[section][key] = value
}

// ToString returns the string in INI format after changes happened to it
func (p Parser) ToString() string {
	var res_string string
	for section, dictionary := range p.dict {
		res_string += "\n"
		res_string += section
		for key, value := range dictionary {
			res_string += "\n"
			res_string += key + " = " + value
		}
	}
	return res_string
}

//helper functions
//get certain lines function to get certain lines that match a given regular expression
func get_certain_lines(file string, r *regexp.Regexp) []string {
	//fmt.Println(file)
	res := []string{}
	lines := strings.Split(file, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}
	for _, line := range lines {
		is_matching := r.MatchString(line)
		if is_matching {
			res = append(res, line)
		}
	}
	return res
}

// make_key_value_dict function returns a dictionary for certain section[key , value])]
func make_key_value_dict(file string, section string) map[string]string {
	section = strings.TrimSpace(section)
	var res_dict = make(map[string]string)
	lines := strings.Split(file, "\n")
	//triming white spaces for comparisions of strings
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}
	length := len(lines)
	for i := 0; i < length; i++ {
		if strings.HasPrefix(lines[i], "[") {
			if lines[i] == section {
				i++
				for ok := true; ok; ok = i < length && !strings.HasPrefix(lines[i], "[") {
					if !strings.HasPrefix(lines[i], ";") && len(lines[i]) > 1 {
						dictionary_key_value := strings.Split(lines[i], "=")
						dictionary_key_value[0] = strings.TrimSpace(dictionary_key_value[0])
						dictionary_key_value[1] = strings.TrimSpace(dictionary_key_value[1])
						res_dict[dictionary_key_value[0]] = dictionary_key_value[1]
					}
					i++
				}
				return res_dict
			}
		}
	}
	return nil
}

func main() {
	p := Parser{}
	str := p.LoadFromFile("file1.ini")
	p.LoadFromString(str)
	fmt.Printf("%q", p.GetSectionNames())
	fmt.Println()
	fmt.Printf("%q", p.GetSections())
	fmt.Println()
	fmt.Printf("%q", p.Getvalue("[database]", "port"))
	fmt.Println()
	str2 := p.ToString()
	fmt.Println(str2)
	fmt.Println()
	p.SetKeyValue("[database]", "ip", "127.1.1.1")
	str3 := p.ToString()
	fmt.Println(str3)
}
