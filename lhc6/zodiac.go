package lhc6

import (
	"fmt"
)

var (
	yearZodiacNumbers = map[int]map[string][]int{
		2020: {
			"rat":     {1, 13, 25, 37, 49},
			"ox":      {12, 24, 36, 48},
			"tiger":   {11, 23, 35, 47},
			"rabbit":  {10, 22, 34, 46},
			"dragon":  {9, 21, 33, 45},
			"snake":   {8, 20, 32, 44},
			"horse":   {7, 19, 31, 43},
			"goat":    {6, 18, 30, 42},
			"monkey":  {5, 17, 29, 41},
			"rooster": {4, 16, 28, 40},
			"dog":     {3, 15, 27, 39},
			"pig":     {2, 14, 26, 38},
		},
		2021: {
			"rat":     {2, 14, 26, 38},
			"ox":      {1, 13, 25, 37, 49},
			"tiger":   {12, 24, 36, 48},
			"rabbit":  {11, 23, 35, 47},
			"dragon":  {10, 22, 34, 46},
			"snake":   {9, 21, 33, 45},
			"horse":   {8, 20, 32, 44},
			"goat":    {7, 19, 31, 43},
			"monkey":  {6, 18, 30, 42},
			"rooster": {5, 17, 29, 41},
			"dog":     {4, 16, 28, 40},
			"pig":     {3, 15, 27, 39},
		},
		2022: {
			"rat":     {3, 15, 27, 39},
			"ox":      {2, 14, 26, 38},
			"tiger":   {1, 13, 25, 37, 49},
			"rabbit":  {12, 24, 36, 48},
			"dragon":  {11, 23, 35, 47},
			"snake":   {10, 22, 34, 46},
			"horse":   {9, 21, 33, 45},
			"goat":    {8, 20, 32, 44},
			"monkey":  {7, 19, 31, 43},
			"rooster": {6, 18, 30, 42},
			"dog":     {5, 17, 29, 41},
			"pig":     {4, 16, 28, 40},
		},
		2023: {
			"rat":     {4, 16, 28, 40},
			"ox":      {3, 15, 27, 39},
			"tiger":   {2, 14, 26, 38},
			"rabbit":  {1, 13, 25, 37, 49},
			"dragon":  {12, 24, 36, 48},
			"snake":   {11, 23, 35, 47},
			"horse":   {10, 22, 34, 46},
			"goat":    {9, 21, 33, 45},
			"monkey":  {8, 20, 32, 44},
			"rooster": {7, 19, 31, 43},
			"dog":     {6, 18, 30, 42},
			"pig":     {5, 17, 29, 41},
		},
		2024: {
			"rat":     {5, 17, 29, 41},
			"ox":      {4, 16, 28, 40},
			"tiger":   {3, 15, 27, 39},
			"rabbit":  {2, 14, 26, 38},
			"dragon":  {1, 13, 25, 37, 49},
			"snake":   {12, 24, 36, 48},
			"horse":   {11, 23, 35, 47},
			"goat":    {10, 22, 34, 46},
			"monkey":  {9, 21, 33, 45},
			"rooster": {8, 20, 32, 44},
			"dog":     {7, 19, 31, 43},
			"pig":     {6, 18, 30, 42},
		},
		2025: {
			"rat":     {6, 18, 30, 42},
			"ox":      {5, 17, 29, 41},
			"tiger":   {4, 16, 28, 40},
			"rabbit":  {3, 15, 27, 39},
			"dragon":  {2, 14, 26, 38},
			"snake":   {1, 13, 25, 37, 49},
			"horse":   {12, 24, 36, 48},
			"goat":    {11, 23, 35, 47},
			"monkey":  {10, 22, 34, 46},
			"rooster": {9, 21, 33, 45},
			"dog":     {8, 20, 32, 44},
			"pig":     {7, 19, 31, 43},
		},
		2026: {
			"rat":     {7, 19, 31, 43},
			"ox":      {6, 18, 30, 42},
			"tiger":   {5, 17, 29, 41},
			"rabbit":  {4, 16, 28, 40},
			"dragon":  {3, 15, 27, 39},
			"snake":   {2, 14, 26, 38},
			"horse":   {1, 13, 25, 37, 49},
			"goat":    {12, 24, 36, 48},
			"monkey":  {11, 23, 35, 47},
			"rooster": {10, 22, 34, 46},
			"dog":     {9, 21, 33, 45},
			"pig":     {8, 20, 32, 44},
		},
		2027: {
			"rat":     {8, 20, 32, 44},
			"ox":      {7, 19, 31, 43},
			"tiger":   {6, 18, 30, 42},
			"rabbit":  {5, 17, 29, 41},
			"dragon":  {4, 16, 28, 40},
			"snake":   {3, 15, 27, 39},
			"horse":   {2, 14, 26, 38},
			"goat":    {1, 13, 25, 37, 49},
			"monkey":  {12, 24, 36, 48},
			"rooster": {11, 23, 35, 47},
			"dog":     {10, 22, 34, 46},
			"pig":     {9, 21, 33, 45},
		},
		2028: {
			"rat":     {9, 21, 33, 45},
			"ox":      {8, 20, 32, 44},
			"tiger":   {7, 19, 31, 43},
			"rabbit":  {6, 18, 30, 42},
			"dragon":  {5, 17, 29, 41},
			"snake":   {4, 16, 28, 40},
			"horse":   {3, 15, 27, 39},
			"goat":    {2, 14, 26, 38},
			"monkey":  {1, 13, 25, 37, 49},
			"rooster": {12, 24, 36, 48},
			"dog":     {11, 23, 35, 47},
			"pig":     {10, 22, 34, 46},
		},
		2029: {
			"rat":     {10, 22, 34, 46},
			"ox":      {9, 21, 33, 45},
			"tiger":   {8, 20, 32, 44},
			"rabbit":  {7, 19, 31, 43},
			"dragon":  {6, 18, 30, 42},
			"snake":   {5, 17, 29, 41},
			"horse":   {4, 16, 28, 40},
			"goat":    {3, 15, 27, 39},
			"monkey":  {2, 14, 26, 38},
			"rooster": {1, 13, 25, 37, 49},
			"dog":     {12, 24, 36, 48},
			"pig":     {11, 23, 35, 47},
		},
		2030: {
			"rat":     {11, 23, 35, 47},
			"ox":      {10, 22, 34, 46},
			"tiger":   {9, 21, 33, 45},
			"rabbit":  {8, 20, 32, 44},
			"dragon":  {7, 19, 31, 43},
			"snake":   {6, 18, 30, 42},
			"horse":   {5, 17, 29, 41},
			"goat":    {4, 16, 28, 40},
			"monkey":  {3, 15, 27, 39},
			"rooster": {2, 14, 26, 38},
			"dog":     {1, 13, 25, 37, 49},
			"pig":     {12, 24, 36, 48},
		},
		2031: {
			"rat":     {12, 24, 36, 48},
			"ox":      {11, 23, 35, 47},
			"tiger":   {10, 22, 34, 46},
			"rabbit":  {9, 21, 33, 45},
			"dragon":  {8, 20, 32, 44},
			"snake":   {7, 19, 31, 43},
			"horse":   {6, 18, 30, 42},
			"goat":    {5, 17, 29, 41},
			"monkey":  {4, 16, 28, 40},
			"rooster": {3, 15, 27, 39},
			"dog":     {2, 14, 26, 38},
			"pig":     {1, 13, 25, 37, 49},
		},
		2032: {
			"rat":     {1, 13, 25, 37, 49},
			"ox":      {12, 24, 36, 48},
			"tiger":   {11, 23, 35, 47},
			"rabbit":  {10, 22, 34, 46},
			"dragon":  {9, 21, 33, 45},
			"snake":   {8, 20, 32, 44},
			"horse":   {7, 19, 31, 43},
			"goat":    {6, 18, 30, 42},
			"monkey":  {5, 17, 29, 41},
			"rooster": {4, 16, 28, 40},
			"dog":     {3, 15, 27, 39},
			"pig":     {2, 14, 26, 38},
		},
		2033: {
			"rat":     {2, 14, 26, 38},
			"ox":      {1, 13, 25, 37, 49},
			"tiger":   {12, 24, 36, 48},
			"rabbit":  {11, 23, 35, 47},
			"dragon":  {10, 22, 34, 46},
			"snake":   {9, 21, 33, 45},
			"horse":   {8, 20, 32, 44},
			"goat":    {7, 19, 31, 43},
			"monkey":  {6, 18, 30, 42},
			"rooster": {5, 17, 29, 41},
			"dog":     {4, 16, 28, 40},
			"pig":     {3, 15, 27, 39},
		},
		2034: {
			"rat":     {3, 15, 27, 39},
			"ox":      {2, 14, 26, 38},
			"tiger":   {1, 13, 25, 37, 49},
			"rabbit":  {12, 24, 36, 48},
			"dragon":  {11, 23, 35, 47},
			"snake":   {10, 22, 34, 46},
			"horse":   {9, 21, 33, 45},
			"goat":    {8, 20, 32, 44},
			"monkey":  {7, 19, 31, 43},
			"rooster": {6, 18, 30, 42},
			"dog":     {5, 17, 29, 41},
			"pig":     {4, 16, 28, 40},
		},
		2035: {
			"rat":     {4, 16, 28, 40},
			"ox":      {3, 15, 27, 39},
			"tiger":   {2, 14, 26, 38},
			"rabbit":  {1, 13, 25, 37, 49},
			"dragon":  {12, 24, 36, 48},
			"snake":   {11, 23, 35, 47},
			"horse":   {10, 22, 34, 46},
			"goat":    {9, 21, 33, 45},
			"monkey":  {8, 20, 32, 44},
			"rooster": {7, 19, 31, 43},
			"dog":     {6, 18, 30, 42},
			"pig":     {5, 17, 29, 41},
		},
		2036: {
			"rat":     {5, 17, 29, 41},
			"ox":      {4, 16, 28, 40},
			"tiger":   {3, 15, 27, 39},
			"rabbit":  {2, 14, 26, 38},
			"dragon":  {1, 13, 25, 37, 49},
			"snake":   {12, 24, 36, 48},
			"horse":   {11, 23, 35, 47},
			"goat":    {10, 22, 34, 46},
			"monkey":  {9, 21, 33, 45},
			"rooster": {8, 20, 32, 44},
			"dog":     {7, 19, 31, 43},
			"pig":     {6, 18, 30, 42},
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
