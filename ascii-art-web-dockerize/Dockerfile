FROM golang:1.22.3-alpine
WORKDIR /home/asciiArtWeb
COPY . .
RUN go build -o asciiArtWeb .
EXPOSE 8080
CMD ["./asciiArtWeb"]
