package entity

import "fmt"

type iConfig interface {
	ToString() string
}

type (
	//Config Patch the ssh config file see (~/.ssh/config)
	Config struct {
		Id             int
		Host           string `ssh:"Host"`
		HostName       string
		User           string
		IdentitiesOnly string
		IdentityFile   string
		Port           int16
		iConfig
	}

	//Configs List of configs on config file
	Configs struct {
		Cgs []Config
		iConfig
	}
)

func (c *Config) ToString() string {
	return fmt.Sprintf(
		`Host %s
    Hostname %s
    User %s
    IdentitiesOnly %s
    IdentityFile %s
    Port %d`,
		c.Host,
		c.HostName,
		c.User,
		c.IdentitiesOnly,
		c.IdentityFile,
		c.Port)
}

func (c *Configs) isHostRepeated() bool {
	return false
}

func (c *Configs) ToString() (s string) {
	for _, p := range c.Cgs {
		s += p.ToString() + "\n"
	}
	return
}
