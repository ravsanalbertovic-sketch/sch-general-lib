package user

const DeletedEventTopicName = "user.deleted"

type UserDeleted struct {
	UserID string
	Email  string
}

func (d UserDeleted) Name() string {
	return DeletedEventTopicName
}
