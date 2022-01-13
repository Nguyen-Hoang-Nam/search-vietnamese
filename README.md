# Search Vietnamese

Simple go package to search vietnamese words

## Installation

```go
go get github.com/Nguyen-Hoang-Nam/search-vietnamese
```

## Usage

```go
package main

import (
    searchvietnamese "github.com/Nguyen-Hoang-Nam/search-vietnamese"
)

func main() {
    checkContain := searchvietnamese.Contain("Nguyễn Hoàng Nam", "nguyen")
}
```

## TODO

- [ ] Support index function
- [ ] Able to match "ễ" and "ê"
- [ ] SIMP
- [ ] Add benchmark

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
