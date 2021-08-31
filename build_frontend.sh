#!/bin/bash
cd front-end && \
npm install && \
npm run build && \
docker build -t gcr.io/albertlockett-test2/sports-betting-front-end:latest  . && \
cd ../

