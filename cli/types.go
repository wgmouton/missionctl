package main

type About struct {
	Version string
	Test    string
}

type MissionTools struct {
	Variables map[string]string
	Tools     map[string]Tool
}

type Tool struct {
	Binary string
	Args   []string
}
