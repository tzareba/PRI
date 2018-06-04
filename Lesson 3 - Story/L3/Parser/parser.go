package parser

import (
	"encoding/json"
	"os"

	"github.com/gophercises/L3/Model"
)

type Provider = func(string) (model.Story, bool)

func CreateProvider(filePath string) Provider {
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var stories map[string]model.Story
	json.NewDecoder(reader).Decode(&stories)

	return func(arcName string) (model.Story, bool) {
		story, ok := stories[arcName]
		return story, ok
	}
}
