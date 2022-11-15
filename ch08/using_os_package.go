package ch08

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func UsingOsPackage() {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Can not get the name of host")
	} else {
		fmt.Printf("Host name: %s\n", hostName)
	}

	printLine()

	envs := os.Environ()
	fmt.Printf("envs's type: %T\n", envs)
	for _, env := range envs {
		fmt.Println(env)
	}

	printLine()

	javaHome := os.Getenv("JAVA_HOME")
	fmt.Println("JAVA_HOME: ", javaHome)
	goRoot := os.Getenv("GOROOT")
	fmt.Println("GOROOT: ", goRoot)
	os.Setenv("MASAGA_HOME", "China GD GZ HuaDu")
	fmt.Println("MASAGA_HOME: ", os.Getenv("MASAGA_HOME"))

	printLine()

	workspace, err := os.Getwd()
	if err != nil {
		fmt.Printf("Can not get workspace: %s\n", err.Error())
	} else {
		uid := os.Getuid()
		gid := os.Getgid()
		pid := os.Getpid()
		fmt.Printf("uid: %d, gid: %d, pid: %d, workspace: %s\n", uid, gid, pid, workspace)
	}

	printLine()

	dirName := "/tmp/abc"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Printf("Make directory `%s` failed - %s\n", dirName, err.Error())
	} else {
		fmt.Printf("Directory `%s` maked.\n", dirName)
	}
	os.Remove(dirName)

	dirName = "/tmp/a/b/c/d"
	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Printf("Make directory `%s` failed - %s\n", dirName, err.Error())
	} else {
		fmt.Printf("Directory `%s` maked.\n", dirName)
	}
	err = os.RemoveAll("/tmp/a")
	if err != nil {
		fmt.Printf("Remove all for /tmp/a failed: %s\n", err.Error())
	}

	commandFile, err := exec.LookPath("ls")
	if err == nil {
		fmt.Println(commandFile)
	}

	user, err := user.Current()
	if err == nil {
		log.Println("username:", user.Name)
		log.Println("uid:", user.Uid)
		log.Println("home dir:", user.HomeDir)
		log.Println("gid:", user.Gid)

		groups, err := user.GroupIds()
		if err == nil {
			log.Println("groups: ", groups)
		}
	} else {
		log.Printf("Get current user information fail: %s", err.Error())
	}
}

func printLine() {
	fmt.Println(">>-----------------------------------------------------------------<<")
}
