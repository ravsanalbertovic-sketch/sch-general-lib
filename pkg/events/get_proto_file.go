package events

import (
	"github.com/ravsanalbertovic-sketch/sch-general-lib/api/proto/v1/events"
)

func GetUserCreatedProto() (string, error) {
	b, err := events.FS.ReadFile("user/created.proto")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GetUserDeletedProto() (string, error) {
	b, err := events.FS.ReadFile("user/deleted.proto")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
