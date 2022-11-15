package factory_design_pattern

import (
	"fmt"

	"github.com/google/uuid"
)

type Job struct {
	Id          string
	Name        string
	Type        string
	Description string
}

func NewJob(name, jobType, description string) *Job {
	return &Job{
		Id:          uuid.NewString(),
		Name:        name,
		Type:        jobType,
		Description: description,
	}
}

func (job *Job) String() string {
	return fmt.Sprintf("(id: %s, name: %s, type: %s, description: %s)", job.Id, job.Name, job.Type, job.Description)
}
