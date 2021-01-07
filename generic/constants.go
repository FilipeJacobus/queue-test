package generic

import (
	"container/list"
	"os"
	"time"
)

const (
	// maximum interval to persist queue rules
	TIMETOPERSIST = 5 * time.Second

	//maximum number of items allowed in the queue
	MAXQUEUELEN = 10
)

var (
	//init channel responsible for clearing the queue
	CL = make(chan bool)

	//init queue
	QUEUE = list.New()

	//cassandra credentials
	HOSTS    = os.Getenv("CASSANDRA_CLUSTER")
	USERNAME = os.Getenv("CASSANDRA_USERNAME")
	PASSWORD = os.Getenv("CASSANDRA_PASSWORD")
)
