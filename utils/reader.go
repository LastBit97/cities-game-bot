package utils

import (
	"encoding/json"
	"io"
	"os"
)

func ReadAndUnmarshal(filename string, v any) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(data, v)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}
