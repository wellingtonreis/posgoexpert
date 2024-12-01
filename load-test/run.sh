docker build -t load-test .
docker run -it --rm load-test /app/cli --url=http://google.com --requests=100 --concurrency=10