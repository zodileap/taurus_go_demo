package cmd

import (
	"fmt"
	"testing"

	"github.com/zodileap/taurus_go/cmd"
)

func TestNew(t *testing.T) {
	c := cmd.New("mkdir", "dir")
	c.Must()
}

func TestSetBaseCmd(t *testing.T) {
	cmd.SetBaseCmd("sudo")
	c := cmd.New("mkdir", "dir").String()
	fmt.Println(c)
}

func TestNewUser(t *testing.T) {
	_, err := cmd.New("cat", "/etc/passwd").Run()
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(r))
	c := cmd.NewUser("root").AddCmd("mkdir3213", "dir")
	fmt.Println(c.String())
	cmd.SetOuterr(true)
	c.Must()
}

func TestSplit(t *testing.T) {
	s := cmd.Split("mkdir dir")
	c := cmd.New(s...)
	fmt.Println(c.String())
	c.Must()
}
