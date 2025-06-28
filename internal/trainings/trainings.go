package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int                   // количество шагов, проделанных за тренировку.
	TrainingType string                // тип тренировки(бег или ходьба).
	Duration     time.Duration         // длительность тренировки.
	Personal     personaldata.Personal // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("некорректный формат данных")
	}

	// Преобразуем количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("некорректное значение шагов")
	}
	if steps <= 0 {
		return fmt.Errorf("шаги должны быть положительным числом")
	}
	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = parts[1]

	// Преобразуем продолжительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("некорректный формат продолжительности")
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	t.Duration = duration

	return nil
}

// ActionInfo возвращает строку с информацией о тренировке и ошибку, если она возникает.
func (t Training) ActionInfo() (string, error) {
	// Проверяем валидность данных
	if t.Steps <= 0 || t.Duration <= 0 || t.Personal.Weight <= 0 || t.Personal.Height <= 0 {
		return "", fmt.Errorf("некорректные параметры тренировки")
	}

	// Вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)

	// Вычисляем среднюю скорость
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	// Вычисляем калории в зависимости от типа тренировки
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	// Формируем строку с информацией о тренировке
	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, avgSpeed, calories)

	return info, nil
}
