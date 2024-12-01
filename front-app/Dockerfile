FROM node:22

WORKDIR /home/node/app
RUN npm install --location=global @angular/cli

COPY package*.json ./
RUN npm install

CMD ["ng", "serve","--host", "0.0.0.0"]

EXPOSE 4200
