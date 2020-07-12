package entity

import (
	"fmt"
	"testing"
)

func TestConfig_ToString(t *testing.T) {
	c := Config{
		Host:           "encuesta",
		HostName:       "34.226.105.63",
		User:           "ubuntu",
		IdentitiesOnly: "yes",
		IdentityFile:   "~/.ssh/ec2-test.pem",
		Port:           22,
	}
	fmt.Println(c.ToString())
}

func TestConfigurations_ToString(t *testing.T) {
	c1 := Config{
		Host:           "encuesta",
		HostName:       "34.226.105.63",
		User:           "ubuntu",
		IdentitiesOnly: "yes",
		IdentityFile:   "~/.ssh/ec2-test.pem",
		Port:           22,
	}
	c2 := Config{
		Host:           "encuesta",
		HostName:       "34.226.105.63",
		User:           "ubuntu",
		IdentitiesOnly: "yes",
		IdentityFile:   "~/.ssh/ec2-test.pem",
		Port:           22,
	}
	c3 := Config{
		Host:           "encuesta",
		HostName:       "34.226.105.63",
		User:           "ubuntu",
		IdentitiesOnly: "yes",
		IdentityFile:   "~/.ssh/ec2-test.pem",
		Port:           22,
	}
	cgs := new(Configs)
	//cgs.Cgs[0] = c1
	//cgs.Cgs[1] = c2
	//cgs.Cgs[2] = c3
	cgs.Cgs = []Config{c1, c2, c3}
	fmt.Println(cgs.ToString())
}
