package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() string
}

// создайте функцию Info()
/*
1.	Перебрать все значения слайса dataset в цикле.
2.	Распарсить каждое знание с помощью метода Parse().
3.	Обработать ошибку парсинга. Если она возникает, нужно перейти к следующей итерации цикла.
4.	Сформировать и вывести строку с информацией об активности с помощью метода ActionInfo().
*/
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			errorBytes := []byte(fmt.Sprintf("%v\n", err))
			fmt.Println(string(errorBytes))
			continue
		}
		fmt.Println(dp.ActionInfo())
		//dp.ActionInfo()
	}
}
