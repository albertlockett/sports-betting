import * as React from "react";
import { ApolloProvider } from "@apollo/client";
import { client } from "../../apollo-client";

type apolloProps = {
  children: JSX.Element;
};

export function Apollo(props: apolloProps): JSX.Element {
  return <ApolloProvider client={client}>{props.children}</ApolloProvider>;
}
