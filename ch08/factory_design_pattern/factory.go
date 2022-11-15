package factory_design_pattern

import "fmt"

type JobHandler interface {
	Handle(job *Job) error
}

var factoryMapping map[string]JobHandler = make(map[string]JobHandler)

func RegisterJobHandler(jobType string, handler JobHandler) {
	factoryMapping[jobType] = handler
}

func HandleJob(job *Job) {
	if handler, ok := factoryMapping[job.Type]; ok {
		handler.Handle(job)
	} else {
		fmt.Printf("Warning: can't find handler for job: %s\n", job)
	}
}
