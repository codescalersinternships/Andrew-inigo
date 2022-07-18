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
func (p *Parser) LoadFromFile(file string) {
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error in loading file")
	}
	Text := string(dat)
	p.text = Text
}

// Load from string function just takes a string and add it to the text attribute
func (p *Parser) LoadFromString(Text string) {
	p.text = Text
}

// Get section names function returns a slice "array" of section names
func (p Parser) GetSectionNames() []string {
	r, _ := regexp.Compile(` *\[.*`)
	sections := get_certain_lines(p.text, r)
	//to return section name only without []
	for i := 0; i < len(sections); i++ {
		sections[i] = sections[i][1 : len(sections[i])-1]
	}
	return sections
}

//Get sections function returns dictionary containg [section , [key , value]]
//and assigns this dictionary to dict attribute in parser
func (p Parser) GetSections() map[string]map[string]string {
	var res_dict = make(map[string]map[string]string)
	sections := p.GetSectionNames() //array of section name instance and return it
	for _, section := range sections {
		dict2 := make_dictionary(p.text, section)
		res_dict[section] = dict2
	}
	for k, v := range res_dict {
		p.dict[k] = v
	}
	return res_dict
}

//Getvalue function return value for certain section and key
func (p Parser) Getvalue(section, key string) string {
	var dict = make(map[string]string)
	dict = make_dictionary(p.text, section)
	return dict[key]
}

//setkeyvalue functions add another key and value to a certain section
func (p Parser) SetKeyValue(section, key, value string) {
	p.GetSections()
	p.dict[section][key] = value
}

//get certain lines function to get certain lines that match a given regular expression
func get_certain_lines(file string, r *regexp.Regexp) []string {
	res := []string{}
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		is_matching := r.MatchString(line)
		if is_matching {
			res = append(res, line)
		}
	}
	return res
}

// make_dictionary function returns a dictionary for certain section[key , value])]
func make_dictionary(file string, section string) map[string]string {
	var res_dict = make(map[string]string)
	lines := strings.Split(file, "\n")
	length := len(lines)
	for i := 0; i < length; i++ {
		if strings.HasPrefix(lines[i], "[") {
			if lines[i][1:len(lines[i])-1] == section {
				i++
				for ok := true; ok; ok = i < length && !strings.HasPrefix(lines[i], "[") {
					if !strings.HasPrefix(lines[i], ";") && lines[i] != "" {
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
	p.LoadFromFile("file1.INI")
}
