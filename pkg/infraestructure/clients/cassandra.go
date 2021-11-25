package clients

import (
	"sync"

	"github.com/gocql/gocql"
)

var (
	onceStartCassandra sync.Once
	instanceCassandra  *gocql.Session
)

func ConnectCassandra() *gocql.Session {
	onceStartCassandra.Do(func() {
		cluster := gocql.NewCluster("127.0.0.1")
		cluster.Keyspace = "books"
		session, err := cluster.CreateSession()
		if err != nil {
			panic("couldn't connect to db")
		}

		instanceCassandra = session
	})

	return instanceCassandra
}
