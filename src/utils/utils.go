package utils

import (
	"encoding/json"

	"github.com/google/uuid"
)

func UnmarshalJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func ParseUUID(uuidStr string) (uuid.UUID, error) {
	uuidValue, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuidValue, nil
}
