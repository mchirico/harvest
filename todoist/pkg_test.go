package todoist

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetAllTasks(t *testing.T) {
	url := "https://beta.todoist.com/API/v8/tasks"
	r, _ := GetData(url)
	fmt.Println(r)

	records := make([]TaskStruct, 0)
	json.Unmarshal([]byte(r), &records)

	log.Printf("here.... %v, %v\n", records[0].Id, len(records))
}

func TestGetAllProjects(t *testing.T) {

	url := "https://beta.todoist.com/API/v8/projects"
	r, _ := GetData(url)
	fmt.Println(r)

	records := make([]TaskStruct, 0)
	json.Unmarshal([]byte(r), &records)

	log.Printf("here.... %v, %v\n", records[0].Id, len(records))

}

func TestCreateTask(t *testing.T) {

	tr, err := CreateTask("Test Task Added", "today at 2:00", 4)
	if err != nil {
		t.Fail()
	}
	fmt.Println(tr.Id)

	time.Sleep(4 * time.Second)

	if tr.Id == 0 {
		t.Fail()
	}
	_, err = DeleteTask(tr.ID())
	if err != nil {
		t.Fail()
	}

}

func TestUpdateTask(t *testing.T) {

	r, err := UpdateTask("2971041043", "Modify task", "today at 2:00", 4)
	fmt.Println(r)
	fmt.Println(err)

}

func TestDeleteTask(t *testing.T) {

	r, err := DeleteTask("2971118902")

	fmt.Println(r.String())
	fmt.Println(err)

}
