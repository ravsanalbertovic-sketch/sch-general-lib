package user

import (
	user_eventsv1 "github.com/ravsanalbertovic-sketch/sch-general-lib/api/pb/v1/events/user"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/vo/id"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const DeletedEventTopicName = "user.deleted"

type UserDeleted struct {
	ID        id.ID
	UserID    string
	Email     string
	CreatedAt time.Time
}

func NewUserDeletedEvent(eventID id.ID, userID, email string, createdAt time.Time) *UserDeleted {
	return &UserDeleted{
		ID:        eventID,
		UserID:    userID,
		Email:     email,
		CreatedAt: createdAt,
	}
}

func (d *UserDeleted) EncodePb() *user_eventsv1.UserDeleted {
	return &user_eventsv1.UserDeleted{
		EventID:   d.ID.String(),
		UserID:    d.UserID,
		Email:     d.Email,
		CreatedAt: timestamppb.New(d.CreatedAt),
	}
}

func DecodePbToUserDeleted(pb *user_eventsv1.UserDeleted) (*UserDeleted, error) {
	eventID, err := id.Parse(pb.EventID)
	if err != nil {
		return nil, err
	}
	userDeletedEvent := NewUserDeletedEvent(eventID, pb.UserID, pb.Email, pb.CreatedAt.AsTime())
	return userDeletedEvent, nil
}

func (d *UserDeleted) GetID() id.ID {
	return d.ID
}

func (d *UserDeleted) Name() string {
	return DeletedEventTopicName
}

func (d *UserDeleted) AggregateType() string {
	return AggregateType
}

func (d *UserDeleted) AggregateID() string {
	return d.UserID
}
