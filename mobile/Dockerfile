FROM ghcr.io/cirruslabs/flutter:stable

WORKDIR /app

COPY . .

RUN flutter pub get
RUN flutter clean

CMD ["sh", "-c", "flutter build apk --release && cp build/app/outputs/flutter-apk/app-release.apk /output/client.apk"]
