import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  useQuery,
  gql,
} from "@apollo/client";

// eslint-disable-next-line 
const ENDPOINT = process.env.ENDPOINT || "http://localhost:4000/graphql"

export const client = new ApolloClient({
  uri: ENDPOINT,
  cache: new InMemoryCache(),
});
