package main
import (
"fmt"
)
func main(){
m :=make(map[string]int)
m["divya"]=40
m["chandu"]=30
m["gopi"]=50
for key,val :=range m{
fmt.Print("[",key,"->",val,"]")
}
fmt.Println("divy age:",m["divya"])
delete(m,"divya")
fmt.Println("divya age:",m["divya"])

v, ok :=m["divya"]
fmt.Println("divya age:",v, "Present:",ok)

v2,ok2 :=m["chandu"]
fmt.Println("chandu age:",v2,"Present:",ok2)
}