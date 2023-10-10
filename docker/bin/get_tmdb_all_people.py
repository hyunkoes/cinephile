import sys
import time
import requests
import os
from dotenv import load_dotenv

# 영화 정보를 담을 array
movies = []
# 인물 정보를 담을 dictionary
characters = {}

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
def get_last_person_id(token):
    url = "https://api.themoviedb.org/3/person/latest"
    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(token)
    }
    response = requests.get(url, headers=headers).json()
    
    return  response['id']
def get_person(token, id):
    url = "https://api.themoviedb.org/3/person/{}?language=en-US".format(id)

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(token)
    }

    response = requests.get(url, headers=headers).json()
    return response
    
def person_to_sql(person):
    
    person_sql = 'INSERT INTO person VALUES ({},"{}","{}","{}","{}",{},{},"{}","{}");\n'
    birth = person.get('birthday', '') or ''
    death = person.get('deathday', '') or ''
    
    return person_sql.format(
            person['id'],person['name'],person['biography'].replace('"', "'"),
            birth, death, person['gender'],
            person['adult'], person['profile_path'], person['known_for_department']
        ).replace('""',"null")
def get_movie_list(page):
    url = "https://api.themoviedb.org/3/movie/popular?language={}&page={}".format("ko-KR",page)

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(token)
    }
    response = requests.get(url, headers=headers).json()
    return  response['results']
def getCreditInfo(movieID):
    url = "https://api.themoviedb.org/3/movie/{}/credits?language=en-US".format(movieID)

    headers = {
        "accept": "application/json",
        "Authorization": "Bearer {}".format(token)
    }
    response = requests.get(url, headers=headers).json()
    cast = response['cast']
    directors = [crew_member for crew_member in response['crew'] if crew_member['job'] == 'Director']
    return cast, directors

def appendPersonInfo(people):
    sql = ""
    for person in people :
        if characters.get(person['id'], False):
            continue
        personInfo = get_person(token, person['id'])
        characters[person['id']] = 1
        sql += person_to_sql(personInfo)
    return sql

def roleQuery(person, movie_id):
    # movieID, personID, role, character, order, name
    person_sql = 'INSERT INTO role VALUES ({},{},"{}","{}",{},"{}");\n'
    if person.get('job', False):
        if person['job'] == 'Directing':
            return person_sql.format(movie_id, person['id'], person['known_for_department'], person['department'], 1, person['name'])
    else:
        if person.get('character', False):
            return person_sql.format(movie_id, person['id'], person['known_for_department'], person['character'], person['order'], person['name'])
    return ""

def relationBetweenPersonMovie(people, movie_id):
    sql = ""
    for person in people:
        try:
            sql += roleQuery(person, movie_id)
        except Exception as e:
            print(e)
    return sql

def run_script():
    global movies
    global characters
    for i in range(1, 500):
        next_movies = get_movie_list(i)
        if ( len(next_movies) == 0 ):
            break
        movies = movies+next_movies
        time.sleep(1)
    person_sql = ""
    relation_sql = ""
    for i in range(0, len(movies)):
        cast, directors = getCreditInfo(movies[i]['id'])
        person_sql += appendPersonInfo(cast)
        person_sql += appendPersonInfo(directors)
        relation_sql += relationBetweenPersonMovie(cast, movies[i]['id'])
        relation_sql += relationBetweenPersonMovie(directors, movies[i]['id'])

    with open('../mysql/initdb.d/popular_people.sql','w+') as f:
        f.write(person_sql)
    with open('../mysql/initdb.d/popular_movie_credits.sql','w+') as f:
        f.write(relation_sql)
    
if __name__ == "__main__":
    global token 
    token = get_token()
    run_script()
    

