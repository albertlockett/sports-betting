import gql from "graphql-tag";
import { fetchDailyValues, fetchExpectedValues } from "./resolvers";

export const typeDefs = gql`
  type ExpectedValue {
    homeTeam: String
    awayTeam: String
    time: String
    odds: Float
    lineAmerican: Int
    lineDecimal: Float
    expectedValue: Float
    type: String
    eventId: String
    latestCollected: Boolean
    side: String
    timeComputed: String
  }

  type DailySummary {
    time: String
    numGames: Int
  }

  input SourcesSearchInput {
    latestCollected: Boolean
    time: String
  }

  extend type Query {
    DailySummary(input: SourcesSearchInput): [DailySummary]
    ExpectedValues(input: SourcesSearchInput): [ExpectedValue]
  }
`;

export const queries = {
  ExpectedValues: fetchExpectedValues,
  DailySummary: fetchDailyValues,
};
