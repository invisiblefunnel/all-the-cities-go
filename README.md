# all-the-cities

Golang port of [zeke/all-the-cities](https://github.com/zeke/all-the-cities).

```go
var (
    cities []allthecities.City
    err    error
)

cities, err = allthecities.Load()
```
