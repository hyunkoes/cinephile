# !bin/bash
# Check if database ready
MAX_TRY=15
CURRENT_TRY=0

# Check parm
# if [ $# -eq 1 ]; then
#     if [ $1 == reset ]; then
#         echo "\033[33mUsing reset parm\033[0m"
#         docker rm -f dev_db
#         docker rm -f logger-db
#         docker rm -f Probrain_redis
#     else
#         echo "\033[31m Wrong parm..\033[0m"
#         exit -1;
#     fi
# fi

# Wrapper
function check_status {
    echo "\033[34m"
    bash -c "./wait-for-it.sh -s -t 60 $1"
    echo "\033[0m"
    if [ $? ==  0 ]; then
        echo "\033[32m ...Success\033[0m";
        #sleep 1
    else
        echo "\033[31m >> ...Failed\033[0m";
        exit -1;
    fi
}

env_value="$1"

if [ "$env_value" = "local" ]; then
    filename=".env.local"
elif [ "$env_value" = "dev" ]; then
    filename=".env.dev"
elif [ "$env_value" = "prod" ]; then
    filename=".env.prod"
else
    echo "Invalid 'env'"
    exit 1
fi

if [ -f $filename ]; then
    export $(cat $filename | sed 's/#.*//g' | xargs) # env file load
    echo "Using development $filename"
else
    echo "There is no $filename !!"
    exit -1
fi
export $(cat $filename | sed 's/#.*//g' | xargs) # env file load

# Docker ping tests
# Mysql ping test
until echo '\q' | docker exec cinephile_mysql_$1 mysql -h localhost -P"${MYSQL_PORT}" -u"${MYSQL_USER}" -p"${MYSQL_PASSWORD}" ${MYSQL_DATABASE}> /dev/null 2>&1; do
    if [ $CURRENT_TRY -eq $MAX_TRY ]; then
        echo "\033[31mMYSQL is unavailable - Abort\033[0m";
        exit -1;
    else
        >&2 echo "\033[33mMySQL is unavailable - sleeping 5sec\033[0m";
    sleep 5
    fi
    
    CURRENT_TRY=$((CURRENT_TRY+1))
done
echo "\033[32mMYSQL is ready\033[0m";

CURRENT_TRY=0
