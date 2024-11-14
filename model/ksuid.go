package model

import (
	"database/sql/driver"
	"fmt"

	"github.com/segmentio/ksuid"
)

type KSUID ksuid.KSUID

// Scan implements the sql.Scanner interface.
func (k *KSUID) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}
	ks, err := ksuid.FromBytes(b)
	if err != nil {
		return err
	}
	*k = KSUID(ks)
	return nil
}

// Value implements the driver.Valuer interface.
func (k KSUID) Value() (driver.Value, error) {
	return ksuid.KSUID(k).Bytes(), nil
}

func NewKSUID() KSUID {
	id, _ := ksuid.NewRandom()
	return KSUID(id)
}
