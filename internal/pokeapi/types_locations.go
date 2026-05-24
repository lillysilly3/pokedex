package pokeapi

type Result struct {
	Name	string		`json:"name"`
	Url		string		`json:"url"`
}

type RespShallowLocations struct {
	Count		int			`json:"count"`
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results		[]Result	`json:"results"`
}

type PokemonInArea struct {
	 PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name`
		} `json:"pokemon"`
	 } `json:"pokemon_encounters"`
}

