FROM golang:1.16-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o /main main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /main /main
USER nonroot:nonroot
ENTRYPOINT ["/main"]