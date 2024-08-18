# Code_League

League Home code challenge

To Run the code 
go run .


To test code 
go test ./...


## Supported Endpoints
- `/echo` - Returns the matrix as is.
    curl --location --request GET 'http://localhost:8080/echo' \
    --form 'file=@"path/matrix.csv"
- `/invert` - Returns the transposed version of the matrix.
    curl --location --request GET 'http://localhost:8080/invert' \
    --form 'file=@"path/matrix.csv"
- `/flatten` - Returns the matrix as a single line string.
    curl --location --request GET 'http://localhost:8080/flatten' \
    --form 'file=@"path/matrix.csv"
- `/sum` - Returns the sum of all integers in the matrix.
    curl --location --request GET 'http://localhost:8080/sum' \
    --form 'file=@"path/matrix.csv"
- `/multiply` - Returns the product of all integers in the matrix.
    curl --location --request GET 'http://localhost:8080/multiply' \
    --form 'file=@"path/matrix.csv"
