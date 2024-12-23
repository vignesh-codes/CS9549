# Development Mode
FROM node:22

WORKDIR /app/client

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--host"]
