# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

FROM golang:1.19-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

# copy all files
COPY . ./

# compile application
RUN CGO_ENABLED=0 go build -o /microservices-go -ldflags="-w -s"

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

# copy executable from build image
COPY --from=build /microservices-go /microservices-go

# set gin to run in release mode
ENV GIN_MODE=release

# export application port
EXPOSE 8080

# command to be used to execute when the image is used to start a container
ENTRYPOINT [ "/microservices-go" ]
