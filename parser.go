package main

import (
	"regexp"
	"strings"
)

type Parser struct {
	file string
	dict map[string]map[string]string
}

func (p Parser) GetSectionNames() []string {
	r, _ := regexp.Compile(` *\[.*`)
	sections := get_certain_lines(p.file, r)
	//to return section name only without []
	for i := 0; i < len(sections); i++ {
		sections[i] = sections[i][1 : len(sections[i])-1]
	}
	return sections
}

//this function assign the dictionary to parser
func (p Parser) GetSections() map[string]map[string]string {
	var res_dict = make(map[string]map[string]string)
	sections := p.GetSectionNames() //array of section name instance and return it
	for _, section := range sections {
		dict2 := make_dictionary(p.file, section)
		res_dict[section] = dict2
	}
	for k, v := range res_dict {
		p.dict[k] = v
	}
	return res_dict
}
func (p Parser) Getvalue(section, key string) string {
	var dict = make(map[string]string)
	dict = make_dictionary(p.file, section)
	return dict[key]
}
func (p Parser) SetKeyValue(section, key, value string) {
	p.GetSections()
	p.dict[section][key] = value
}

//function to get certain lines that match a given regular expression
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

// helper function returns a dictionary for certain section[key , value])]
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
