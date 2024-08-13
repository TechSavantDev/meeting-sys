FROM node

WORKDIR /app

RUN npm install -g create-vite@latest

EXPOSE 9000
