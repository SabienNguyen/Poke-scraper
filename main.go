package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	var pokemonType1 string
	var pokemonType2 string
	flag.StringVar(&pokemonType1, "type1", "", "Type of Pokemon to scrape")
	flag.StringVar(&pokemonType2, "type2", "", "Type of Pokemon to scrape")
	flag.Parse()
	pageToScrape := "https://pokemondb.net/pokedex/all/"

	pokemonData := [][]string{{"id", "Name", "Type", "Total", "HP", "Attack", "Defense", "Sp.Atk", "Sp.Def", "Speed"}}

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Printf("%v\n", e.Attr("href"))
	})

	c.OnHTML(".data-table", func(e *colly.HTMLElement) {
		// For each row in the table, extract Pokemon data
		e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			var pokemonInfo []string
			el.ForEach("td", func(index int, elem *colly.HTMLElement) {
				pokemonInfo = append(pokemonInfo, elem.Text)
			})
			if strings.Contains(strings.ToLower(pokemonInfo[2]), strings.ToLower(pokemonType1)) && strings.Contains(strings.ToLower(pokemonInfo[2]), pokemonType2) {
				pokemonData = append(pokemonData, pokemonInfo)
			}
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.Visit(pageToScrape)

	file, err := os.Create("Pokemon.csv")
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range pokemonData {
		err := writer.Write(value)
		if err != nil {
			log.Fatalf("Error writing record to file: %s", err)
		}
	}

	fmt.Println("CSV file saved successfully.")

}
