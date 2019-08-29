package internal

type TaskData struct{
	TaskID int
	Progress int
	Taken bool
	TakenAt int64
	Handling bool
}
