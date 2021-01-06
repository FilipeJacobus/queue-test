package generic

import "os"

var (
	//cassandra credentials
	HOSTS    = os.Getenv("CASSANDRA_CLUSTER")
	USERNAME = os.Getenv("CASSANDRA_USERNAME")
	PASSWORD = os.Getenv("CASSANDRA_PASSWORD")
)
