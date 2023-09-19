package main

import (
	"log"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type DocumentSlice []interface{}

// Implement sort.Interface for DocumentSlice
func (p DocumentSlice) Len() int { return len(p) }
func (p DocumentSlice) Less(i, j int) bool {
	// Handle "null" documents from trailing ---
	if p[i] == nil {
		return true
	} else if p[j] == nil {
		return false
	}

	var l = make(map[string]interface{})
	for k, v := range p[i].(map[string]interface{}) {
		l[k] = v
	}
	var r = make(map[string]interface{})
	for k, v := range p[j].(map[string]interface{}) {
		r[k] = v
	}

	apiLess := l["apiVersion"].(string) < r["apiVersion"].(string)
	kindLess := l["apiVersion"].(string) == r["apiVersion"].(string) && l["kind"].(string) < r["kind"].(string)
	nameLess := l["apiVersion"] == r["apiVersion"] &&
		l["kind"] == r["kind"] &&
		l["metadata"].(map[string]interface{})["name"].(string) < (r["metadata"].(map[string]interface{}))["name"].(string)

	return apiLess || kindLess || nameLess

}

func (p DocumentSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	input := os.Args[1]

	// Read YAML file
	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	var documents DocumentSlice

	decoder := yaml.NewDecoder(file)
	for {
		var doc interface{}
		err = decoder.Decode(&doc)
		if err != nil {
			break
		}
		documents = append(documents, doc)
	}

	if err.Error() != "EOF" {
		log.Fatalf("error: %v", err)
	}

	sort.Sort(documents)

	encoder := yaml.NewEncoder(os.Stdout)
	for _, doc := range documents {
		err := encoder.Encode(&doc)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	if err := encoder.Close(); err != nil {
		log.Fatalf("error %v", err)
	}
}
