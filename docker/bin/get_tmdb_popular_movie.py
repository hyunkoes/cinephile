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
def genre_define(token):
    url = "https://api.themoviedb.org/3/genre/movie/list?language={}".format("ko-kr")
    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(token)
    }
    response = requests.get(url, headers=headers).json()
    return  response['genres']
    
def movies_to_sql(movies):
    movies.sort(key=lambda x: x.get('id'))
    movie_sql = 'INSERT INTO movie VALUES ({},{},"{}","{}","{}","{}","{}");\n'
    genre_sql = 'INSERT INTO genre VALUES ({}, {});\n'
    channel_sql = 'INSERT INTO channel (movie_id) VALUES ({});\n'
    movie_sql_total = ""
    genre_sql_total = ""
    channel_sql_total = ""
    for movie in movies :
        date = movie.get('release_date', '').replace('"','')
        movie_sql_total += movie_sql.format(movie['id'],movie['adult'],movie['original_title'].replace('"', "'"),movie['title'].replace('"', "'") or '',movie['poster_path'],date,movie['overview'].replace('"', "'")).replace('""',"null")
        channel_sql_total += channel_sql.format(movie['id'])
        for genre in movie['genre_ids']:
            genre_sql_total += genre_sql.format(movie['id'],genre)
    return movie_sql_total + genre_sql_total + channel_sql_total
def genre_to_sql(genres):
    genre_sql = 'INSERT INTO genre VALUES ({}, {});\n'
    genre_sql_total = ""
    for genre in genres :
        print(genre)
        # print(genre['id'], genre['name'])
        genre_sql_total += genre_sql.format(genre['id'],genre['name'])
    return genre_sql_total


def run_script():
    token = get_token()
    movies = []
    for i in range(1, 100):
        next_movies = get_movie_list(token,i)
        if ( len(next_movies) == 0 ):
            break
        movies = movies+next_movies
        time.sleep(1)
    sql = movies_to_sql(movies) + genre_to_sql(genre_define(token))
    print(genre_define(token))
    with open('../mysql/initdb.d/popular_movie.sql','w+') as f:
        f.write(sql)
    

if __name__ == "__main__":
    run_script()
