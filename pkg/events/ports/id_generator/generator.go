package id_generator

import "github.com/ravsanalbertovic-sketch/sch-general-lib/pkg/events/vo/id"

type Generator interface {
	Generate() id.ID
}
