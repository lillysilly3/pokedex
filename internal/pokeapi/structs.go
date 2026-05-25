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
	 PokemonEncounters	[]struct {
		Pokemon 		struct {
			Name 		string		`json:"name"`
		} 							`json:"pokemon"`
	 }								`json:"pokemon_encounters"`
}

type Pokemon struct {
	Id				int		`json:"id"`
	Name			string	`json:"name"`
	BaseExperience	int		`json:"base_experience"`
	Height			int		`json:"height"`
	Weight			int		`json:"weight"`
	Stats			[]struct {
		BaseStat	int		`json:"base_stat"`
		Effort		int		`json:"effort"`
		Stat		struct {
			Name	string	`json:"name"`
			URL		string	`json:"url"`
		}					`json:"stat"`
	}						`json:"stats"`
	Types			[]struct {
		Slot		int		`json:"slot"`
		Type		struct {
			Name	string	`json:"name"`
			URL		string	`json:"url"`
		}					`json:"type"`
	}						`json:"types"`
}

