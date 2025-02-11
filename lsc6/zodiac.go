package lsc6

import (
	"fmt"
)

var (
	yearZodiacNumbers = map[int]map[string][]int{
		2019: {
			"rat":     {12, 24, 36, 48, 60},
			"ox":      {11, 23, 35, 47, 59},
			"tiger":   {10, 22, 34, 46, 58},
			"rabbit":  {9, 21, 33, 45, 57},
			"dragon":  {8, 20, 32, 44, 56},
			"snake":   {7, 19, 31, 43, 55},
			"horse":   {6, 18, 30, 42, 54},
			"goat":    {5, 17, 29, 41, 53},
			"monkey":  {4, 16, 28, 40, 52},
			"rooster": {3, 15, 27, 39, 51},
			"dog":     {2, 14, 26, 38, 50},
			"pig":     {1, 13, 25, 37, 49},
		},
		2020: {
			"rat":     {1, 13, 25, 37, 49},
			"ox":      {12, 24, 36, 48, 60},
			"tiger":   {11, 23, 35, 47, 59},
			"rabbit":  {10, 22, 34, 46, 58},
			"dragon":  {9, 21, 33, 45, 57},
			"snake":   {8, 20, 32, 44, 56},
			"horse":   {7, 19, 31, 43, 55},
			"goat":    {6, 18, 30, 42, 54},
			"monkey":  {5, 17, 29, 41, 53},
			"rooster": {4, 16, 28, 40, 52},
			"dog":     {3, 15, 27, 39, 51},
			"pig":     {2, 14, 26, 38, 50},
		},
		2021: {
			"rat":     {2, 14, 26, 38, 50},
			"ox":      {1, 13, 25, 37, 49},
			"tiger":   {12, 24, 36, 48, 60},
			"rabbit":  {11, 23, 35, 47, 59},
			"dragon":  {10, 22, 34, 46, 58},
			"snake":   {9, 21, 33, 45, 57},
			"horse":   {8, 20, 32, 44, 56},
			"goat":    {7, 19, 31, 43, 55},
			"monkey":  {6, 18, 30, 42, 54},
			"rooster": {5, 17, 29, 41, 53},
			"dog":     {4, 16, 28, 40, 52},
			"pig":     {3, 15, 27, 39, 51},
		},
		2022: {
			"rat":     {3, 15, 27, 39, 51},
			"ox":      {2, 14, 26, 38, 50},
			"tiger":   {1, 13, 25, 37, 49},
			"rabbit":  {12, 24, 36, 48, 60},
			"dragon":  {11, 23, 35, 47, 59},
			"snake":   {10, 22, 34, 46, 58},
			"horse":   {9, 21, 33, 45, 57},
			"goat":    {8, 20, 32, 44, 56},
			"monkey":  {7, 19, 31, 43, 55},
			"rooster": {6, 18, 30, 42, 54},
			"dog":     {5, 17, 29, 41, 53},
			"pig":     {4, 16, 28, 40, 52},
		},
		2023: {
			"rat":     {4, 16, 28, 40, 52},
			"ox":      {3, 15, 27, 39, 51},
			"tiger":   {2, 14, 26, 38, 50},
			"rabbit":  {1, 13, 25, 37, 49},
			"dragon":  {12, 24, 36, 48, 60},
			"snake":   {11, 23, 35, 47, 59},
			"horse":   {10, 22, 34, 46, 58},
			"goat":    {9, 21, 33, 45, 57},
			"monkey":  {8, 20, 32, 44, 56},
			"rooster": {7, 19, 31, 43, 55},
			"dog":     {6, 18, 30, 42, 54},
			"pig":     {5, 17, 29, 41, 53},
		},
		2024: {
			"rat":     {5, 17, 29, 41, 53},
			"ox":      {4, 16, 28, 40, 52},
			"tiger":   {3, 15, 27, 39, 51},
			"rabbit":  {2, 14, 26, 38, 50},
			"dragon":  {1, 13, 25, 37, 49},
			"snake":   {12, 24, 36, 48, 60},
			"horse":   {11, 23, 35, 47, 59},
			"goat":    {10, 22, 34, 46, 58},
			"monkey":  {9, 21, 33, 45, 57},
			"rooster": {8, 20, 32, 44, 56},
			"dog":     {7, 19, 31, 43, 55},
			"pig":     {6, 18, 30, 42, 54},
		},
		2025: {
			"rat":     {6, 18, 30, 42, 54},
			"ox":      {5, 17, 29, 41, 53},
			"tiger":   {4, 16, 28, 40, 52},
			"rabbit":  {3, 15, 27, 39, 51},
			"dragon":  {2, 14, 26, 38, 50},
			"snake":   {1, 13, 25, 37, 49},
			"horse":   {12, 24, 36, 48, 60},
			"goat":    {11, 23, 35, 47, 59},
			"monkey":  {10, 22, 34, 46, 58},
			"rooster": {9, 21, 33, 45, 57},
			"dog":     {8, 20, 32, 44, 56},
			"pig":     {7, 19, 31, 43, 55},
		},
		2026: {
			"rat":     {7, 19, 31, 43, 55},
			"ox":      {6, 18, 30, 42, 54},
			"tiger":   {5, 17, 29, 41, 53},
			"rabbit":  {4, 16, 28, 40, 52},
			"dragon":  {3, 15, 27, 39, 51},
			"snake":   {2, 14, 26, 38, 50},
			"horse":   {1, 13, 25, 37, 49},
			"goat":    {12, 24, 36, 48, 60},
			"monkey":  {11, 23, 35, 47, 59},
			"rooster": {10, 22, 34, 46, 58},
			"dog":     {9, 21, 33, 45, 57},
			"pig":     {8, 20, 32, 44, 56},
		},
		2027: {
			"rat":     {8, 20, 32, 44, 56},
			"ox":      {7, 19, 31, 43, 55},
			"tiger":   {6, 18, 30, 42, 54},
			"rabbit":  {5, 17, 29, 41, 53},
			"dragon":  {4, 16, 28, 40, 52},
			"snake":   {3, 15, 27, 39, 51},
			"horse":   {2, 14, 26, 38, 50},
			"goat":    {1, 13, 25, 37, 49},
			"monkey":  {12, 24, 36, 48, 60},
			"rooster": {11, 23, 35, 47, 59},
			"dog":     {10, 22, 34, 46, 58},
			"pig":     {9, 21, 33, 45, 57},
		},
		2028: {
			"rat":     {9, 21, 33, 45, 57},
			"ox":      {8, 20, 32, 44, 56},
			"tiger":   {7, 19, 31, 43, 55},
			"rabbit":  {6, 18, 30, 42, 54},
			"dragon":  {5, 17, 29, 41, 53},
			"snake":   {4, 16, 28, 40, 52},
			"horse":   {3, 15, 27, 39, 51},
			"goat":    {2, 14, 26, 38, 50},
			"monkey":  {1, 13, 25, 37, 49},
			"rooster": {12, 24, 36, 48, 60},
			"dog":     {11, 23, 35, 47, 59},
			"pig":     {10, 22, 34, 46, 58},
		},
		2029: {
			"rat":     {10, 22, 34, 46, 58},
			"ox":      {9, 21, 33, 45, 57},
			"tiger":   {8, 20, 32, 44, 56},
			"rabbit":  {7, 19, 31, 43, 55},
			"dragon":  {6, 18, 30, 42, 54},
			"snake":   {5, 17, 29, 41, 53},
			"horse":   {4, 16, 28, 40, 52},
			"goat":    {3, 15, 27, 39, 51},
			"monkey":  {2, 14, 26, 38, 50},
			"rooster": {1, 13, 25, 37, 49},
			"dog":     {12, 24, 36, 48, 60},
			"pig":     {11, 23, 35, 47, 59},
		},
		2030: {
			"rat":     {11, 23, 35, 47, 59},
			"ox":      {10, 22, 34, 46, 58},
			"tiger":   {9, 21, 33, 45, 57},
			"rabbit":  {8, 20, 32, 44, 56},
			"dragon":  {7, 19, 31, 43, 55},
			"snake":   {6, 18, 30, 42, 54},
			"horse":   {5, 17, 29, 41, 53},
			"goat":    {4, 16, 28, 40, 52},
			"monkey":  {3, 15, 27, 39, 51},
			"rooster": {2, 14, 26, 38, 50},
			"dog":     {1, 13, 25, 37, 49},
			"pig":     {12, 24, 36, 48, 60},
		},
		2031: {
			"rat":     {12, 24, 36, 48, 60},
			"ox":      {11, 23, 35, 47, 59},
			"tiger":   {10, 22, 34, 46, 58},
			"rabbit":  {9, 21, 33, 45, 57},
			"dragon":  {8, 20, 32, 44, 56},
			"snake":   {7, 19, 31, 43, 55},
			"horse":   {6, 18, 30, 42, 54},
			"goat":    {5, 17, 29, 41, 53},
			"monkey":  {4, 16, 28, 40, 52},
			"rooster": {3, 15, 27, 39, 51},
			"dog":     {2, 14, 26, 38, 50},
			"pig":     {1, 13, 25, 37, 49},
		},
		2032: {
			"rat":     {1, 13, 25, 37, 49},
			"ox":      {12, 24, 36, 48, 60},
			"tiger":   {11, 23, 35, 47, 59},
			"rabbit":  {10, 22, 34, 46, 58},
			"dragon":  {9, 21, 33, 45, 57},
			"snake":   {8, 20, 32, 44, 56},
			"horse":   {7, 19, 31, 43, 55},
			"goat":    {6, 18, 30, 42, 54},
			"monkey":  {5, 17, 29, 41, 53},
			"rooster": {4, 16, 28, 40, 52},
			"dog":     {3, 15, 27, 39, 51},
			"pig":     {2, 14, 26, 38, 50},
		},
		2033: {
			"rat":     {2, 14, 26, 38, 50},
			"ox":      {1, 13, 25, 37, 49},
			"tiger":   {12, 24, 36, 48, 60},
			"rabbit":  {11, 23, 35, 47, 59},
			"dragon":  {10, 22, 34, 46, 58},
			"snake":   {9, 21, 33, 45, 57},
			"horse":   {8, 20, 32, 44, 56},
			"goat":    {7, 19, 31, 43, 55},
			"monkey":  {6, 18, 30, 42, 54},
			"rooster": {5, 17, 29, 41, 53},
			"dog":     {4, 16, 28, 40, 52},
			"pig":     {3, 15, 27, 39, 51},
		},
		2034: {
			"rat":     {3, 15, 27, 39, 51},
			"ox":      {2, 14, 26, 38, 50},
			"tiger":   {1, 13, 25, 37, 49},
			"rabbit":  {12, 24, 36, 48, 60},
			"dragon":  {11, 23, 35, 47, 59},
			"snake":   {10, 22, 34, 46, 58},
			"horse":   {9, 21, 33, 45, 57},
			"goat":    {8, 20, 32, 44, 56},
			"monkey":  {7, 19, 31, 43, 55},
			"rooster": {6, 18, 30, 42, 54},
			"dog":     {5, 17, 29, 41, 53},
			"pig":     {4, 16, 28, 40, 52},
		},
		2035: {
			"rat":     {4, 16, 28, 40, 52},
			"ox":      {3, 15, 27, 39, 51},
			"tiger":   {2, 14, 26, 38, 50},
			"rabbit":  {1, 13, 25, 37, 49},
			"dragon":  {12, 24, 36, 48, 60},
			"snake":   {11, 23, 35, 47, 59},
			"horse":   {10, 22, 34, 46, 58},
			"goat":    {9, 21, 33, 45, 57},
			"monkey":  {8, 20, 32, 44, 56},
			"rooster": {7, 19, 31, 43, 55},
			"dog":     {6, 18, 30, 42, 54},
			"pig":     {5, 17, 29, 41, 53},
		},
		2036: {
			"rat":     {5, 17, 29, 41, 53},
			"ox":      {4, 16, 28, 40, 52},
			"tiger":   {3, 15, 27, 39, 51},
			"rabbit":  {2, 14, 26, 38, 50},
			"dragon":  {1, 13, 25, 37, 49},
			"snake":   {12, 24, 36, 48, 60},
			"horse":   {11, 23, 35, 47, 59},
			"goat":    {10, 22, 34, 46, 58},
			"monkey":  {9, 21, 33, 45, 57},
			"rooster": {8, 20, 32, 44, 56},
			"dog":     {7, 19, 31, 43, 55},
			"pig":     {6, 18, 30, 42, 54},
		},
	}
	yearNumberZodiacs = map[int]map[int]string{}
)

func init() {
	// init yearNumberZodiacs
	for year, zodiacNumbers := range yearZodiacNumbers {
		numberZodiacs := map[int]string{}
		for zodiac, numbers := range zodiacNumbers {
			for _, number := range numbers {
				numberZodiacs[number] = zodiac
			}
		}

		yearNumberZodiacs[year] = numberZodiacs
	}
}

// GetZodiacNumbers 取得全部生肖的號碼
func GetZodiacNumbers(year int) (map[string][]int, error) {
	zodiacNumbers, ok := yearZodiacNumbers[year]
	if !ok {
		return nil, fmt.Errorf("not supported for %d", year)
	}

	return zodiacNumbers, nil
}

// GetZodiacNumber 取得單一生肖的號碼
func GetZodiacNumber(year int, zodiac string) ([]int, error) {
	zodiacNumbers, err := GetZodiacNumbers(year)
	if err != nil {
		return nil, err
	}

	numbers, ok := zodiacNumbers[zodiac]
	if !ok {
		return nil, fmt.Errorf("invalid zodiac %s", zodiac)
	}

	return numbers, nil
}

// GetNumberZodiacs 取得全部號碼的生肖
func GetNumberZodiacs(year int) (map[int]string, error) {
	numberZodiacs, ok := yearNumberZodiacs[year]
	if !ok {
		return nil, fmt.Errorf("not supported for %d", year)
	}

	return numberZodiacs, nil
}

// GetNumberZodiac 取得單一號碼的生肖
func GetNumberZodiac(yaer, number int) (string, error) {
	numberZodiacs, err := GetNumberZodiacs(yaer)
	if err != nil {
		return "", err
	}

	zodiac, ok := numberZodiacs[number]
	if !ok {
		return "", fmt.Errorf("invalid number %d", number)
	}

	return zodiac, nil
}

// GetDutyZodiac 取得值年生肖
func GetDutyZodiac(year int) (string, error) {
	numberZodiacs, err := GetNumberZodiacs(year)
	if err != nil {
		return "", err
	}

	return numberZodiacs[1], nil
}
