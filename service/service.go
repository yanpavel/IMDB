package service

import (
	"fmt"
	"os"

	s "github.com/yanpavel/imdb/storage"
)

type Service struct {
	repository s.MemTable
}

func NewService(repo s.MemTable) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Process() {
	file, err := os.Open("../commitlog.log")
	if err != nil {
		file, err = os.Create("../commitlog.log")
		if err != nil {
			panic("unable to create log file")
		}
	}

	s.transferData()

	operation := make([]string, 3)

	for {
		firstMenu()

		var firstStep uint8
		fmt.Scan(&firstStep)

		var id string
		var name string

		switch firstStep {
		case 1:
			fmt.Println("Введите ID и имя персоны для добавления:")
			fmt.Scan(&id, &name)
			operation = []string{"1", id, name}
			//s.repository.Add(id, name)
		case 2:
			fmt.Println("Введите ID и имя персоны для обновления записи:")
			fmt.Scan(&id, &name)
			operation = []string{"2", id, name}
			//s.repository.Change(id, &name)
		case 3:
			fmt.Println("Введите ID персоны для удаления:")
			fmt.Scan(&id)
			operation = []string{"3", id}
			//s.repository.Delete(id)
		default:
			fmt.Print("Невалидный ввод")
			continue
		}

		writeTo(file, operation)
		s.transferData()
	}
}

func (s *Service) transferData() error {
	// Прочитать журнал после отметки
	// если есть записи выполнить команду с репозиторием
	return nil
}

func writeTo(file *os.File, operation []string) error {
	// for k, o := range operation {
	// 	if k[0] != "" {
	// 		file.WriteString(o[0] + " " + o[1] + " " + o[2] + "\n")
	// 		continue
	// 	}
	// 	file.WriteString(o[0] + " " + o[1] + "\n")
	// }
	return nil
}

func firstMenu() {
	fmt.Println("Выберите команду:")
	fmt.Println("1 - Добавить персону, 2 - Изменить персону, 3 - Удалить персону")
}
