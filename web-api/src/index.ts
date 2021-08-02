import express = require("express");
import { ApolloServer, gql } from "apollo-server-express";

async function start() {
  // Construct a schema, using GraphQL schema language
  const typeDefs = gql`
    type Query {
      hello: String
    }
  `;

  // Provide resolver functions for your schema fields
  const resolvers = {
    Query: {
      hello: () => "Hello world!",
    },
  };

  const server = new ApolloServer({ typeDefs, resolvers });
  await server.start();

  const app = express();
  server.applyMiddleware({ app });

  const staticPath = __dirname + "/../static";
  app.get("/graphiql.html", express.static(staticPath));

  await new Promise((resolve) =>
    app.listen({ port: 4000 }, () => {
      resolve(null);
    })
  );
  console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`);
  return { server, app };
}

start();
