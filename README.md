# zodiacwuxing-go

Mark Six each year zodiac and wuxing numbers

# Installation

```sh
go get github.com/jetfueltw/zodiacwuxing-go
```

```go
import "github.com/jetfueltw/zodiacwuxing-go"
```

# Useage

取得 2020 年的生肖號碼
```go
zodiacNumbers, err := zodiacwuxing.GetZodiacNumbers(2020)
// [
//     "rat"     => [1, 13, 25, 37, 49],
//     "ox"      => [12, 24, 36, 48],
//     "tiger"   => [11, 23, 35, 47],
//     "rabbit"  => [10, 22, 34, 46],
//     "dragon"  => [9, 21, 33, 45],
//     ...
// ]
```

取得 2020 年的號碼生肖
```go
numberZodiacs, err := zodiacwuxing.GetNumberZodiacs(2020)
// [
//     1 => "rat",
//     2 => "pig",
//     3 => "dog",
//     4 => "rooster",
//     5 => "monkey",
//     ...
// ]
```

取得 2020 年的生肖
```go
dutyZodiac, err := zodiacwuxing.GetDutyZodiac(year)
// "rat"
```

取得 2020 年的五行號碼
```go
wuxingNumbers, err := zodiacwuxing.GetWuxingNumbers(2020)
// [
//     "metal" => [6, 7, 20, 21, 28, 29, 36, 37],
//     "wood"  => [2, 3, 10, 11, 18, 19, 32, 33, 40, 41, 48, 49],
//     "water" => [8, 9, 16, 17, 24, 25, 38, 39, 46, 47],
//     "fire"  => [4, 5, 12, 13, 26, 27, 34, 35, 42, 43],
//     "earth" => [1, 14, 15, 22, 23, 30, 31, 44, 45],
// ]
```

取得 2020 年的號碼五行
```go
numberWuxings, err := zodiacwuxing.GetNumberWuxings(2020)
// [
//     1 => "earth",
//     2 => "wood",
//     3 => "wood",
//     4 => "fire",
//     5 => "fire",
//     ...
// ]
```
