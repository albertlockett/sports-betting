import * as React from "react";
import classNames from "classnames";

type teamInfoProps = {
  code: string;
  side: string;
};

interface NameMap {
  [name: string]: string;
}

const NAME_FOR_CODE: NameMap = {
  BAL: "Baltimore Orioles",
  BOS: "Boston Red Sox",
  NYY: "New York Yankees",
  TBD: "Tampa Bay Rays",
  TOR: "Toronto Blue Jays",
  CSW: "Chicago White Sox",
  CLE: "Cleveland Guardians",
  DET: "Detroit Tigers",
  KCR: "Kansas City Royals",
  MIN: "Minnesota Twins",
  HOU: "Houston Astros",
  ANA: "Los Angeles Angels",
  OAK: "Oakland Athletics",
  SEA: "Seattle Mariners",
  TEX: "Texas Rangers",
  ATL: "Atlanta Braves",
  FLA: "Miami Marlins",
  NYM: "New York Mets",
  PHI: "Philadelphia Phillies",
  WSN: "Washington Nationals",
  CHC: "Chicago Cubs",
  CIN: "Cincinnati Reds",
  MIL: "Milwaukee Brewers",
  PIT: "Pittsburgh Pirates",
  STL: "St Louis Cardinals",
  ARZ: "Arizona Diamondbacks",
  COL: "Colorado Rockies",
  LAD: "Los Angeles Dodgers",
  SDP: "San Diego Padres",
  SFG: "San Francisco Giants",
};

function nameForCode(code: string): string {
  return NAME_FOR_CODE[code];
}

export default function TeamInfo(props: teamInfoProps): JSX.Element {
  return (
    <div
      className={classNames("team-info", {
        home: "home" === props.side,
        away: "away" === props.side,
      })}
    >
      <div className="name">{nameForCode(props.code)}</div>
    </div>
  );
}
