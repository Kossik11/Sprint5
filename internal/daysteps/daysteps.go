package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные о дневной прогулке.
type DaySteps struct {
	Steps                 int           // количество шагов
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // персональные данные пользователя
}

// Parse разбирает строку с данными о прогулке формата "678,0h50m" и заполняет поля структуры.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("некорректный формат данных")
	}

	// Парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("некорректное значение шагов")
	}
	if steps <= 0 {
		return fmt.Errorf("шаги должны быть положительным числом")
	}
	ds.Steps = steps

	// Парсим длительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("некорректный формат продолжительности")
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

// ActionInfo возвращает строку с информацией о прогулке и ошибку, если она возникает.
func (ds DaySteps) ActionInfo() (string, error) {
	// Проверяем валидность данных
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Personal.Weight <= 0 || ds.Personal.Height <= 0 {
		return "", fmt.Errorf("некорректные параметры прогулки")
	}

	// Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	// Вычисляем калории
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формируем строку с информацией
	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)

	return info, nil
}
