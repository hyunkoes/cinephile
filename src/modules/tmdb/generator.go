package tmdb

func TmdbPosterAPI(path string) string {
	return `https://image.tmdb.org/t/p/original` + path
}
