cd ..

docker run -d -v res/backdoor:/var/www/html --name php-running-razboynik php-razboynik
docker rm -f /php-running-razboynik
