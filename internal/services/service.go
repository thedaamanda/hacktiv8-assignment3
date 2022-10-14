package services

import (
	"assignment3/internal/entity"
	"assignment3/internal/helpers"
	"encoding/json"
	"os"
)

func GetFiles() string {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files := basePath + "/internal/repository/data.json"

	return files
}

func Randomize(data *entity.Data) {
	data.Status.Water = helpers.RandomInt()
	data.Status.Wind = helpers.RandomInt()
}

func ReadJSON(files string) (entity.Data, error) {
	content, err := os.ReadFile(files)

	if err != nil {
		return entity.Data{}, err
	}

	var data entity.Data
	err = json.Unmarshal(content, &data)

	if err != nil {
		return entity.Data{}, err
	}

	return data, nil
}

func WriteJSON(data entity.Data, files string) error {
	content, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = os.WriteFile(files, content, 0644)

	if err != nil {
		return err
	}

	return nil
}
