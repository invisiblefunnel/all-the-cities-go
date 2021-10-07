# all-the-cities [![Tests](https://github.com/invisiblefunnel/all-the-cities-go/actions/workflows/go.yml/badge.svg)](https://github.com/invisiblefunnel/all-the-cities-go/actions/workflows/go.yml)

Golang port of [zeke/all-the-cities](https://github.com/zeke/all-the-cities).

```go
var (
    cities []allthecities.City
    err    error
)

cities, err = allthecities.Load()
```
