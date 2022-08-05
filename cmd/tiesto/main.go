package main

import (
	"encoding/json"
	"log"
	"os"
)

// possible inaccurate

var rawData = []map[string]interface{}{
	{
		"name":      "marco",
		"hometown":  "Quito, EQU",
		"residence": "Redmond, WA",
	},
	{
		"name":      "ben",
		"hometown":  "Raredo, TX",
		"residence": "Boise, ID",
	},
	{
		"name":      "steven",
		"hometown":  "Boston, MA",
		"residence": "Northampton, MA",
	},
	{
		"name":      "russell",
		"hometown":  "Los Angeles, CA",
		"residence": "Seattle, WA",
	},
	{
		"name":      "emma",
		"hometown":  "London, UK",
		"residence": "San Francisco, CA",
	},
	{
		"name":      "ian",
		"hometown":  "Parkersburg, WV",
		"residence": "Atlanta, GA",
	},
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(rawData); err != nil {
		log.Fatalf("failed to encode raw data: %v", err)
	}
}
