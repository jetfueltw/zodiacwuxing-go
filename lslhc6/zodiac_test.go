package lslhc6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetZodiacNumbers(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020", func(t *testing.T) {
		zodiacNumbers, err := GetZodiacNumbers(2020)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[string][]int{
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
		}, zodiacNumbers)
	})

	t.Run("2036", func(t *testing.T) {
		zodiacNumbers, err := GetZodiacNumbers(2036)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[string][]int{
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
		}, zodiacNumbers)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetZodiacNumbers(2019)
		assert.EqualError(err, "not supported for 2019")
	})
}

func TestGetZodiacNumber(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020 rat", func(t *testing.T) {
		zodiacNumber, err := GetZodiacNumber(2020, "rat")
		if !assert.NoError(err) {
			return
		}

		assert.Equal([]int{1, 13, 25, 37, 49}, zodiacNumber)
	})

	t.Run("2020 pig", func(t *testing.T) {
		zodiacNumber, err := GetZodiacNumber(2020, "pig")
		if !assert.NoError(err) {
			return
		}

		assert.Equal([]int{2, 14, 26, 38, 50}, zodiacNumber)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetZodiacNumber(2019, "rat")
		assert.EqualError(err, "not supported for 2019")
	})

	t.Run("error zodiac", func(t *testing.T) {
		_, err := GetZodiacNumber(2020, "dummy")
		assert.EqualError(err, "invalid zodiac dummy")
	})
}

func TestGetNumberZodiacs(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020", func(t *testing.T) {
		numberZodiacs, err := GetNumberZodiacs(2020)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[int]string{
			1:  "rat",
			2:  "pig",
			3:  "dog",
			4:  "rooster",
			5:  "monkey",
			6:  "goat",
			7:  "horse",
			8:  "snake",
			9:  "dragon",
			10: "rabbit",
			11: "tiger",
			12: "ox",
			13: "rat",
			14: "pig",
			15: "dog",
			16: "rooster",
			17: "monkey",
			18: "goat",
			19: "horse",
			20: "snake",
			21: "dragon",
			22: "rabbit",
			23: "tiger",
			24: "ox",
			25: "rat",
			26: "pig",
			27: "dog",
			28: "rooster",
			29: "monkey",
			30: "goat",
			31: "horse",
			32: "snake",
			33: "dragon",
			34: "rabbit",
			35: "tiger",
			36: "ox",
			37: "rat",
			38: "pig",
			39: "dog",
			40: "rooster",
			41: "monkey",
			42: "goat",
			43: "horse",
			44: "snake",
			45: "dragon",
			46: "rabbit",
			47: "tiger",
			48: "ox",
			49: "rat",
			50: "pig",
			51: "dog",
			52: "rooster",
			53: "monkey",
			54: "goat",
			55: "horse",
			56: "snake",
			57: "dragon",
			58: "rabbit",
			59: "tiger",
			60: "ox",
		}, numberZodiacs)
	})

	t.Run("2036", func(t *testing.T) {
		numberZodiacs, err := GetNumberZodiacs(2036)
		if !assert.NoError(err) {
			return
		}

		assert.Equal(map[int]string{
			1:  "dragon",
			2:  "rabbit",
			3:  "tiger",
			4:  "ox",
			5:  "rat",
			6:  "pig",
			7:  "dog",
			8:  "rooster",
			9:  "monkey",
			10: "goat",
			11: "horse",
			12: "snake",
			13: "dragon",
			14: "rabbit",
			15: "tiger",
			16: "ox",
			17: "rat",
			18: "pig",
			19: "dog",
			20: "rooster",
			21: "monkey",
			22: "goat",
			23: "horse",
			24: "snake",
			25: "dragon",
			26: "rabbit",
			27: "tiger",
			28: "ox",
			29: "rat",
			30: "pig",
			31: "dog",
			32: "rooster",
			33: "monkey",
			34: "goat",
			35: "horse",
			36: "snake",
			37: "dragon",
			38: "rabbit",
			39: "tiger",
			40: "ox",
			41: "rat",
			42: "pig",
			43: "dog",
			44: "rooster",
			45: "monkey",
			46: "goat",
			47: "horse",
			48: "snake",
			49: "dragon",
			50: "rabbit",
			51: "tiger",
			52: "ox",
			53: "rat",
			54: "pig",
			55: "dog",
			56: "rooster",
			57: "monkey",
			58: "goat",
			59: "horse",
			60: "snake",
		}, numberZodiacs)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetNumberZodiacs(2019)
		assert.EqualError(err, "not supported for 2019")
	})
}

func TestGetNumberZodiac(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020 1", func(t *testing.T) {
		numberZodiac, err := GetNumberZodiac(2020, 1)
		if !assert.NoError(err) {
			return
		}

		assert.Equal("rat", numberZodiac)
	})

	t.Run("2020 60", func(t *testing.T) {
		numberZodiac, err := GetNumberZodiac(2020, 60)
		if !assert.NoError(err) {
			return
		}

		assert.Equal("ox", numberZodiac)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetNumberZodiac(2037, 1)
		assert.EqualError(err, "not supported for 2037")
	})

	t.Run("error number", func(t *testing.T) {
		_, err := GetNumberZodiac(2020, 61)
		assert.EqualError(err, "invalid number 61")
	})
}

func TestGetDutyZodiac(t *testing.T) {
	assert := assert.New(t)

	t.Run("2020", func(t *testing.T) {
		dutyZodiac, err := GetDutyZodiac(2020)
		if !assert.NoError(err) {
			return
		}

		assert.Equal("rat", dutyZodiac)
	})

	t.Run("2036", func(t *testing.T) {
		dutyZodiac, err := GetDutyZodiac(2036)
		if !assert.NoError(err) {
			return
		}

		assert.Equal("dragon", dutyZodiac)
	})

	t.Run("error year", func(t *testing.T) {
		_, err := GetDutyZodiac(2000)
		assert.EqualError(err, "not supported for 2000")
	})
}
