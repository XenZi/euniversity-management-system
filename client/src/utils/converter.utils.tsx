export const RemoveQuotationMarksFromString = (inputString: string): string => {
  return inputString.substring(1, inputString.length - 1);
};

export const castFromApplicationTypeNumberToActualString = (
  num: number
): string => {
  switch (num) {
    case 1:
      return "Budget";
    case 2:
      return "Self financing";
    case 3:
      return "Disability";
    case 4:
      return "Sensitive groups";
    default:
      return "Unkown";
  }
};

export const castFromToaletTypeNumberToActualString = (num: number): string => {
  switch (num) {
    case 0:
      return "Room shared";
    case 1:
      return "Floor shared";
    case 2:
      return "Room based";
    default:
      return "Unkown";
  }
};

export const castFromApplicationStatusToActualString = (
  num: number
): string => {
  switch (num) {
    case 0:
      return "Review";
    case 1:
      return "Accepted";
    case 2:
      return "Denied";
    case 3:
      return "Pending";
    default:
      return "Unkown";
  }
};
