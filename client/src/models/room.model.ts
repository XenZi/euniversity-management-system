import { User } from "./user.model";

export interface Room {
  id: string;
  dormID: string;
  squareFoot: number;
  numberOfBeds: number;
  toalet: number;
  students: User[];
}
