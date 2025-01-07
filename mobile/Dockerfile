FROM ghcr.io/cirruslabs/flutter:stable

WORKDIR /app

COPY . .

# Modify the env to the prod one
# Later do a more suitable switch from dev to prod (vice versa)
RUN mv .env.prod .env

RUN flutter pub get
RUN flutter clean

CMD ["sh", "-c", "flutter build apk --release"]