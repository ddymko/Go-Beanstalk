FROM golang:1.11.1-stretch as stage

WORKDIR /src
COPY . /src
RUN go build -o eb-go-app
EXPOSE 8081
CMD [ "./eb-go-app" ]

