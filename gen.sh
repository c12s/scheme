protoc -I. --go_out=plugins=grpc:$GOPATH/src/. celestial/celestial.proto
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. blackhole/blackhole.proto
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. core/core.proto

echo "Generating done."
