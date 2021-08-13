import { gql } from "@apollo/client";

export const GET_GAME_LIST_FUNCTION = gql`
  query expected_vals($time: String) {
    ExpectedValues(input: { time: $time }) {
      HomeTeam: homeTeam
      AwayTeam: awayTeam
      Time: time
      Odds: odds
      LineAmerican: lineAmerican
      LineDecimal: lineDecimal
      ExpectedValue: expectedValue
      Type: type
      EventId: eventId
      latestCollected: latestCollected
      Side: side
      TimeComputed: timeComputed
    }
  }
`;
