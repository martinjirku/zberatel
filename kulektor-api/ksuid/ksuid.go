package ksuid

import (
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/segmentio/ksuid"
)

type KSUID ksuid.KSUID

// Scan implements the sql.Scanner interface.
func (k *KSUID) Scan(src interface{}) error {
	var b string
	switch s := src.(type) {
	case []byte:
		b = string(s)
	case string:
		b = s
	default:
		return fmt.Errorf("not a string or []byte: %#T", s)
	}
	ks, err := ksuid.Parse(b)
	if err != nil {
		return err
	}
	*k = KSUID(ks)
	return nil
}

// Value implements the driver.Valuer interface.
func (k KSUID) Value() (driver.Value, error) {
	return ksuid.KSUID(k).String(), nil
}

func NewKSUID() KSUID {
	id, _ := ksuid.NewRandom()
	return KSUID(id)
}

func (k *KSUID) UnmarshalGQL(src interface{}) error {
	var (
		ks  ksuid.KSUID
		err error
	)
	switch v := src.(type) {
	case []byte:
		ks, err = ksuid.FromBytes(v)
	case string:
		fmt.Printf("test %s", v)
		ks, err = ksuid.Parse(v)
	default:
		return fmt.Errorf("expected []byte|string, got %T", src)
	}
	if err != nil {
		return err
	}
	*k = KSUID(ks)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (k KSUID) MarshalGQL(w io.Writer) {
	w.Write([]byte(`"` + ksuid.KSUID(k).String() + `"`))
}
