package cache

import (
	"encoding/json"
)

type Movie struct {
	Id string	`json:"id"`
	Title string	`json:"title"`	
	Description string	`json:"description"` 
}

func (m Movie) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Movie) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	return nil
}

