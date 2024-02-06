docker-compose down
docker build -t golang-my-market-api .
docker-compose up -d
docker rmi $(docker images -qa -f 'dangling=true')
