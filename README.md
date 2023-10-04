# zodiacwuxing-go

Mark Six each year zodiac and wuxing numbers

# Installation

```sh
go get github.com/jetfueltw/zodiacwuxing-go
```

```go
import "github.com/jetfueltw/zodiacwuxing-go/lhc6"
```

# Useage

取得 2020 年全部生肖的號碼
```go
zodiacNumbers, err := lhc6.GetZodiacNumbers(2020)
// [
//     "rat"     => [1, 13, 25, 37, 49],
//     "ox"      => [12, 24, 36, 48],
//     "tiger"   => [11, 23, 35, 47],
//     "rabbit"  => [10, 22, 34, 46],
//     "dragon"  => [9, 21, 33, 45],
//     ...
// ]
```

取得 2020 年單一生肖的號碼
```go
zodiacNumber, err := lhc6.GetZodiacNumber(2020, "rat")
// [1, 13, 25, 37, 49]
```

取得 2020 年全部號碼的生肖
```go
numberZodiacs, err := lhc6.GetNumberZodiacs(2020)
// [
//     1 => "rat",
//     2 => "pig",
//     3 => "dog",
//     4 => "rooster",
//     5 => "monkey",
//     ...
// ]
```

取得 2020 年單一號碼的生肖
```go
numberZodiac, err := lhc6.GetNumberZodiac(2020, 1)
// "rat"
```

取得 2020 年的值年生肖
```go
dutyZodiac, err := lhc6.GetDutyZodiac(2020)
// "rat"
```

取得 2020 年全部五行的號碼
```go
wuxingNumbers, err := lhc6.GetWuxingNumbers(2020)
// [
//     "metal" => [6, 7, 20, 21, 28, 29, 36, 37],
//     "wood"  => [2, 3, 10, 11, 18, 19, 32, 33, 40, 41, 48, 49],
//     "water" => [8, 9, 16, 17, 24, 25, 38, 39, 46, 47],
//     "fire"  => [4, 5, 12, 13, 26, 27, 34, 35, 42, 43],
//     "earth" => [1, 14, 15, 22, 23, 30, 31, 44, 45],
// ]
```

取得 2020 年單一五行的號碼
```go
wuxingNumber, err := lhc6.GetWuxingNumber(2020, "metal")
// [6, 7, 20, 21, 28, 29, 36, 37]
```

取得 2020 年全部號碼的五行
```go
numberWuxings, err := lhc6.GetNumberWuxings(2020)
// [
//     1 => "earth",
//     2 => "wood",
//     3 => "wood",
//     4 => "fire",
//     5 => "fire",
//     ...
// ]
```

取得 2020 年單一號碼的五行
```go
numberWuxing, err := lhc6.GetNumberWuxing(2020, 1)
// "earth"
```
