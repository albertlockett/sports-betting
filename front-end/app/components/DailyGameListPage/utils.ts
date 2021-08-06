import { ExpectedValue, ExpectedValuePair } from "./types";

export function pairGames(games: ExpectedValue[]): ExpectedValuePair[] {
  const pairs = games.reduce(
    (acc, curr) => {
      const key = curr.AwayTeam + curr.HomeTeam;
      const pair = acc[key] ?? [];
      pair.push(curr);
      acc[key] = pair;
      return acc;
    },
    {} as {
      [_: string]: ExpectedValue[];
    }
  );
  return Object.values(pairs).map((games) => {
    const [game1, game2] = games;
    if (game1.Side === "home") {
      return { home: game1, away: game2 };
    } else {
      return { home: game2, away: game1 };
    }
  });
}
