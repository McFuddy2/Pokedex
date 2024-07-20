package main

import (
	"github.com/McFuddy2/Pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient 				pokeapi.Client
    NextLocationAreaURL 	 	*string 
	PreviousLocationAreaURL  	*string
	caughtPokemon				map[string]pokeapi.Pokemon
}

func main() {

	cnfg := config{
		pokeapiClient: 	pokeapi.NewClient(time.Minute*90),
		caughtPokemon:	make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cnfg)
}

/*
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

var cache *pokecache.Cache

func exploreCommand(areaName string) {
	fmt.Printf("Exploring %s...\n", areaName)

	for _, loc := range cache.Locations{
		if loc.Name == areaName {
			url = loc.URL
			break
		}
	}
	if url == "" {
		log.Printf("Location %s not found.", areaName)
		return
	}
	

	respData, err := fetchLocationData(url)
	if err != nil {
		log.Printf("Error fetching data from PokeAPI: %v", err)
		return
	}
	defer respData.Body.Close()

	var locationArea LocationAreaResponse
	err = json.NewDecoder(respData.Body).Decode(&locationArea)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	var pokemonNames []string
	for _, encounter := range locationArea.Results {
		pokemonNames = append(pokemonNames, encounter.Name)
	}

	encodedPokemonNames, err := json.Marshal(pokemonNames)
	if err != nil {
		log.Printf("Error encoding pokemonNames to JSON: %v", err)
		return
	}

	cache.Set(areaName, encodedPokemonNames)

	displayPokemonNames(pokemonNames)
}

func displayPokemonNames(names []string) {
	fmt.Println("Found Pokemon:")
	english := language.English
	tc := cases.Title(english)
	for _, name := range names {
		title := tc.String(name)
		fmt.Printf(" - %s\n", title)
	}
}

func fetchLocations(config *Config, direction string) (map[string]string, *Config, error) {
	var url string

	//are we moving to the next set of 20 or the previous set of 20.... or just the first page?
	if direction == "next" {
		if config.Next != "" {
			url = config.Next
		} else {
			url = config.BaseURL
		}
	} else if direction == "previous" {
		if config.Previous != "" {
			url = config.Previous
		} else {
			fmt.Println("Silly Swanna, we are already at the beginning of the list. I'll just give you the first page again.")
			url = config.BaseURL
		}
	} else {
		url = config.BaseURL
	}

	//we are checking to see if that info is already in the cache
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(resp)

	defer resp.Body.Close()

	// data is just a ton of numbers
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	// locations is now a list of locations
	locations, newConfig, err := parseLocationData(data)
	if err != nil {
		return nil, nil, err
	}


	return locations, newConfig, nil
}

	func parseLocationData(data []byte) (map[string]string, *Config, error) {
		var result struct {
			Results []struct { 
				Name string `json:"name"`
				URL string `json:"url"`
				} `json:"results"`
			Next     string 								`json:"next"`
			Previous string 								`json:"previous"`
		}

		// this turns all the random numbers from data into a list of locations and URLs
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, nil, err
		}

		// here we make sure locations is a map that is Name:URL
		locations := make(map[string]string)
		for _, loc := range result.Results {
			locations[loc.Name] = loc.URL
		}


		config := &Config{
			Next:     result.Next,
			Previous: result.Previous,
		}
	
		return locations, config, nil
	}
	
func displayLocations(locations map[string]string) {
	for key := range locations {
		fmt.Println(key)
	}
}
	*/