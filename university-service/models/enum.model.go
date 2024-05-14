package models

import "fmt"

type StudentStatus int

const (
	FullTimeStudent StudentStatus = iota
	PartTimeStudent
)

func (s StudentStatus) String() string {
	switch s {
	case FullTimeStudent:
		return "FULLTIME"
	case PartTimeStudent:
		return "PARTTIME"
	default:
		return fmt.Sprintf("Unknown Form value: %d", int(s))
	}

}

type ProfessorStatus int

const (
	FullTImeProfessor ProfessorStatus = iota
	PartTimeProfessor
)

func (p ProfessorStatus) String() string {
	switch p {
	case FullTImeProfessor:
		return "FULLTIME"
	case PartTimeProfessor:
		return "PARTTIME"
	default:
		return fmt.Sprintf("Unknown Form value: %d", int(p))
	}

}
