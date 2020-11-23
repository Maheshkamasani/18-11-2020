package main
import (
"fmt"
)
func main1(){
st :=set.New()
st.Insert(1)
st.Insert(2)
fmt.Println(st.Has(1))
fmt.Println(st.Has(3))
}