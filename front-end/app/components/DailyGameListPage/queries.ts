import { gql } from "@apollo/client";

export const GET_GAME_LIST_FUNCTION = gql`
  query get_daily_games($date: String) {
    elastic7 {
      search(
        index: "expected-values"
        body: {
          query: {
            bool: {
              must: [
                { term: { Time: { value: $date } } }
                { term: { LatestCollected: { value: false } } }
              ]
            }
          }
        }
      )
    }
  }
`;
