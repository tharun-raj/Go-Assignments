package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
	"time"
)
 
var taskId = 0

//Creating a Struct Task 
type Task struct{
	id int
	todo string
	isDone bool
	date string
}

func getInput(cmd string) string{
	fmt.Printf("%s : ", cmd)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

//create addTask Method for struct Task
func addTask(t *[]Task)  {
	fmt.Println("--------- ADD TASK ---------")
	taskId += 1
	var localTask Task
	todo := getInput("What is the task")
	currentTime := time.Now() 
	localTask.id = taskId 
	localTask.todo = todo
	localTask.date = currentTime.Format("2006-01-02")
	*t = append(*t, localTask)
	
}
func listTask(t []Task)  {
	fmt.Println("---------- List of Tasks ---------")
	for _,value := range(t){
		fmt.Println()
		fmt.Printf("id:\t%d\ntask:\t%sisDone:\t%t\nDate:\t%s\n", value.id, value.todo, value.isDone, value.date)
		fmt.Println()
	}

}

func isDone(t []Task) {
	fmt.Println("---------- Is Done  ---------")
	input := getInput("Enter the ID")
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil{
		fmt.Println("Enter a valid Id")
	} else{
		for i, v := range(t){
			if v.id == id{
				v.isDone = true
				t[i] = v
				break
			}
		}
	}
	

}

func deleteTask(t *[]Task) {
	var updTask []Task
	input := getInput("Enter the ID")
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil{
		fmt.Println("Enter a valid Id")
	} else{
		for _, v := range(*t){
			if v.id == id{
				continue
			}
			updTask = append(updTask, v)
		}
	}
	*t = updTask

}

// Main function to display the menu and handle user input
func main() {
   var todoTasks []Task
   for {
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Mark Task as Done")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Exit")
 
        input :=  getInput("Enter choice")
        choice, err := strconv.Atoi(strings.TrimSpace(input))
 
        if err != nil {
            fmt.Println("Invalid choice, please try again.")
            continue
        }

		switch(choice){
		case 1: addTask(&todoTasks)
		case 2: listTask(todoTasks)
		case 3: isDone(todoTasks)
		case 4: deleteTask(&todoTasks)
		case 5: os.Exit(0)
		default:
		}
    }
}
 