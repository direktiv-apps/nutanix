#!/bin/sh

docker build -t nutanix . && docker run -p 9191:8080 nutanix