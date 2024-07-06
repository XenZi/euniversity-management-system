export enum ToaletType {
  RoomShared = 0,
  FloorShared = 1,
  RoomBased = 2,
}

export enum ApplicationType {
  Budget = 1,
  SelfFinancing = 2,
  Disability = 3,
  SensitiveGroups = 4,
}

export enum ApplicationStatus {
  Review = 0,
  Accepted = 1,
  Denied = 2,
  Pending = 3,
}

export enum EDrugForm {
  SYRUP = 0,
  TABLET = 1,
  GEL = 2,
  CAPSULE = 3,
}

export enum EPrescriptionType {
  ISSUED = 0,
  CLAIMED = 1,
  ISSUED_REPEAT = 2,
  CLAIMED_REPEAT = 3,
}

export enum EAppointmentType {
  EXAMINATION = 0,
  INTERVENTION = 1,
}

export enum EAppointmentStatus {
  SCHEDULED = 0,
  COMPLETED = 1,
}

