for plugin in `ls $1`
do
    echo "Build $1/$plugin"
    ./build.sh $1 $plugin
done
