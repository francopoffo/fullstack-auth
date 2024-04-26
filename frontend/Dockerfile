FROM node:latest AS build

WORKDIR /app

COPY package*.json ./

RUN npm i -g pnpm
RUN pnpm install

COPY . .

RUN pnpm build

EXPOSE 3000

CMD ["pnpm", "start"]
