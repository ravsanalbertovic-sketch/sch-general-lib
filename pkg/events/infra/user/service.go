package user

import (
	"fmt"
	user_eventsv1 "github.com/ravsanalbertovic-sketch/sch-general-lib/api/pb/v1/events/user"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/ports/event"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/vo/id"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const AggregateType = "user"

func ConvertPBToEvent(pb proto.Message) (event.Event, error) {
	switch v := pb.(type) {
	case *user_eventsv1.UserCreated:
		eventID, err := id.Parse(v.EventID)
		if err != nil {
			return nil, err
		}
		return NewUserCreatedEvent(eventID, v.UserID, v.Email, v.CreatedAt.AsTime()), nil
	case *user_eventsv1.UserDeleted:
		eventID, err := id.Parse(v.EventID)
		if err != nil {
			return nil, err
		}
		return NewUserDeletedEvent(eventID, v.UserID, v.Email, v.CreatedAt.AsTime()), nil
	default:
		return nil, fmt.Errorf("unknown proto type")
	}
}

func ConvertEventToPB(event event.Event) (proto.Message, error) {
	switch v := event.(type) {
	case *UserCreated:
		return &user_eventsv1.UserCreated{
			EventID:   v.ID.String(),
			UserID:    v.UserID,
			Email:     v.Email,
			CreatedAt: timestamppb.New(v.CreatedAt),
		}, nil
	case *UserDeleted:
		return &user_eventsv1.UserDeleted{
			EventID:   v.ID.String(),
			UserID:    v.UserID,
			Email:     v.Email,
			CreatedAt: timestamppb.New(v.CreatedAt),
		}, nil
	default:
		return nil, fmt.Errorf("unknown event type")
	}
}
