export interface FoodCard {
  id: string;
  student_pin: string;
  expires: string; // ISO date format
  used_point: number[]; // Assuming this is an array of points
  messroom_name: string;
  mass_room_id: string;
  balance: number;
}
