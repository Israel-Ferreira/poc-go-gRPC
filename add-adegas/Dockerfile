FROM golang:1.23.2-alpine AS build

WORKDIR /usr/app

COPY . /usr/app

RUN CGO_ENABLED=0 GOOS=linux go build -o add-adegas-api main.go


FROM scratch

COPY --from=build /usr/app/add-adegas-api   .

EXPOSE 4000

CMD ["./add-adegas-api"]