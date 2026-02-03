package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

type Base struct {
	CreatedByID 	ID
	UpdatedByID		ID
	created_at 		time.Time
	updated_at 		*time.Time
}

type ID string

func NewID() ID {
	return ID(ksuid.New().String())
}

func NewIDFromString(idStr string) (ID,error){
	parsedId,err := ksuid.Parse(idStr)
	if err!=nil {
		return "",err
	}

	return ID(parsedId.String()),nil
}

func (id ID) String() string {
	return string(id)
}

func (id ID) MarshalJSON() ([]byte,error) {
	return json.Marshal(string(id))
}

func (id ID) IsValid() bool {
	_, err := ksuid.Parse(string(id))
	return err == nil
}

func (id *ID) IsNil() bool {
	return id == nil || string(*id) == ""
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var idStr string
	if err := json.Unmarshal(data, &idStr); err != nil {
		return err
	}
	parsedID, err := ksuid.Parse(idStr)
	if err != nil {
		return fmt.Errorf("invalid id")
	}
	*id = ID(parsedID.String())
	return nil
}