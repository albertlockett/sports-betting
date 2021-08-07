import { gql } from "@apollo/client";

export const GET_GAME_LIST_FUNCTION = gql`
  query get_daily_games($date: String) {
    elastic7 {
      search(
        index: "expected-values"
        size: 1000
        body: {
          query: {
            bool: {
              must: [
                { term: { Time: { value: $date } } }
                { term: { LatestCollected: { value: true } } }
              ]
            }
          }
        }
      )
    }
  }
`;
