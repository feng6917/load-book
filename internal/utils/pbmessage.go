package utils

import (
	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func Unmarshal2Message(to proto.Message, args ...interface{}) error {
	ur := &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	for _, i := range args {
		b, _ := json.Marshal(i)
		if err := ur.Unmarshal(bytes.NewReader(b), to); err != nil {
			return err
		}
	}
	return nil
}
