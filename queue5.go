package main
import (
"fmt"
)
type Queue struct{
items []int
}
func(q *queue) Enqueue(i int){
q.items=append(q.items,i)
}
func (q *Queue)Dequeue() int{

toRemove :=q.items[0]
q.items = q.items[l:]
return toRemove
}


func main() {
myQueue :=Queue{}
fmt.Println(myQueue)
myQueue.Enqueue(100)
myQueue.Enqueue(200)
myQueue.Enqueue(300)
myQueue.Dequeue()
fmt.Println(myQueue)
}
