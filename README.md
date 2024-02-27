## Web Scraper
Webscraper to collect pokemon information off the pokemondb.net pokedex
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
