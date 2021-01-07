package methods

import (
	"time"

	"generic"
)

func (StringService) Persist(c PersistRequest) (string, error) {
	//restart timer
	generic.CL <- true

	//add items in queue
	generic.QUEUE.PushBack(c)

	//checks if the queue is full
	if generic.QUEUE.Len() >= generic.MAXQUEUELEN {

		//starts data persistence
		_, err := persistDataConc()
		if err != nil {
			panic(err)
		}

	}

	return "success", nil
}

func persistDataConc() (string, error) {
	//create struct list
	var q []PersistRequest
	for generic.QUEUE.Len() > 0 {
		e := generic.QUEUE.Front()
		q = append(q, e.Value.(PersistRequest))
		//clear queue
		generic.QUEUE.Remove(e)

	}

	//add list to memory address
	qm := &q

	//persist list
	// ch := make(chan bool)
	go persistList(qm)

	return "success", nil
}

func persistList(q *[]PersistRequest) {
	for _, v := range *q {
		insertData(v)
	}

}

func Timer() {

	select {
	//restart timer
	case <-generic.CL:
		Timer()
	//if the call interval exceeds the allowed time it persists the data
	case <-time.After(generic.TIMETOPERSIST):
		if generic.QUEUE.Len() > 0 {
			persistDataConc()
		}
		Timer()
	}

}
