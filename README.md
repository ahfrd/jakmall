# Guide

## Build image docker 
```
docker build -t ahfrd/jm:v1 .
```

## Run jm image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name jm-v1 ahfrd/jm:v1
```# AQ
# jm
# jakmall
