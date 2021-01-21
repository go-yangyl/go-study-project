## json

```go
type FinishTasks struct {
	TaskId   int   `json:"task_id"`
	CreateAt int64 `json:"create_at"`
}

func FinishTasksMarshal(finishTask []*FinishTasks) string {
	str, err := json.Marshal(finishTask)
	if err != nil {
		return ""
	}
	return string(str)
}

func FinishTasksUnmarshal(str []byte) []*FinishTasks {
	finishTask := new([]*FinishTasks)
	err := json.Unmarshal(str, finishTask)
	if err != nil {
		return nil
	}
	return *finishTask
}
```