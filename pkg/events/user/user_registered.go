package user

const CreatedEventTopicName = "user.created"

type UserCreated struct {
	UserID    string
	Email     string
	CreatedAt string
}

func (r UserCreated) Name() string {
	return CreatedEventTopicName
}
