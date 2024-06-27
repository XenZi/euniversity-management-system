import { User } from "./user.model";

export interface Application {
  id?: string;
  dormitoryAdmission?: string;
  applicationType?: number;
  verifiedStudent?: boolean;
  healthInsurance?: boolean;
  applicationStatus?: number;
  student?: User;
}

/*
	ID                    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DormitoryAdmissionsID string             `json:"dormitoryAdmissionsID" bson:"dormitoryAdmissionsID"`
	ApplicationType       ApplicationType    `json:"applicationType" bson:"applicationType"`
	VerifiedStudent       bool               `json:"verifiedStudent" bson:"verifiedStudent"`
	HealthInsurance       bool               `json:"healthInsurance" bson:"healthInsurance"`
	ApplicationStatus     ApplicationStatus  `json:"applicationStatus" bson:"applicationStatus"`
	Student               Student            `json:"student" bson:"student"`
*/

/*
	Review ApplicationStatus = iota
	Accepted
	Denied
	Pending
*/
