package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationData(name *string) (LocationAreaData, error) {
	url := BaseURL + "/location-area/" + *name

	cachedData, exists := c.cache.Get(url)

	if exists {
		data := LocationAreaData{}
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return LocationAreaData{}, err
		}

		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationAreaData{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaData{}, err
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreaData{}, err
	}

	data := LocationAreaData{}
	if err := json.Unmarshal(resBody, &data); err != nil {
		return LocationAreaData{}, err
	}

	c.cache.Add(url, resBody)

	return data, nil
}
