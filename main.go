package main

import(
	"fmt"
	"flag"
)

func main(){
	add := flag.String("add", "N/A", "accepts task(string) and adds it to the list")
	rm := flag.Int("rm", '-1', "removes the tasks from the list. Takes in an index(int)")
	done := flag.Int("done", '-1',"marks the task as done. Accepts an index(int)")
	
	
}
