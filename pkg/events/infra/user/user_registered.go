package user

import (
	user_eventsv1 "github.com/ravsanalbertovic-sketch/sch-general-lib/api/pb/v1/events/user"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/vo/id"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const CreatedEventTopicName = "user.created"

type UserCreated struct {
	ID        id.ID
	UserID    string
	Email     string
	CreatedAt time.Time
}

func NewUserCreatedEvent(eventID id.ID, userID, email string, createdAt time.Time) *UserCreated {
	return &UserCreated{
		ID:        eventID,
		UserID:    userID,
		Email:     email,
		CreatedAt: createdAt,
	}
}

func (r *UserCreated) EncodePb() *user_eventsv1.UserCreated {
	return &user_eventsv1.UserCreated{
		EventID:   r.ID.String(),
		UserID:    r.UserID,
		Email:     r.Email,
		CreatedAt: timestamppb.New(r.CreatedAt),
	}
}

func DecodePbToUserCreated(pb *user_eventsv1.UserCreated) (*UserCreated, error) {
	eventID, err := id.Parse(pb.EventID)
	if err != nil {
		return nil, err
	}
	userCreatedEvent := NewUserCreatedEvent(eventID, pb.UserID, pb.Email, pb.CreatedAt.AsTime())
	return userCreatedEvent, nil
}

func (r *UserCreated) GetID() id.ID {
	return r.ID
}

func (r *UserCreated) Name() string {
	return CreatedEventTopicName
}

func (r *UserCreated) AggregateType() string {
	return AggregateType
}

func (r *UserCreated) AggregateID() string {
	return r.UserID
}
