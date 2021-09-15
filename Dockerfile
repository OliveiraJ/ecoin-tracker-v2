FROM golang:1.17-alpine3.13 as builder
RUN mkdir /app
WORKDIR /app
COPY . /app/
VOLUME /app/data
EXPOSE 10000
RUN go build -o /bin/ecointracker
CMD ["/bin/ecointracker"]
