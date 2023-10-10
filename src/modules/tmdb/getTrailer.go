package tmdb

import (
	. "cinephile/modules/dto"
	"cinephile/modules/logging"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var tmdb_token string

func GetTrailer(movie_id int) []Trailer {
	tmdb_token = os.Getenv("TMDB_TOKEN")
	trailers := make([]Trailer, 0)
	var trailer Trailer
	url := fmt.Sprintf(`https://api.themoviedb.org/3/movie/%d/videos?language=en-US`, movie_id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+tmdb_token)

	resp, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Warn(err.Error())
		return []Trailer{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logging.Warn(err.Error())
		return []Trailer{}
	}
	var payload interface{}
	_ = json.Unmarshal(body, &payload)
	jsonData := payload.(map[string]interface{})
	results, ok := jsonData["results"].([]interface{})
	if !ok {
		panic("Failed to extract results")
	}
	for _, item := range results {
		result := item.(map[string]interface{})
		site := result["site"].(string)
		key := result["key"].(string)
		official := result["official"].(bool)
		fmt.Println(site, key)
		var url string
		if strings.ToLower(site) == "youtube" {
			fmt.Println("Youtube임 ㅋ")
			url = "https://www.youtube.com/watch?v=" + key
			fmt.Println(site, key, url)
		}
		trailer.Site = site
		trailer.Key = key
		trailer.Official = official
		trailer.Url = url
		trailers = append(trailers, trailer)
	}
	return trailers
}
