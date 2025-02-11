package lsc6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWuxingNumbers(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020", func(t *testing.T) {
		wuxingNumbers, err := GetWuxingNumbers(2020)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[string][]int{
			"metal": {6, 7, 20, 21, 28, 29, 36, 37, 50, 51, 58, 59},
			"wood":  {2, 3, 10, 11, 18, 19, 32, 33, 40, 41, 48, 49},
			"water": {8, 9, 16, 17, 24, 25, 38, 39, 46, 47, 54, 55},
			"fire":  {4, 5, 12, 13, 26, 27, 34, 35, 42, 43, 56, 57},
			"earth": {1, 14, 15, 22, 23, 30, 31, 44, 45, 52, 53, 60},
		}, wuxingNumbers)
	})

	t.Run("2036", func(t *testing.T) {
		wuxingNumbers, err := GetWuxingNumbers(2036)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[string][]int{
			"metal": {6, 7, 14, 15, 22, 23, 36, 37, 44, 45, 52, 53},
			"wood":  {4, 5, 18, 19, 26, 27, 34, 35, 48, 49, 56, 57},
			"water": {2, 3, 10, 11, 24, 25, 32, 33, 40, 41, 54, 55},
			"fire":  {12, 13, 20, 21, 28, 29, 42, 43, 50, 51, 58, 59},
			"earth": {1, 8, 9, 16, 17, 30, 31, 38, 39, 46, 47, 60},
		}, wuxingNumbers)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetWuxingNumbers(2019)
		assert.EqualError(err, "not supported for 2019")
	})
}

func TestGetWuxingNumber(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020 water", func(t *testing.T) {
		wuxingNumber, err := GetWuxingNumber(2020, "water")
		if !assert.NoError(err) {
			return
		}

		assert.Equal([]int{8, 9, 16, 17, 24, 25, 38, 39, 46, 47, 54, 55}, wuxingNumber)
	})

	t.Run("2020 earth", func(t *testing.T) {
		wuxingNumber, err := GetWuxingNumber(2020, "earth")
		if !assert.NoError(err) {
			return
		}

		assert.Equal([]int{1, 14, 15, 22, 23, 30, 31, 44, 45, 52, 53, 60}, wuxingNumber)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetWuxingNumber(2019, "fire")
		assert.EqualError(err, "not supported for 2019")
	})

	t.Run("error zodiac", func(t *testing.T) {
		_, err := GetWuxingNumber(2020, "dummy")
		assert.EqualError(err, "invalid wuxing dummy")
	})
}

func TestGetNumberWuxings(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020", func(t *testing.T) {
		numberWuxings, err := GetNumberWuxings(2020)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[int]string{
			1:  "earth",
			2:  "wood",
			3:  "wood",
			4:  "fire",
			5:  "fire",
			6:  "metal",
			7:  "metal",
			8:  "water",
			9:  "water",
			10: "wood",
			11: "wood",
			12: "fire",
			13: "fire",
			14: "earth",
			15: "earth",
			16: "water",
			17: "water",
			18: "wood",
			19: "wood",
			20: "metal",
			21: "metal",
			22: "earth",
			23: "earth",
			24: "water",
			25: "water",
			26: "fire",
			27: "fire",
			28: "metal",
			29: "metal",
			30: "earth",
			31: "earth",
			32: "wood",
			33: "wood",
			34: "fire",
			35: "fire",
			36: "metal",
			37: "metal",
			38: "water",
			39: "water",
			40: "wood",
			41: "wood",
			42: "fire",
			43: "fire",
			44: "earth",
			45: "earth",
			46: "water",
			47: "water",
			48: "wood",
			49: "wood",
			50: "metal",
			51: "metal",
			52: "earth",
			53: "earth",
			54: "water",
			55: "water",
			56: "fire",
			57: "fire",
			58: "metal",
			59: "metal",
			60: "earth",
		}, numberWuxings)
	})

	t.Run("2036", func(t *testing.T) {
		numberWuxings, err := GetNumberWuxings(2036)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[int]string{
			1:  "earth",
			2:  "water",
			3:  "water",
			4:  "wood",
			5:  "wood",
			6:  "metal",
			7:  "metal",
			8:  "earth",
			9:  "earth",
			10: "water",
			11: "water",
			12: "fire",
			13: "fire",
			14: "metal",
			15: "metal",
			16: "earth",
			17: "earth",
			18: "wood",
			19: "wood",
			20: "fire",
			21: "fire",
			22: "metal",
			23: "metal",
			24: "water",
			25: "water",
			26: "wood",
			27: "wood",
			28: "fire",
			29: "fire",
			30: "earth",
			31: "earth",
			32: "water",
			33: "water",
			34: "wood",
			35: "wood",
			36: "metal",
			37: "metal",
			38: "earth",
			39: "earth",
			40: "water",
			41: "water",
			42: "fire",
			43: "fire",
			44: "metal",
			45: "metal",
			46: "earth",
			47: "earth",
			48: "wood",
			49: "wood",
			50: "fire",
			51: "fire",
			52: "metal",
			53: "metal",
			54: "water",
			55: "water",
			56: "wood",
			57: "wood",
			58: "fire",
			59: "fire",
			60: "earth",
		}, numberWuxings)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetNumberWuxings(2019)
		assert.EqualError(err, "not supported for 2019")
	})
}

func TestGetNumberWuxing(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020 1", func(t *testing.T) {
		numberWuxing, err := GetNumberWuxing(2020, 1)
		if !assert.NoError(err) {
			return
		}

		assert.Equal("earth", numberWuxing)
	})

	t.Run("2020 60", func(t *testing.T) {
		numberWuxing, err := GetNumberWuxing(2020, 60)
		if !assert.NoError(err) {
			return
		}

		assert.Equal("earth", numberWuxing)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetNumberWuxing(2037, 1)
		assert.EqualError(err, "not supported for 2037")
	})

	t.Run("error number", func(t *testing.T) {
		_, err := GetNumberWuxing(2020, 0)
		assert.EqualError(err, "invalid number 0")
	})
}
