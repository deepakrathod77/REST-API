package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "project"
	session, err = cluster.CreateSession()
	cluster.Consistency = gocql.Quorum
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
func GetSession() *gocql.Session {
	return session
}
