import * as React from "react";
import { useQuery } from "@apollo/client";
import moment from "moment";

import Game from "./components/Game";

import { GET_GAME_LIST_FUNCTION } from "./queries";
import { pairGames } from "./utils";
import { selectHits } from "../../util/selectors";

import "./styles.scss";

export default function DailyGameListPage(): JSX.Element {
  const today = moment(new Date()).format("YYYY-MM-DDT00:00:00\\Z");
  const { data, loading, error } = useQuery(GET_GAME_LIST_FUNCTION, {
    variables: { date: today },
  });

  if (loading) {
    return <h1>Loading</h1>;
  }
  if (error) {
    return <h1>Error {error}</h1>;
  }

  const hits = selectHits(data).map(({ _source }) => _source);

  const gamePairs = pairGames(hits);

  return (
    <div className="daily-game-list">
      {gamePairs.map((gamePair) => {
        const key = gamePair.home.AwayTeam + gamePair.home.HomeTeam;
        return (
          <Game key={key} homeGame={gamePair.home} awayGame={gamePair.away} />
        );
      })}
    </div>
  );
}
