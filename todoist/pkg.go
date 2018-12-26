package todoist

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/levigross/grequests"
	"github.com/mchirico/harvest/configure"
	"strconv"
)

type DueStruct struct {
	Recurring bool   `json:"recurring"`
	String    string `json:"string"`
	Date      string `json:"date"`
	DateTime  string `json:"datetime"`
	Timezone  string `json:"timezone"`
}

type TaskStruct struct {
	Id           int       `json:"id"`
	ProjectID    int       `json:"project_id"`
	Content      string    `json:"content"`
	Complete     bool      `json:"complete"`
	Label        []int     `json:"label_ids"`
	Order        int       `json:"order"`
	Indent       int       `json:"indent"`
	Priority     int       `json:"priority"`
	CommentCount int       `json:"comment_count"`
	Due          DueStruct `json:"due"`
	Url          string    `json:"url"`
}

func (t *TaskStruct) ID() string {
	return strconv.Itoa(t.Id)
}

func GetData(url string) (string, error) {

	token := configure.GetTODOtoken()

	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %v", token.Token)
	ro.Headers = headers

	result, err := grequests.Get(url, &ro)

	return result.String(), err
}

/*
curl "https://beta.todoist.com/API/v8/tasks" \
    -X POST \
    --data '{"content": "Appointment with Maria", "due_string": "tomorrow at 12:00", "due_lang": "en", "priority": 4}' \
    -H "Content-Type: application/json" \
    -H "X-Request-Id: $(uuidgen)" \
    -H "Authorization: Bearer $token"
*/

// --data '{"content": "Appointment with Maria", "due_string": "tomorrow at 12:00", "due_lang": "en", "priority": 4}' \
// 2971041043
func CreateTask(content string,
	due_string string, priority int) (TaskStruct, error) {

	token := configure.GetTODOtoken()

	uid := uuid.New()
	url := "https://beta.todoist.com/API/v8/tasks"

	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["X-Request-Id"] = uid.String()
	headers["Authorization"] = fmt.Sprintf("Bearer %v", token.Token)
	ro.Headers = headers

	type Task struct {
		Content  string `json:"content"`
		Due      string `json:"due_string"`
		Lang     string `json:"due_lang"`
		Priority int    `json:"priority"`
	}

	ts := TaskStruct{}

	task := Task{}
	task.Content = content
	task.Due = due_string
	task.Lang = "en"
	task.Priority = priority

	ro.JSON = task

	result, err := grequests.Post(url, &ro)
	fmt.Println(result.String())

	json.Unmarshal([]byte(result.String()), &ts)

	return ts, err

}

func UpdateTask(id string, content string,
	due_string string, priority int) (string, error) {

	token := configure.GetTODOtoken()

	uid := uuid.New()
	url := fmt.Sprintf("https://beta.todoist.com/API/v8/tasks/%s", id)

	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	headers["X-Request-Id"] = uid.String()
	headers["Authorization"] = fmt.Sprintf("Bearer %v", token.Token)
	ro.Headers = headers

	type Task struct {
		Content  string `json:"content"`
		Due      string `json:"due_string"`
		Lang     string `json:"due_lang"`
		Priority int    `json:"priority"`
	}

	task := Task{}
	task.Content = content
	task.Due = due_string
	task.Lang = "en"
	task.Priority = priority

	data, err := json.Marshal(task)

	if err != nil {
		return "", err
	}

	ro.JSON = task
	fmt.Println(string(data))

	result, err := grequests.Post(url, &ro)

	return result.String(), err

}

func DeleteTask(id string) (*grequests.Response, error) {
	token := configure.GetTODOtoken()
	url := fmt.Sprintf("https://beta.todoist.com/API/v8/tasks/%s", id)
	//url = "http://httpbin.org/delete"

	ro := grequests.RequestOptions{}
	headers := map[string]string{}
	//headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %v", token.Token)
	ro.Headers = headers

	result, err := grequests.Delete(url, &ro)

	return result, err
}

func GetAllTasks() ([]TaskStruct, error) {
	url := "https://beta.todoist.com/API/v8/tasks"
	r, _ := GetData(url)

	records := make([]TaskStruct, 0)
	err := json.Unmarshal([]byte(r), &records)

	return records, err
}
