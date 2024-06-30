import { User } from "./user.model";

export interface Application {
  id?: string;
  dormitoryAdmissionsID?: string;
  applicationType?: number;
  verifiedStudent?: boolean;
  healthInsurance?: boolean;
  applicationStatus?: number;
  student?: User;
}
