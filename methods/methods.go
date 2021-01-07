package methods

import (
	"container/list"
)

func (StringService) Persist(c PersistRequest, queue *list.List) (string, error) {

	queue.PushBack(c)
	if queue.Len() >= 10 {

		//create struct list
		var q []PersistRequest
		for i := 0; i < queue.Len(); i++ {
			e := queue.Front()
			q = append(q, e.Value.(PersistRequest))

		}

		//add list to memory address
		qm := &q

		//clear queue
		for i := 0; i < queue.Len(); i++ {
			e := queue.Front()
			queue.Remove(e)
		}

		//persist list
		ch := make(chan bool)
		go persistList(qm, ch)
		if <-ch {
			return "success", nil
		}
	}

	return "success", nil
}

func persistList(q *[]PersistRequest, ch chan bool) {
	for _, v := range *q {
		insertData(v)
	}
	ch <- true
}
