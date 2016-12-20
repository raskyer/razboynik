cd ..

docker run --rm -v /res/backdoor:/var/www/html/t --name php-running-razboynik php-razboynik
#docker rm -f /php-running-razboynik
