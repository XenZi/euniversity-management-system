export interface Dorm {
  id: string;
  name: string;
  location: string;
  prices: Prices[];
}

export interface Prices {
  applicationType: number;
  price: number;
}
