package schemes

import (
	"github.com/twmb/franz-go/pkg/sr"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
	"reflect"
	"sync"
)

type Registry struct {
	Serde *sr.Serde
	mu    sync.RWMutex
	types map[int]reflect.Type
	Sfg   singleflight.Group
}

func NewRegistry() *Registry {
	return &Registry{
		Serde: sr.NewSerde(),
		types: make(map[int]reflect.Type),
	}
}

func (r *Registry) RegisterEvent(id int, msg proto.Message) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Serde.Register(
		id,
		msg,
		sr.Index(0),
		sr.EncodeFn(func(v any) ([]byte, error) { return proto.Marshal(v.(proto.Message)) }),
		sr.DecodeFn(func(b []byte, v any) error { return proto.Unmarshal(b, v.(proto.Message)) }),
	)
	r.types[id] = reflect.TypeOf(msg).Elem()
}

// NewInstanceById создает пустой объект по ID схемы
func (r *Registry) NewInstanceById(id int) any {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.types[id]
	if !ok {
		return nil
	}
	return reflect.New(t).Interface()
}
