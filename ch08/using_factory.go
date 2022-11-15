package ch08

import (
	"fmt"

	. "github.com/study-onboard/go/ch08/factory_design_pattern"
	_ "github.com/study-onboard/go/ch08/factory_design_pattern/handlers"
)

func UsingFactory() {
	job := NewJob(
		"Build codes",
		"build",
		"Build codes of project",
	)
	HandleJob(job)

	fmt.Println("----------------------------------------------------------------------")

	job = NewJob(
		"Package artifact",
		"package",
		"Package artifact for deployment",
	)
	HandleJob(job)

	fmt.Println("----------------------------------------------------------------------")

	job = NewJob(
		"Deploy artifact",
		"deploy",
		"Deploy artifact to cluster",
	)
	HandleJob(job)
}
