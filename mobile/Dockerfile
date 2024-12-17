FROM ghcr.io/cirruslabs/flutter:stable

WORKDIR /app

COPY pubspec.yaml pubspec.lock ./

RUN flutter pub get

COPY . .

RUN flutter clean
RUN flutter pub get

CMD ["sh", "-c", "flutter build apk --release && cp build/app/outputs/flutter-apk/app-release.apk /output/client.apk"]