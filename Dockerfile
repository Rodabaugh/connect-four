FROM debian:stable-slim
WORKDIR /app
COPY connect-four /app/
COPY static /app/static/
RUN chmod +x /app/connect-four

RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates

CMD ["./connect-four"]