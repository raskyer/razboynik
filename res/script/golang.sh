cd ..

docker run -d --name php-running-razboynik php-razboynik
docker run -it --rm -v ~/Documents/dev/razboynik:/go/src/github.com/eatbytes/razboynik --name go-running-razboynik go-razboynik
docker rm -f /php-running-razboynik
