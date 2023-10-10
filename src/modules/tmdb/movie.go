package tmdb

import (
	. "cinephile/modules/dto"
	"cinephile/modules/logging"
	"encoding/json"
	"fmt"
	"io"
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
		var url string
		if strings.ToLower(site) == "youtube" {
			url = "https://www.youtube.com/watch?v=" + key
		}
		trailer.Site = site
		trailer.Key = key
		trailer.Official = official
		trailer.Url = url
		trailers = append(trailers, trailer)
	}
	return trailers
}

func GetStillCut(movie_id int) []string {
	tmdb_token = os.Getenv("TMDB_TOKEN")
	url := fmt.Sprintf(`https://api.themoviedb.org/3/movie/%d/images`, movie_id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI5NjJlYWFhN2NhNDk4ZjA4ZGM1OWQzNTg5N2VjZWIwOSIsInN1YiI6IjY0YWZiYzg3NmEzNDQ4MDBlYThmNTczOCIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.96gO_hH6s2BKQt_UDj3Bgp-yuEfufwxaKBf9_TrjQwM")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var payload interface{}
	_ = json.Unmarshal(body, &payload)
	jsonData := payload.(map[string]interface{})
	results, _ := jsonData["backdrops"].([]interface{})
	stillcuts := make([]string, 0)
	for _, item := range results {
		result := item.(map[string]interface{})
		stillcut := TmdbPosterAPI(result[`file_path`].(string))
		stillcuts = append(stillcuts, stillcut)
	}
	return stillcuts
}
