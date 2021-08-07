import gql from "graphql-tag";
import { fetchExpectedValues } from "./resolvers";

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

  input SourcesSearchInput {
    latestCollected: Boolean
    time: String
  }

  extend type Query {
    ExpectedValues(input: SourcesSearchInput): [ExpectedValue]
  }
`;

export const queries = {
  ExpectedValues: fetchExpectedValues,
};
