FROM ghcr.io/cirruslabs/flutter:stable

WORKDIR /app

COPY pubspec.yaml pubspec.lock ./

RUN flutter pub get

COPY . .

# Modify the env to the prod one
# Later do a more suitable switch from dev to prod (vice versa)
RUN mv .env.prod .env

RUN flutter clean
RUN flutter pub get

CMD ["sh", "-c", "flutter build apk --release && cp build/app/outputs/flutter-apk/app-release.apk /output/client.apk"]