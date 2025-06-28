package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, datastring := range dataset {
		err := dp.Parse(datastring)
		if err != nil {
			log.Printf("Ошибка парсинга %s: %v", datastring, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации об активности: %v", err)
		} else {
			fmt.Println(info)
		}
	}
}
