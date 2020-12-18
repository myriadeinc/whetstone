FROM golang:1.15.6-buster
WORKDIR /usr/src/app
COPY ./ /usr/src/app
RUN useradd gouser && usermod -a -G root gouser && mkdir /home/gouser
RUN chown gouser:gouser /home/gouser
RUN chown gouser:gouser /usr/src/app

USER gouser
RUN mkdir -p build
RUN go mod download
RUN go build -o build/whetstone

ENTRYPOINT [ "./build/whetstone" ]