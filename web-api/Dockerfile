FROM node:16-alpine3.14 AS build
WORKDIR /usr/src/app
COPY package*.json ./
COPY . .
RUN npm install
RUN npm run build

FROM node:16-alpine3.14
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install --only=production
COPY --from=build /usr/src/app/build ./build
COPY ./static/graphiql.html ./static/graphiql.html
EXPOSE 4000
CMD npm start