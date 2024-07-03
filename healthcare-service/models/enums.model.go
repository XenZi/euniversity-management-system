package models

import (
	"fmt"
)

type EForm int

const (
	Syrup EForm = iota
	Tablet
	Gel
	Capsule
)

func (f EForm) String() string {
	switch f {
	case Syrup:
		return "SYRUP"
	case Tablet:
		return "TABLET"
	case Gel:
		return "GEL"
	case Capsule:
		return "CAPSULE"
	default:
		return fmt.Sprintf("Unknown Form value: %d", int(f))
	}
}

type EPrescriptionStatus int

const (
	Issued EPrescriptionStatus = iota
	Claimed
	IssuedRepeat
	ClaimedRepeat
)

func (f EPrescriptionStatus) String() string {
	switch f {
	case Issued:
		return "ISSUED"
	case Claimed:
		return "CLAIMED"
	case IssuedRepeat:
		return "ISSUED_REPEAT"
	case ClaimedRepeat:
		return "CLAIMED_REPEAT"
	default:
		return fmt.Sprintf("Unknown PrescriptionStatus value: %d", int(f))
	}
}

type EAppointmentType int

const (
	Examination EAppointmentType = iota
	Intervention
)

func (f EAppointmentType) String() string {
	switch f {
	case Examination:
		return "EXAMINATION"
	case Intervention:
		return "INTERVENTION"
	default:
		return fmt.Sprintf("Unknown AppointmentType value: %d", int(f))
	}
}

type EAppointmentStatus int

const (
	Scheduled EAppointmentStatus = iota
	Completed
)

func (f EAppointmentStatus) String() string {
	switch f {
	case Scheduled:
		return "SCHEDULED"
	case Completed:
		return "COMPLETED"
	default:
		return fmt.Sprintf("Unknown AppointmentStatus value: %d", int(f))
	}
}
