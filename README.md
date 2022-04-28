# Search Vietnamese

![Github action](https://github.com/Nguyen-Hoang-Nam/search-vietnamese/actions/workflows/go.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/Nguyen-Hoang-Nam/search-vietnamese/badge.svg?branch=main)](https://coveralls.io/github/Nguyen-Hoang-Nam/search-vietnamese?branch=main)

Advance search tool for Vietnamese language. Able to match e and éèẻẽẹêếềểễệÉÈẺẼẸÊẾỀỂỄỆ

## Installation

```text
go get github.com/Nguyen-Hoang-Nam/search-vietnamese
```

## Usage

```go
package main

import (
    searchvietnamese "github.com/Nguyen-Hoang-Nam/search-vietnamese"
)

func main() {
    searchvietnamese.Contains("Nguyễn Hoàng Nam", "nguyen") // true
    searchvietnamese.Contains("Nguyễn Hoàng Nam", "nguyên") // true
    searchvietnamese.Contains("Nguyễn Hoàng Nam", "ngyên") // false

    index := searchvietnamese.Index("Nguyễn Hoàng Nam", "hoang") // 7
    index := searchvietnamese.Index("Nguyễn Hoàng Nam", "hang") // -1
}
```

## Benchmark

```text
goos: linux
goarch: amd64
pkg: github.com/Nguyen-Hoang-Nam/search-vietnamese
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkContain-8      	2194798	      530.8 ns/op
BenchmarkToAlphabet-8   	4899396	      239.9 ns/op
```

## Case sensitive

Default behaviour of all functions is match with case-insensitive.
To specific case-sensitive you may use function with postfix `Sensitive`

```go
searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "nguyen") // false
searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "Nguyên") // true
searchvietnamese.ContainsSensitive("Nguyễn Hoàng Nam", "ngyên") // false

index := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "hoang") // -1
index := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "Hoang") // 7
index := searchvietnamese.IndexSensitive("Nguyễn Hoàng Nam", "hang") // -1
```

## Strict mode

If you want more performance, this may help you by assuming that your text is
valid Vietnamese text. This means there are only alphabet and Vietnamese character.

This may cause true negative if your text has any kind of UTF-8 other than I
mention above.

```go
searchvietnamese.StrictContains("Nguyễn Hoàng Nam", "nguyen") // true
searchvietnamese.StrictContains("Nguyễn Hoàng Nam", "nguyên") // true
searchvietnamese.StrictContains("Nguyễn Hoàng Nam", "ngyên") // false

index := searchvietnamese.StrictIndex("Nguyễn Hoàng Nam", "hoang") // 7
index := searchvietnamese.StrictIndex("Nguyễn Hoàng Nam", "hang") // -1
```

### Compare normal mode and strict mode

```text
goos: linux
goarch: amd64
pkg: github.com/Nguyen-Hoang-Nam/search-vietnamese
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkContain-8            	2276416	      519.3 ns/op
BenchmarkToAlphabet-8         	4838238	      246.9 ns/op
BenchmarkStrictContain-8      	2785992	      426.2 ns/op
BenchmarkStrictToAlphabet-8   	7058474	      167.2 ns/op
```

## TODO

- [ ] Convert text to regex to search Vietnamese

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
