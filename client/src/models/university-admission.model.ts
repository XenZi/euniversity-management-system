import { University } from "./university.model";
import { User } from "./user.model";

export interface UniversityAdmission{
    id: string;
    citizen: User;
    dateAndTime: string;
    university: University;
}