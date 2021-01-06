package methods

import (
	"container/list"
	"fmt"
)

func (StringService) Persist(c PersistRequest, queue *list.List) (string, error) {

	queue.PushBack(c)
	if queue.Len() >= 10 {

		for i := 0; i < queue.Len(); i++ {
			e := queue.Front()
			fmt.Println(i+1, "persist >", e.Value)
			stc := e.Value.(PersistRequest)

			//implement goroutines and channels here ↓↓
			insertData(stc)
		}
		//clear queue || TODO:implement correct method
		for i := 0; i < queue.Len(); i++ {
			e := queue.Front()
			queue.Remove(e)
		}

	}

	return "success", nil
}
