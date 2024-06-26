export interface Room {
  id: string;
  dormID: string;
  squareFoot: number;
  numberOfBeds: number;
  toalet: number;
}
export enum ToaletType {
  RoomShared = 0,
  FloorShared = 1,
  RoomBased = 2,
}
