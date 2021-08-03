import * as React from "react";
import classnames from "classnames";

type numberProps = {
  expectedValue: number;
  lineAmerican: number;
  lineDecimal: number;
  odds: number;
  side: string;
};

export default function Numbers(props: numberProps): JSX.Element {
  return (
    <div
      className={classnames("game-item-numbers", {
        home: "home" === props.side,
        away: "away" === props.side,
      })}
    >
      <div
        className={classnames("ev-number", {
          positive: props.expectedValue > 1,
          negative: props.expectedValue < 1,
        })}
      >
        <span className="value">{props.expectedValue.toFixed(4)}</span>
      </div>
      <div className="handicap-number">
        <span className="value">{(props.odds * 100).toFixed(2)}%</span>
      </div>
      <div className="line">
        <span className="line-american value">{props.lineAmerican}</span>
        {"/"}
        <span className="line-decimal value">
          {props.lineDecimal.toFixed(3)}
        </span>
      </div>
    </div>
  );
}
