package lhc6

import (
	"fmt"
)

var (
	yearWuxingNumbers = map[int]map[string][]int{
		2020: {
			"metal": {6, 7, 20, 21, 28, 29, 36, 37},
			"wood":  {2, 3, 10, 11, 18, 19, 32, 33, 40, 41, 48, 49},
			"water": {8, 9, 16, 17, 24, 25, 38, 39, 46, 47},
			"fire":  {4, 5, 12, 13, 26, 27, 34, 35, 42, 43},
			"earth": {1, 14, 15, 22, 23, 30, 31, 44, 45},
		},
		2021: {
			"metal": {7, 8, 21, 22, 29, 30, 37, 38},
			"wood":  {3, 4, 11, 12, 19, 20, 33, 34, 41, 42, 49},
			"water": {9, 10, 17, 18, 25, 26, 39, 40, 47, 48},
			"fire":  {5, 6, 13, 14, 27, 28, 35, 36, 43, 44},
			"earth": {1, 2, 15, 16, 23, 24, 31, 32, 45, 46},
		},
		2022: {
			"metal": {1, 8, 9, 22, 23, 30, 31, 38, 39},
			"wood":  {4, 5, 12, 13, 20, 21, 34, 35, 42, 43},
			"water": {10, 11, 18, 19, 26, 27, 40, 41, 48, 49},
			"fire":  {6, 7, 14, 15, 28, 29, 36, 37, 44, 45},
			"earth": {2, 3, 16, 17, 24, 25, 32, 33, 46, 47},
		},
		2023: {
			"metal": {1, 2, 9, 10, 23, 24, 31, 32, 39, 40},
			"wood":  {5, 6, 13, 14, 21, 22, 35, 36, 43, 44},
			"water": {11, 12, 19, 20, 27, 28, 41, 42, 49},
			"fire":  {7, 8, 15, 16, 29, 30, 37, 38, 45, 46},
			"earth": {3, 4, 17, 18, 25, 26, 33, 34, 47, 48},
		},
		2024: {
			"metal": {2, 3, 10, 11, 24, 25, 32, 33, 40, 41},
			"wood":  {6, 7, 14, 15, 22, 23, 36, 37, 44, 45},
			"water": {12, 13, 20, 21, 28, 29, 42, 43},
			"fire":  {1, 8, 9, 16, 17, 30, 31, 38, 39, 46, 47},
			"earth": {4, 5, 18, 19, 26, 27, 34, 35, 48, 49},
		},
		2025: {
			"metal": {3, 4, 11, 12, 25, 26, 33, 34, 41, 42},
			"wood":  {7, 8, 15, 16, 23, 24, 37, 38, 45, 46},
			"water": {13, 14, 21, 22, 29, 30, 43, 44},
			"fire":  {1, 2, 9, 10, 17, 18, 31, 32, 39, 40, 47, 48},
			"earth": {5, 6, 19, 20, 27, 28, 35, 36, 49},
		},
		2026: {
			"metal": {4, 5, 12, 13, 26, 27, 34, 35, 42, 43},
			"wood":  {8, 9, 16, 17, 24, 25, 38, 39, 46, 47},
			"water": {1, 14, 15, 22, 23, 30, 31, 44, 45},
			"fire":  {2, 3, 10, 11, 18, 19, 32, 33, 40, 41, 48, 49},
			"earth": {6, 7, 20, 21, 28, 29, 36, 37},
		},
		2027: {
			"metal": {5, 6, 13, 14, 27, 28, 35, 36, 43, 44},
			"wood":  {9, 10, 17, 18, 25, 26, 39, 40, 47, 48},
			"water": {1, 2, 15, 16, 23, 24, 31, 32, 45, 46},
			"fire":  {3, 4, 11, 12, 19, 20, 33, 34, 41, 42, 49},
			"earth": {7, 8, 21, 22, 29, 30, 37, 38},
		},
		2028: {
			"metal": {6, 7, 14, 15, 28, 29, 36, 37, 44, 45},
			"wood":  {10, 11, 18, 19, 26, 27, 40, 41, 48, 49},
			"water": {2, 3, 16, 17, 24, 25, 32, 33, 46, 47},
			"fire":  {4, 5, 12, 13, 20, 21, 34, 35, 42, 43},
			"earth": {1, 8, 9, 22, 23, 30, 31, 38, 39},
		},
		2029: {
			"metal": {7, 8, 15, 16, 29, 30, 37, 38, 45, 46},
			"wood":  {11, 12, 19, 20, 27, 28, 41, 42, 49},
			"water": {3, 4, 17, 18, 25, 26, 33, 34, 47, 48},
			"fire":  {5, 6, 13, 14, 21, 22, 35, 36, 43, 44},
			"earth": {1, 2, 9, 10, 23, 24, 31, 32, 39, 40},
		},
		2030: {
			"metal": {1, 8, 9, 16, 17, 30, 31, 38, 39, 46, 47},
			"wood":  {12, 13, 20, 21, 28, 29, 42, 43},
			"water": {4, 5, 18, 19, 26, 27, 34, 35, 48, 49},
			"fire":  {6, 7, 14, 15, 22, 23, 36, 37, 44, 45},
			"earth": {2, 3, 10, 11, 24, 25, 32, 33, 40, 41},
		},
		2031: {
			"metal": {1, 2, 9, 10, 17, 18, 31, 32, 39, 40, 47, 48},
			"wood":  {13, 14, 21, 22, 29, 30, 43, 44},
			"water": {5, 6, 19, 20, 27, 28, 35, 36, 49},
			"fire":  {7, 8, 15, 16, 23, 24, 37, 38, 45, 46},
			"earth": {3, 4, 11, 12, 25, 26, 33, 34, 41, 42},
		},
		2032: {
			"metal": {2, 3, 10, 11, 18, 19, 32, 33, 40, 41, 48, 49},
			"wood":  {1, 14, 15, 22, 23, 30, 31, 44, 45},
			"water": {6, 7, 20, 21, 28, 29, 36, 37},
			"fire":  {8, 9, 16, 17, 24, 25, 38, 39, 46, 47},
			"earth": {4, 5, 12, 13, 26, 27, 34, 35, 42, 43},
		},
		2033: {
			"metal": {3, 4, 11, 12, 19, 20, 33, 34, 41, 42, 49},
			"wood":  {1, 2, 15, 16, 23, 24, 31, 32, 45, 46},
			"water": {7, 8, 21, 22, 29, 30, 37, 38},
			"fire":  {9, 10, 17, 18, 25, 26, 39, 40, 47, 48},
			"earth": {5, 6, 13, 14, 27, 28, 35, 36, 43, 44},
		},
		2034: {
			"metal": {4, 5, 12, 13, 20, 21, 34, 35, 42, 43},
			"wood":  {2, 3, 16, 17, 24, 25, 32, 33, 46, 47},
			"water": {1, 8, 9, 22, 23, 30, 31, 38, 39},
			"fire":  {10, 11, 18, 19, 26, 27, 40, 41, 48, 49},
			"earth": {6, 7, 14, 15, 28, 29, 36, 37, 44, 45},
		},
		2035: {
			"metal": {5, 6, 13, 14, 21, 22, 35, 36, 43, 44},
			"wood":  {3, 4, 17, 18, 25, 26, 33, 34, 47, 48},
			"water": {1, 2, 9, 10, 23, 24, 31, 32, 39, 40},
			"fire":  {11, 12, 19, 20, 27, 28, 41, 42, 49},
			"earth": {7, 8, 15, 16, 29, 30, 37, 38, 45, 46},
		},
		2036: {
			"metal": {6, 7, 14, 15, 22, 23, 36, 37, 44, 45},
			"wood":  {4, 5, 18, 19, 26, 27, 34, 35, 48, 49},
			"water": {2, 3, 10, 11, 24, 25, 32, 33, 40, 41},
			"fire":  {12, 13, 20, 21, 28, 29, 42, 43},
			"earth": {1, 8, 9, 16, 17, 30, 31, 38, 39, 46, 47},
		},
	}
	yearNumberWuxings = map[int]map[int]string{}
)

func init() {
	// init yearNumberWuxings
	for year, wuxingNumbers := range yearWuxingNumbers {
		numberWuxings := map[int]string{}
		for wuxing, numbers := range wuxingNumbers {
			for _, number := range numbers {
				numberWuxings[number] = wuxing
			}
		}

		yearNumberWuxings[year] = numberWuxings
	}
}

// GetWuxingNumbers 取得全部五行的號碼
func GetWuxingNumbers(year int) (map[string][]int, error) {
	wuxingNumbers, ok := yearWuxingNumbers[year]
	if !ok {
		return nil, fmt.Errorf("not supported for %d", year)
	}

	return wuxingNumbers, nil
}

// GetWuxingNumber 取得單一五行的號碼
func GetWuxingNumber(year int, wuxing string) ([]int, error) {
	wuxingNumbers, err := GetWuxingNumbers(year)
	if err != nil {
		return nil, err
	}

	numbers, ok := wuxingNumbers[wuxing]
	if !ok {
		return nil, fmt.Errorf("invalid wuxing %s", wuxing)
	}

	return numbers, nil
}

// GetNumberWuxings 取得全部號碼的五行
func GetNumberWuxings(year int) (map[int]string, error) {
	numberWuxings, ok := yearNumberWuxings[year]
	if !ok {
		return nil, fmt.Errorf("not supported for %d", year)
	}

	return numberWuxings, nil
}

// GetNumberWuxing 取得單一號碼的五行
func GetNumberWuxing(year, number int) (string, error) {
	numberWuxings, err := GetNumberWuxings(year)
	if err != nil {
		return "", err
	}

	wuxing, ok := numberWuxings[number]
	if !ok {
		return "", fmt.Errorf("invalid number %d", number)
	}

	return wuxing, nil
}
