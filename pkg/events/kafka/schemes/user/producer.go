package user

import (
	"context"
	user_eventsv1 "github.com/ravsanalbertovic-sketch/sch-general-lib/api/pb/v1/events/user"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/kafka/schemes"
	"github.com/twmb/franz-go/pkg/sr"
	"google.golang.org/protobuf/proto"
)

func InitUserProducer(ctx context.Context, registry *schemes.Registry, srURL string) error {
	client, err := sr.NewClient(sr.URLs(srURL))
	if err != nil {
		return err
	}
	configs := []struct {
		subject string
		msg     proto.Message
		getter  func() (string, error)
	}{
		{CreatedSubject, &user_eventsv1.UserCreated{}, events.GetUserCreatedProto},
		{DeletedSubject, &user_eventsv1.UserDeleted{}, events.GetUserDeletedProto},
	}
	for _, cfg := range configs {
		protoStr, err := cfg.getter()
		if err != nil {
			return err
		}
		schema, err := client.CreateSchema(ctx, cfg.subject, sr.Schema{
			Schema: protoStr,
			Type:   sr.TypeProtobuf,
		})
		if err != nil {
			return err
		}
		registry.RegisterEvent(schema.ID, cfg.msg)
	}
	return nil
}
