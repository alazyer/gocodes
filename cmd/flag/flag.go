package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	namespace string
	project   string
	region    string
)

func main() {
	basicFlag()
	basicFlagSet()
}

func basicFlag() {
	// flag.XXXVar example
	flag.StringVar(&namespace, "namespace", "alauda", "Name of the organization")
	flag.StringVar(&project, "project", "system", "Name of the project")
	flag.StringVar(&region, "region", "ace", "Name of the region")

	// flag.XXX example
	var username = flag.String("username", "ylzhang", "username in organization")

	// flag.Var example, with flag.Value interface
	flagVar := new(FlagVar)
	flag.Var(flagVar, "flag-var", "custom flag var test")
	// flag.Parse(os.Args[1:])
	flag.Parse()

	// Output
	fmt.Printf("Args: %v", flag.Args())
	fmt.Printf("Namespace: %v, Project: %v, Region: %v\n", namespace, project, region)
	fmt.Printf("Username: %s\n", *username)
	fmt.Printf("Custom FlagVar: %s\n", *flagVar)

}

func basicFlagSet() {
	fmt.Println("Flagset example")
	projectCmd := flag.NewFlagSet("project", flag.ExitOnError)
	projectEnable := projectCmd.Bool("enable", false, "enable")
	projectName := projectCmd.String("name", "abc", "name")

	// go run ./flag.go project --name 123
	projectCmd.Parse(os.Args[2:])

	// Output
	fmt.Printf("Args: %v\n", projectCmd.Args())
	fmt.Printf("projectName: %v\n", *projectName)
	fmt.Printf("projectEnable: %v\n", *projectEnable)
}

type FlagVar string

var _ flag.Value = new(FlagVar)

func (f *FlagVar) Set(value string) error {
	*f = FlagVar(value)
	return nil
}

func (f *FlagVar) String() string {
	return string(*f)
}

func (f *FlagVar) Type() string {
	return "flagVar"
}
