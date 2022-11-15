package handlers

import (
	"fmt"
	. "github.com/study-onboard/go/ch08/factory_design_pattern"
)

type PackageJobHandler struct {
}

func (handler *PackageJobHandler) Handle(job *Job) error {
	fmt.Printf("PackageJobHandler - packaging for job: %s\n", job)
	fmt.Printf("PackageJobHandler - job has been packaged: %s\n", job)
	return nil
}

func init() {
	RegisterJobHandler("package", &PackageJobHandler{})
	fmt.Println("PackageJobHandler has been registered for type - package")
}
