package main

import (
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	dict map[string]map[string]string
}

// LoadFromFile function just reads the file context and generate the dictionary
func (p *Parser) LoadFromFile(file string) error {
	dat, err := os.ReadFile(file)
	if err == nil {
		text := string(dat)
		p.LoadFromString(text)
		return nil
	}
	return err
}

//  LoadFromString function just reads String and generate the dictionary
func (p *Parser) LoadFromString(Text string) {
	p.dict = make(map[string]map[string]string)
	p.Parsing(Text)
}

//parsing function Generate the dictionary
func (p *Parser) Parsing(text string) {
	p.dict = make(map[string]map[string]string)
	lines := strings.Split(text, "\n")
	var section_name string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		is_section, _ := regexp.MatchString(`^\[([^]]*)\]\s*$`, line)
		if is_section {
			section_name = line
			p.dict[section_name] = make(map[string]string)
			continue
		}
		is_key_val, _ := regexp.MatchString(`^(\w*)\s*=\s*(.*?)\s*$`, line)
		if is_key_val {
			key_value := strings.Split(line, "=")
			key_value[0] = strings.TrimSpace(key_value[0])
			key_value[1] = strings.TrimSpace(key_value[1])
			p.dict[section_name][key_value[0]] = key_value[1]
			continue
		}
	}
}

// SaveToFile function saves ini to file
func (p *Parser) SaveToFile(out_file string) {
	file, _ := os.Create(out_file)
	file.WriteString(p.ToString())
	file.Close()
}

// Get section names function returns a slice "array" of section names
func (p Parser) GetSectionNames() []string {
	res := []string{}
	for section := range p.dict {
		res = append(res, section)
	}
	return res
}

func (p Parser) GetSections() map[string]map[string]string {
	return p.dict
}

//Get function return value for certain section and key
func (p Parser) Get(section, key string) string {
	return p.dict[section][key]
}

//setkeyvalue functions add another key and value to a certain section
func (p *Parser) Set(section, key, value string) {
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
