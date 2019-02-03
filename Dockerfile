FROM golang:1.11.1

#Determine Work directory
WORKDIR /go/src/app

#Install Go dep
RUN go get -u github.com/golang/dep/cmd/dep

#Copy dependencies file
COPY . .

#Install dependencies
RUN dep ensure
RUN go install

#Run the thing
CMD [ "app" ]