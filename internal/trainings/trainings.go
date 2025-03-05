package trainings

import (
	"fmt"
	"strings"
	"time"
	"errors"
	"strconv"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}


// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	
	s := strings.Split(datastring, ",")

	if len(s) == 3 {
		step, err := strconv.Atoi(s[0])
		if err != nil {
			return err
		}
		t.Steps = step

		if s[1] != "Бег" && s[1] != "Ходьба" {
			err := errors.New("неизвестный тип тренировки")
			return err
		}
		t.TrainingType = s[1]

		duration, err := time.ParseDuration(s[2])
		if err != nil {
			return err
		}
		t.Duration = duration
	} else {
		err := errors.New("Invalid arguments count")
		return err
	}
	return nil
}
// создайте метод ActionInfo()
/*1.	Вычислить дистанцию, используя функцию из пакета spentenergy.
2.	Вычислить среднюю скорость, используя функцию из пакета spentenergy.
3.	Проверить, какой вид тренировки содержится в структуре Training. Для каждого из видов тренировки рассчитать калории, используя функцию из пакета spentenergy.
4.	Сформируйте и верните строку, образец которой был выше.
5.	Если был передан неизвестный тип тренировки, верните "неизвестный тип тренировки".
*/
func (t Training) ActionInfo() string {
	var ccal float64

	var s string
	s+=string(t.Steps)
	s+=t.TrainingType
	s+=string(t.Duration)
	err := t.Parse(s)

	if err != nil {
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", err, 0.0, 0.0, 0.0, 0.0)
	}

	dist := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	switch t.TrainingType {
		case "Бег": {
			ccal = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
		}
		case "Ходьба": {
			ccal = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		}
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, speed, ccal)
}

