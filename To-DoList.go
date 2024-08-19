package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ToDo struct{
	title string
	description string
	date string
}

var toDoList []ToDo

func displayToDoList(){
	fmt.Println("TO-DO LIST:")
	fmt.Println("-------------------------------------")
	for i:=0; i< len(toDoList); i++{
		fmt.Printf("%d.",i+1)
		fmt.Printf("\tTITLE: %s\n", toDoList[i].title)
		fmt.Printf("\tDESCRIPTION: %s\n", toDoList[i].description)
		fmt.Printf("\tDATE: %s\n", toDoList[i].date)
		fmt.Println("-------------------------------------")
	}
}

func addToDo(command string){

	fields := strings.Split(command, "|")
	if len(fields) != 2{return}
	
	var toDo ToDo
	toDo.title= strings.TrimSpace(fields[0])
	toDo.description = strings.TrimSpace(fields[1])
	toDo.date = time.Now().Format("2006-01-02 15:04")

	toDoList = append(toDoList, toDo)
}

func removeToDo(command string){
	fields := strings.Split(command, " ")
	if len(fields) != 2{return}

	num, _ := strconv.Atoi(fields[1])
	num--

	if num<0 || num > len(toDoList){return}

	toDoList = append(toDoList[:num], toDoList[num+1:]...)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("This is my To-Do List CLI application")
	fmt.Println("To enter a To-Do enter it with this format: \"title | description\"")
	fmt.Println("To display your To-Do List enter: \"display\"")
	fmt.Println("To remove a To-Do enter: \"remove idx\"  ")
	fmt.Println("To terminate application enter: \"exit\"")

	var continueFlag bool = true 
	for continueFlag{
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if strings.Compare(command, "display")==0 {
			displayToDoList()
		} else if strings.Compare(command, "exit")==0{
			continueFlag = false
		} else if strings.Contains(command, "remove") {
			removeToDo(command)
		} else {
			addToDo(command)
		}

	}

}