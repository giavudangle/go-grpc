package serializer_test

import (
	"testing"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/giavudangle/go-grpc/sample"
	"github.com/giavudangle/go-grpc/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop_one := sample.NewLaptop()
	laptop_two := &pb.Laptop{}

	err := serializer.WriteProtoBufferToBinaryFile(laptop_one, binaryFile)
	require.NoError(t, err)

	err = serializer.ReadProtoBufferFromBinaryFile(binaryFile, laptop_two)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop_one, laptop_two))

	err = serializer.WriteProtoBufferToJSONFile(laptop_one, jsonFile)
	require.NoError(t, err)
}
