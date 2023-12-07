FROM node:18.16.0-slim as builder

ARG REACT_APP_SERVER_URL
ENV REACT_APP_SERVER_URL=$REACT_APP_SERVER_URL

WORKDIR /app

COPY package*.json ./ 

RUN npm install

COPY ./src ./src

COPY ./public ./public

COPY ./tsconfig.json ./tsconfig.json

RUN npm run build 

FROM nginx:1.23

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=builder /app/build .

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]