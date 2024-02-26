package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {

	pageToScrape := "https://pokemondb.net/pokedex/all"
	pokemonData := [][]string{{"id", "Name", "Type", "Total", "HP", "Attack", "Defense", "Sp.Atk", "Sp.Def", "Speed"}}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
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
			pokemonData = append(pokemonData, pokemonInfo)
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
