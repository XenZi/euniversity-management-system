package models

type ApplicationForDorm struct {
	ApplicationType   ApplicationType   `json:"applicationType"`
	VerifiedStudent   bool              `json:"verifiedStudent"`
	HealthInsurance   bool              `json:"healthInsurance"`
	ApplicationStatus ApplicationStatus `json:"applicationStatus"`
	Student           Student           `json:"student"`
}
