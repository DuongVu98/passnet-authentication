package channels

var grpcChannel = make(chan interface{})
func GetGrpcChannel() chan interface{} {
	return grpcChannel
}