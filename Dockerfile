FROM golang:1.20-alpine as build

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . /app
EXPOSE 3030
##buat executeable
RUN go build -o todo .

##jalankan executeable
CMD ["/app/todo"]
