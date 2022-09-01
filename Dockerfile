FROM golang:1.19

#create Folder
RUN  mkdir -p /usr/src/app

#project path
WORKDIR /usr/src/app

#source to target
COPY go.mod go.sum ./
#download need packet
RUN go mod download && go mod verify


COPY . .
RUN go build -v -o demo_server

#

#

ENTRYPOINT ["/usr/src/app/demo_server"]