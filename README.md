# Search Vietnamese

![Github action](https://github.com/Nguyen-Hoang-Nam/search-vietnamese/actions/workflows/go.yml/badge.svg)

Simple go package to search vietnamese words

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
BenchmarkContain-8       1561740               755.8 ns/op
```

## TODO

- [x] Support index function
- [x] Able to match "ễ" and "ê"
- [x] Add benchmark
- [x] Support fuzzy search
- [x] Add case-sensitive mode
- [x] Add CI/CD
- [ ] Add coverage badge

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
