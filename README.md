#GrizzlyAPI
## A W.I.P RestfulAPI written in Golang with the Echo framework for Grand Oaks Highschool
---

### To run
```go
go mod tidy && go run main.go
```
### Endpoints:

#### Orange/Blue Calendar

```
http://localhost:8000/date/:month/:day
```
Returns a JSON object representing the 'type' of day of the given date `blue`, `orange`, or `no school`

Example usage:
```ts
fetch("http://localhost:8000/date/12/4")
```
Returns
```ts
{
"date": "12/4",
"type": "Blue"
}
```
