import * as React from "react";

import { ExpectedValue } from "../types"

import AtSign from "./AtSign";
import Numbers from "./Numbers";
import TeamInfo from "./TeamInfo";

type gameProps = {
  homeGame: ExpectedValue;
  awayGame: ExpectedValue;
};

export default function Game(props: gameProps): JSX.Element {
  return (
    <div className="game-item">
      <Numbers
        expectedValue={props.awayGame.ExpectedValue}
        lineAmerican={props.awayGame.LineAmerican}
        lineDecimal={props.awayGame.LineDecimal}
        odds={props.awayGame.Odds}
        side="away"
      />
      <TeamInfo code={props.awayGame.AwayTeam} side="away" />
      <AtSign />
      <TeamInfo code={props.homeGame.HomeTeam} side="home" />
      <Numbers
        expectedValue={props.homeGame.ExpectedValue}
        lineAmerican={props.homeGame.LineAmerican}
        lineDecimal={props.homeGame.LineDecimal}
        odds={props.homeGame.Odds}
        side="home"
      />
    </div>
  );
}
