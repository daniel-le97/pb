FROM golang:latest

RUN apt add -v build-base
RUN apt add -v ca-certificates
RUN apt add -v git \
    unzip \
    openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

# Add this line to copy the "static" folder into the image


RUN go build -o pb-build

EXPOSE 8080

CMD ["/app/pb-build", "serve", "--http=0.0.0.0:8090"]