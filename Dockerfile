FROM golang:1.11.1

#Determine Work directory
WORKDIR /go/src/go-book-database

#Install Go dep
RUN sudo apt install go-dep

#Copy dependencies file
COPY . .

#Install dependencies
RUN dep ensure
RUN go install

#Run the thing
CMD [ "go-book-database" ]