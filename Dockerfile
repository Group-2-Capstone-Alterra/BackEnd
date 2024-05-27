FROM golang:1.22.2-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

# Debug: tampilkan isi direktori sebelum build
RUN ls -R /app

RUN go build -o PetPalApp .

CMD ["./PetPalApp"]
