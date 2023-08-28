output=$(docker ps | grep mysql)
id=$(echo "$output" | awk '{print $1}')


docker rm -f $id

sudo rm -rf ../db_volume

docker-compose -f ../docker-compose.yaml up --force-recreate --build 
