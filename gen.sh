protoc -I. --go_out=plugins=grpc:$GOPATH/src/. apollo/apollo.proto
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. blackhole/blackhole.proto
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. core/core.proto
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. celestial/celestial.proto
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. gravity/gravity.proto
echo "Generating done."
