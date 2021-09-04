FROM node:16-alpine3.14 AS build
ARG ENDPOINT
ENV ENDPOINT=$ENDPOINT
RUN apk add make
RUN apk add g++
RUN npm install -g npm@7.21.1
RUN apk add --no-cache python3 py3-pip
WORKDIR /fuck
COPY . ./
RUN npm install
RUN npm rebuild
RUN npm run build
CMD ["npm", "run", "serve"]

# FROM node:16-alpine3.14 
# RUN apk add thttpd
# RUN adduser -D static
# USER static
# WORKDIR /home/static
# COPY --from=build /fuck/dist ./
# CMD ["thttpd", "-D", "-h", "0.0.0.0", "-p", "3000", "-d", "/home/static", "-u", "static", "-l", "-", "-M", "60"]

