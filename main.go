package main

import(
	"fmt"
	"flag"
	"encoding/json"
	"os"
)

type Task struct{
	Id int 				`json:"id"`
	Name string 		`json:"name"`
	Status string 		`json:"status"`

}

func main(){

	add := flag.String("add", "N/A", "accepts task(string) and adds it to the list")
	rm := flag.Int("rm", -1, "removes the tasks from the list. Takes in an index(int)")

	flag.Parse()

	_ = rm
		
	data := getData()

	if *add != "N/A" {
		fmt.Println("Added:", *add)
		newTask := Task{len(data), *add, "due"}
		data = append(data, newTask)
		updateData(data)
	}

	if *rm != -1{
		fmt.Println("Removed:", *rm)
		data = append(data[:*rm], data[*rm+1:]...)
		updateData(data)
	}
	formatPrintData(data)
}


func formatPrintData(tasks []Task){
	for _, value := range tasks{
		fmt.Printf("%v  |   %s \n", value.Id, value.Name)
	}
}

func updateData(tasks []Task){
		b, err:= json.Marshal(tasks)
		if err != nil{
			return
		}

		errr := os.WriteFile("task.json", b, 0666)
		if errr != nil{
			fmt.Println("Cant write to task.json:",errr)
		}
}
func getData() []Task{

	var tasks []Task
	data, err := os.ReadFile("task.json")
	if err != nil {
		return tasks
	}

	err = json.Unmarshal(data, &tasks) // Check for Unmarshal errors
	if err != nil {
		return tasks
	}

	return tasks
}
