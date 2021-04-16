# League Backend Challenge

## Run with
		go run .
## Send request with:
		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
		curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
		curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
		curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
		curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
