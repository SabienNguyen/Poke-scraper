## Poke-scraper
Poke-scraper is a tool that allows users to collect data on pokemon from the pokemondb database! 

## installation
```
go 1.20
require (
  github.com/gocolly/colly/
)
```
## run
```
go build main.go
go run main.go
```

## using flags

```
go run main.go -type1=xxxx -type2=xxxx //default is ""
```
