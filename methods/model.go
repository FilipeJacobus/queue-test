package methods

import (
	// "fmt"

	"time"

	"generic"

	"github.com/gocql/gocql"
)

var session *gocql.Session

func initCql() *gocql.Session {

	cluster := gocql.NewCluster(generic.HOSTS)
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: generic.USERNAME, Password: generic.PASSWORD}
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return session
}

func insertData(stc PersistRequest) {

	defer initCql().Close()

	if err := session.Query(`INSERT INTO telemetry (id, latitude, longitude) VALUES (?, ?, ?)`,
		stc.CliID, stc.Latitude, stc.Longitude).Exec(); err != nil {
		panic(err)
	}
}
