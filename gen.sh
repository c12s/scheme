protoc blackhole/blackhole.proto -I. --go_out=plugin=grpc:.
protoc core/core.proto -I. --go_out=plugin=grpc:.
protoc celestial/celestial.proto -I. --go_out=plugin=grpc:.
echo "Generating done."
