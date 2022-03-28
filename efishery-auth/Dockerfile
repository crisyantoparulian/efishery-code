FROM node:10.16.3-stretch

WORKDIR /app

COPY . /app

RUN npm install

EXPOSE 8081
CMD [ "npm", "run", "start" ]