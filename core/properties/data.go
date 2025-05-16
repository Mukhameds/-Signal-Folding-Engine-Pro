package core

import (
	"encoding/json"
	
	"os"
)

type AAProperty struct {
	Name        string  `json:"name"`
	Mass        float64 `json:"mass"`
	Charge      float64 `json:"charge"`
	Hydrophobic float64 `json:"hydrophobic"`
	Volume      float64 `json:"volume"`
}

var AAProperties map[string]AAProperty

func LoadAAProperties(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&AAProperties)
}

func GetProperty(code string) AAProperty {
	if prop, ok := AAProperties[code]; ok {
		return prop
	}
	return AAProperty{} // пустой fallback
}
