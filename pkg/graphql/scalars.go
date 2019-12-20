package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

var (
	msgOK  = &Message{"Success!"}
	msgErr = &Message{"Fail!"}
)

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(id.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch id := v.(type) {
	case string:
		return uuid.Parse(id)
	case []byte:
		return uuid.ParseBytes(id)
	default:
		return uuid.Nil, fmt.Errorf("%[1]T(%[1]v) is not uuid", v)
	}
}
