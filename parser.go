package parser

import (
	"errors"
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
		err2 := p.LoadFromString(text)
		if err2 != nil {
			return errors.New("error in the syntax of file")
		}
		return nil
	}
	return errors.New("error in loading the file")
}

//  LoadFromString function just reads String and generate the dictionary
func (p *Parser) LoadFromString(Text string) error {
	p.dict = make(map[string]map[string]string)
	error := p.Parsing(Text)
	if error != nil {
		return errors.New("error in the syntax of file")
	}
	return nil
}

//parsing function Generate the dictionary
func (p *Parser) Parsing(text string) error {
	p.dict = make(map[string]map[string]string)
	lines := strings.Split(text, "\n")
	var section_name string
	flag := true
	for _, line := range lines {
		line = strings.TrimSpace(line)
		is_section, _ := regexp.MatchString(`^\[([^]]*)\]\s*$`, line)
		if is_section {
			section_name = line
			//section_name = section_name[1 : len(section_name)-1]
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
		is_comment, _ := regexp.MatchString(`; .*`, line)
		if !is_comment && line != "" {
			flag = false
			break
		}
	}
	if !flag {
		return errors.New("file format is inccorect")
	}
	return nil
}

// SaveToFile function saves ini to file
func (p *Parser) SaveToFile(out_file string) error {
	file, err := os.Create(out_file)
	file.Close()
	if err != nil {
		return errors.New("error in making output file")
	}
	file.WriteString(p.ToString())
	return nil
}

// Get section names function returns a slice "array" of section names
func (p Parser) GetSectionNames() []string {
	res := []string{}
	for section := range p.dict {
		res = append(res, section)
	}
	return res
}

// Get sections return the dictionary
func (p Parser) GetSections() map[string]map[string]string {
	return p.dict
}

//Get function return value for certain section and key
func (p Parser) Get(section, key string) (string, error) {
	str, bool := p.dict[section][key]
	if !bool {
		return str, errors.New("invalid section or key")
	}
	return str, nil
}

//Set functions add another key and value to a certain section
func (p *Parser) Set(section, key, value string) {
	if p.dict[section] == nil {
		temp := make(map[string]string)
		temp[key] = value
		p.dict[section] = temp
	} else {
		p.dict[section][key] = value
	}
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
