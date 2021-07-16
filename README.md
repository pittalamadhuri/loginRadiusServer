# loginRadiusServer

This server has one post API which takes a text file as input and returns the 10 most repeated words with their occurence count.

## Installation

Clone the repository, then install the requirements and start the web server.

    $ go get
    $ go run main.go
this will start the server on port 8080

## API
>* Type: POST <br>
>* URL: http://localhost:8080/radius/ <br>
>* Content-Type: multipart/form-data <br>
>* name: textFile <br>

### Postman link
>* https://www.getpostman.com/collections/34013650c8cc90158dc7


