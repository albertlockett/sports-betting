import * as React from "react";

import AtSign from "./AtSign";
import Numbers from "./Numbers";
import TeamInfo from "./TeamInfo";

type gameProps = {
  game: {
    HomeTeam: string;
    AwayTeam: string;
    Time: string;
    Odds: number;
    LineAmerican: number;
    LineDecimal: number;
    Type: string;
    LatestCollected: boolean;
    Side: string;
    ExpectedValue: number;
  };
};

export default function Game(props: gameProps): JSX.Element {
  // TODO 2 different games need be passed in
  return (
    <div className="game-item">
      <Numbers
        expectedValue={props.game.ExpectedValue}
        lineAmerican={props.game.LineAmerican}
        lineDecimal={props.game.LineDecimal}
        odds={props.game.Odds}
        side="away"
      />
      <TeamInfo code={props.game.AwayTeam} side="away" />
      <AtSign />
      <TeamInfo code={props.game.HomeTeam} side="home" />
      <Numbers
        expectedValue={props.game.ExpectedValue}
        lineAmerican={props.game.LineAmerican}
        lineDecimal={props.game.LineDecimal}
        odds={props.game.Odds}
        side="home"
      />
    </div>
  );
}
