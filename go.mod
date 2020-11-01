module github.com/agronskiy/golang-learning

go 1.15

replace google.golang.org/grpc/examples => ./

require (
	google.golang.org/grpc v1.33.1
	google.golang.org/grpc/examples v0.0.0-20201030163418-c8ef9bc95712 // indirect
	google.golang.org/protobuf v1.25.0
)
