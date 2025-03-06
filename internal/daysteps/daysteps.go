package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {

	s := strings.Split(datastring, ",")

	if len(s) == 2 {
		step, err := strconv.Atoi(s[0])

		if err != nil {
			return err
		}
		ds.Steps = step

		duration, err := time.ParseDuration(s[1])

		if err != nil {
			return err
		}
		ds.Duration = duration
	} else {
		err := errors.New("Invalid arguments count")
		return err
	}
	return nil
}

// создайте метод ActionInfo()
/*1.	Вычислите дистанцию.
2.	Вычислите количество сожжённых калорий.
3.	Сформируйте и верните строку с информацией.
*/

func (ds DaySteps) ActionInfo() string {
    dist := spentenergy.Distance(ds.Steps)
    ccal := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

    return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, ccal)
}


/*func (ds DaySteps) ActionInfo() string {
	var s string
	s+=string(ds.Steps)
	s+=string(ds.Duration)
	err := ds.Parse(s)
	if err != nil {
		errorBytes := []byte(fmt.Sprintf("%v\n", err))
		return string(errorBytes)
	}

	dist := float64(ds.Steps) * StepLength / spentenergy.MInKm

	ccal := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, ccal)
}*/
