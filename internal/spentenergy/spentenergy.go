package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 || weight <= 0 || height <= 0 || steps <= 0 {
		return 0, fmt.Errorf("Некорректные параметры")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	walkingSpentCalories := (weight * meanSpeed * durationInMinutes * walkingCaloriesCoefficient) / minInH
	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 || weight <= 0 || height <= 0 || steps <= 0 {
		return 0, fmt.Errorf("Некорректные параметры")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	runningSpentCalories := (weight * meanSpeed * durationInMinutes) / minInH
	return runningSpentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	meanspeed := Distance(steps, height) / duration.Hours()
	return meanspeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := stepLengthCoefficient * height
	distanceKilometr := (stepLength * float64(steps)) / mInKm
	return distanceKilometr
}
