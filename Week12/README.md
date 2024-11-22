# Cloud Student Portfolio

## How to run the Go application:

```bash
git clone https://github.com/sojoudian/CLO2003_fall24.git
cd CLO2003_fall24/Week10
go run main.go
```

## How to build the Cloud Engineer Portfolio Docker image
```bash
git pull
docker build -t cloud_portfolio:0.1 .
```

## How to run the Cloud Engineer Portfolio Docker image
```bash
docker run -p 80:80 cloud_portfolio:0.1
```
# how to push the image to the docker hub
```bash
docker login
docker tag cloud_portfolio:0.1 maziar/cloud_portfolio:0.1
docker push maziar/cloud_portfolio:0.1
```
## Build for all UNIX like Operating Systems

```bash
docker buildx build --platform linux/amd64,linux/arm64,windows/amd64,linux/arm/v7 -t maziar/myapp:latest --push .
````



