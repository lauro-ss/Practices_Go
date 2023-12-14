package service

import (
	"httpserver/model"
	"os"
	"strconv"
	"strings"
)

const PATH = "./data/"

func GetAllAnimalCsv(f string) ([]model.Animal, error) {
	b, err := os.ReadFile(PATH + "animal.csv")
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(b), "\n")[1:]

	o := make([]model.Animal, len(s))
	for i, v := range s {
		c := strings.Split(v, ";")
		o[i].Id, err = strconv.Atoi(c[0])
		if err != nil {
			return nil, err
		}
		o[i].Name = c[1]
		o[i].Icon = c[2]
	}

	return o, nil
}
