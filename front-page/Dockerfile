FROM node:current-alpine3.13 AS nodebuilder
COPY . /node/front-page
RUN cd /node/front-page && npm install && npm run build