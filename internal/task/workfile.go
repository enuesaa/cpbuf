package task

func NewWorkfile(filename string) Workfile {
	return Workfile{ filename: filename }
}

type Workfile struct {
	filename string
}
// get work file path
// copy to buf dir
