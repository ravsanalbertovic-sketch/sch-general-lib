package id

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
)

type ID struct {
	value [16]byte
}

func From16Bytes(b [16]byte) ID {
	return ID{value: b}
}

func (id ID) Bytes() [16]byte {
	return id.value
}

func (id ID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		id.value[0:4], id.value[4:6], id.value[6:8], id.value[8:10], id.value[10:16],
	)
}
func Parse(s string) (ID, error) {
	uid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return From16Bytes(uid), nil
}

func (id ID) Value() (driver.Value, error) {
	return id.String(), nil
}

func (id *ID) Scan(value interface{}) error {
	if value == nil {
		id.value = [16]byte{}
		return nil
	}
	var s string
	switch v := value.(type) {
	case string:
		s = v
	case []byte:
		s = string(v) // преобразуем []byte в строку
	default:
		return fmt.Errorf("cannot scan %T into ID", value)
	}
	//u, err := uuid.Parse(s)
	u, err := Parse(s)
	if err != nil {
		return fmt.Errorf("cannot parse UUID from string: %w", err)
	}
	id.value = u.value
	//copy(id.value[:], u[:])
	return nil
}
