package event_id

import (
	"github.com/google/uuid"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/ports/id_generator"
	"github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/vo/id"
)

type generator struct{}

func NewGenerator() id_generator.Generator {
	return &generator{}
}

func (g *generator) Generate() id.ID {
	return id.From16Bytes(uuid.New())
}
