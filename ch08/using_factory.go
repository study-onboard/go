package ch08

import (
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
}
