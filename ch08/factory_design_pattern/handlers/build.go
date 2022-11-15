package handlers

import (
	"fmt"
	. "github.com/study-onboard/go/ch08/factory_design_pattern"
)

type BuildJobHandler struct {
}

func (handler *BuildJobHandler) Handle(job *Job) error {
	fmt.Printf("BuildJobHandler - buiding for job: %s\n", job)
	fmt.Printf("BuildJobHandler - job %s built", job)
	return nil
}

func init() {
	RegisterJobHandler("build", &BuildJobHandler{})
	fmt.Println("BuildJobHandler has been registered for type - build")
}
