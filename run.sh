#!/bin/sh

env="$1"
case "$1" in
	dev) # dev env up
        docker-compose -f "docker/docker-compose.dev.yaml" up  -d --build cinephile_mysql certbot
        ./check-db-ready.sh $env
        docker-compose -f "docker/docker-compose.dev.yaml" up --build cinephile_server nginx
	;;
	deploy)
		#release env up
        docker-compose -f "docker/docker-compose.prod.yaml" up -d --build cinephile_mysql certbot
        ./check-db-ready.sh $env
        docker-compose -f "docker/docker-compose.prod.yaml" up -d --build cinephile_server nginx
	;;
    local)
        docker-compose -f "docker/docker-compose.local.yaml" up -d --build cinephile_mysql
        ./check-db-ready.sh $env
    ;;
	down) # env down
		if [ "$2" == "" ] ; then
			docker-compose -f "docker/docker-compose.yaml" down
		else
			echo "-down : No option"
		fi
	;;
	*) # exception
	echo "'$1' is unknown command + '$2'"
	;;
esac
