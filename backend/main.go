package main
import (
	e "easy_ssh/ssh/entity"
	"easy_ssh/storage"
	"fmt"
	"log"
)


func main() {

	var cfgs e.Configs
	var cf e.Config

	c, err := storage.SqlLite("easy_ssh.db")
	if nil != err {
		log.Println(err)
	}

	r, err := c.Query("SELECT * FROM hosts")
	if nil != err {
		log.Println(err)
	}

	for r.Next() {
		err := r.Scan(
			&cf.Id,
			&cf.Host,
			&cf.HostName,
			&cf.User,
			&cf.IdentitiesOnly,
			&cf.IdentityFile,
			&cf.Port,
		)
		if nil != err {
			log.Println(err)
		}
		cfgs.Cgs = append(cfgs.Cgs, cf)
	}
	fmt.Println(cfgs.ToString())
}
