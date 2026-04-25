package user

import (
	"context"
	"fmt"
	user_eventsv1 "github.com/ravsanalbertovic-sketch/sch-general-lib/api/pb/v1/events/user"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/infra/kafka/schemes"
	user2 "github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/infra/user"
	"github.com/twmb/franz-go/pkg/sr"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
)

const (
	CreatedSubject = user2.CreatedEventTopicName + "-value"
	DeletedSubject = user2.DeletedEventTopicName + "-value"
)

func DecodeRecord(
	ctx context.Context,
	client *sr.Client,
	registry *schemes.Registry,
	record []byte,
) (any, error) {
	id, _, err := registry.Serde.DecodeID(record)
	if err != nil {
		return nil, fmt.Errorf("failed to decode id: %w", err)
	}
	event := registry.NewInstanceById(id)
	if event == nil {
		key := strconv.Itoa(id)
		res, err, _ := registry.Sfg.Do(key, func() (interface{}, error) {
			schema, err := client.SchemaByID(ctx, id)
			if err != nil {
				return nil, err
			}
			return schema, nil
		})
		if err != nil {
			return nil, err
		}
		schemaText := res.(string)
		var msg proto.Message
		if strings.Contains(schemaText, "message UserCreated") {
			msg = &user_eventsv1.UserCreated{}
		} else if strings.Contains(schemaText, "message UserDeleted") {
			msg = &user_eventsv1.UserDeleted{}
		} else {
			return nil, fmt.Errorf("unknown message type in schema for id %d", id)
		}
		registry.RegisterEvent(id, msg)
		event = registry.NewInstanceById(id)
	}
	if err = registry.Serde.Decode(record, event); err != nil {
		return nil, err
	}
	return event, nil
}

/* реализация для confluent sr
func DecodeRecord(
	ctx context.Context,
	client *sr.Client,
	registry *schemes.Registry,
	record []byte,
) (any, error) {
	id, _, err := registry.Serde.DecodeID(record)
	if err != nil {
		return nil, fmt.Errorf("failed to decode id: %w", err)
	}
	event := registry.NewInstanceById(id)
	if event == nil {
		key := strconv.Itoa(id)
		res, err, _ := registry.Sfg.Do(key, func() (interface{}, error) {
			subjects, err := client.SubjectsByID(ctx, id)
			if err != nil {
				return nil, err
			}
			if len(subjects) == 0 {
				return nil, fmt.Errorf("id %d not found in registry", id)
			}
			return subjects, nil
		})
		if err != nil {
			return nil, err
		}
		subjects := res.([]string)
		var msg proto.Message
		found := false
		for _, s := range subjects {
			switch s {
			case CreatedSubject:
				msg = &user_eventsv1.UserCreated{}
				found = true
			case DeletedSubject:
				msg = &user_eventsv1.UserDeleted{}
				found = true
			}
			if found {
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("id %d has no matching subjects in local registry", id)
		}
		registry.RegisterEvent(id, msg)
		event = registry.NewInstanceById(id)
	}
	if err = registry.Serde.Decode(record, event); err != nil {
		return nil, err
	}
	return event, nil
}
*/
