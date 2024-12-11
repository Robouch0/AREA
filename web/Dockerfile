FROM node:18-alpine

WORKDIR /app

COPY package.json ./

RUN npm install

COPY . .

# Modify the env to the prod one
# Later do a more suitable switch from dev to prod (vice versa)
RUN mv .env.prod .env

EXPOSE 8081

## Change it to a full build for prod
CMD ["npm", "run", "dev"]
