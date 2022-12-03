package tuto

type Task struct {
	Name string
	Done bool
}
type Tasklist struct {
	Name  string
	Tasks []Task
}

type Person struct {
	Name         string
	Age          int
	LikesGo      bool
	LikesNetwork bool
}
