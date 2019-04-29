docker container stop jarvisnews
docker container rm jarvisnews
docker run -d --name jarvisnews \
    -v $PWD/cfg:/app/jarvisnews/cfg \
    -v $PWD/logs:/app/jarvisnews/logs \
    -v $PWD/dat:/app/jarvisnews/dat \
    -p 7053:7053 \
    -p 7788:7788 \
    jarvisnews