import sys
import time
import requests
import os
from dotenv import load_dotenv

def get_token():
    load_dotenv()
    API_KEY = os.getenv('TMDB_TOKEN')
    url = "https://api.themoviedb.org/3/authentication"

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(API_KEY)
    }
    response = requests.get(url, headers=headers).json()
    if (response['status_code'] != 1):
        print("[ERR] Wrong access token ! ")
        sys.exit(0)
    return API_KEY
def get_movie_list(token, page):
    url = "https://api.themoviedb.org/3/movie/popular?language={}&page={}".format("ko-KR",page)

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(token)
    }
    response = requests.get(url, headers=headers).json()
    return  response['results']

def movies_to_sql(movies):
    movie_sql = 'INSERT INTO movie VALUES ({},{},"{}","{}","{}","{}","{}");\n'
    genre_sql = 'INSERT INTO genre VALUES ({}, {});\n'
    movie_sql_total = ""
    genre_sql_total = ""
    for movie in movies :
        movie['overview'] = movie['overview'].replace('"', "'")
        movie_sql_total += movie_sql.format(movie['id'],movie['adult'],movie['original_title'],movie['title'] or '',movie['poster_path'],movie.get('release_date', ''),movie['overview'])
        for genre in movie['genre_ids']:
            genre_sql_total += genre_sql.format(movie['id'],genre)
    return movie_sql_total + genre_sql_total



def run_script():
    token = get_token()
    movies = []
    for i in range(1, 100):
        next_movies = get_movie_list(token,i)
        if ( len(next_movies) == 0 ):
            break
        movies = movies+next_movies
        time.sleep(1)
    sql = movies_to_sql(movies)
    with open('../mysql/initdb.d/popular_movie.sql','w+') as f:
        f.write(sql)
    

if __name__ == "__main__":
    run_script()
