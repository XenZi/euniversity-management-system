export interface User {
  fullName: string;
  gender: string;
  identityCardNumber: string;
  citizenship: string;
  personalIdentificationNumber: string;
  residence: string;
  roles: string[];
}

export interface Residence {
  address: string;
  placeOfResidence: string;
  municipalityOfResidence: string;
  countryOfResidence: string;
}

export interface BirthData {
  dateOfBirth: string;
  municapilityOfBirth: string;
  countryOfBirth: string;
}
