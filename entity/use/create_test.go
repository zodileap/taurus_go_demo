package use

import (
	"testing"

	"github.com/google/uuid"
	"github.com/yohobala/taurus_go/tlog"

	entity "taurus_go_demo/entity/gen"
)

func TestCreate(t *testing.T) {
	db, err := entity.NewUser()
	defer db.Close()
	if err != nil {
		t.Errorf(err.Error())
	}
	u, err := db.Users.New(
		uuid.New().String(),
		"test",
		db.Users.WithDesc("test"),
	)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = db.Save()
	if err != nil {
		t.Errorf(err.Error())
	}
	tlog.Print(u.ID)
	tlog.Print(u.CreatedTime)
	tlog.Print(u.Desc)
}
