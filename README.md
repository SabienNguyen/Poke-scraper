## Poke-scraper
Poke-scraper is a tool that allows users to collect data on pokemon from the pokemondb database! Can be useful for anyone searching for an easy way to view stats of pokemon filtered by pokemon type. Can be great for putting together your pokemon team or collecting data for analytics.
## What it does
Scans the pokemondb database webpage to collect the pokemon id, name and stats and collects it all into a csv format. This tool also allows for filtering based off of pokemon type. 
## How it works
The program starts with checking for user defined flags and then creates an empty 2d array, pokemonData, that will hold our scraped data. Using colly, the scraper is able to identify the database table, rows and columns. This information gets added to a pokemon info array that is appended to the pokemonData array of arrays. The pokemonData gets saved into a csv file called pokemon.csv which the users can download and use or view for later. 
## Example
```
go run main.go -type1=fire -type2=flying
```
<img width="776" alt="fire_flying" src="https://github.com/SabienNguyen/Poke-scraper/assets/32147674/26da4083-4ebc-413e-a002-a411dba2173d">

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
