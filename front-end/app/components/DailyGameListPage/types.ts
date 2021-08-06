export type ExpectedValue = {
  Side: string;
  HomeTeam: string;
  AwayTeam: string;
  Time: string;
  Odds: number;
  LineAmerican: number;
  LineDecimal: number;
  Type: string;
  LatestCollected: boolean;
  ExpectedValue: number;
};

export type ExpectedValuePair = {
  home: ExpectedValue;
  away: ExpectedValue;
};
