package main

import (
	"fmt"
)

func main() {
	
	zoo := map[string]struct{}{}
	
	
	zoo["Walrus"] = struct{}{}
	zoo["Parrot"] = struct{}{}
	zoo["Lion"] = struct{}{}
	zoo["Giraffe"] = struct{}{}
	
	
	zoo["Bear"] = struct{}{}

	
	zoo["Lion"] = struct{}{}
	
	
	delete(zoo, "Parrot")
	
	fmt.Println(zoo)
	
	
}