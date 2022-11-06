package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func PbToJson(message proto.Message) (string, error) {

	data, err := protojson.Marshal(message)

	return string(data), err
}
