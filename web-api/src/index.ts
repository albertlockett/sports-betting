import express = require("express");
import cors from "cors";
import { GraphQLSchema, GraphQLObjectType } from "graphql";
import elasticsearch from "elasticsearch";
import { elasticApiFieldConfig } from "graphql-compose-elasticsearch";

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

  const schema = new GraphQLSchema({
    query: new GraphQLObjectType({
      name: "Query",
      fields: {
        elastic7: elasticApiFieldConfig(
          // you may provide existed Elastic Client instance
          new elasticsearch.Client({
            host: "http://localhost:9200",
            apiVersion: "7.x",
          })
        ),
      },
    }),
  });

  const server = new ApolloServer({ typeDefs, resolvers, schema });
  await server.start();

  const app = express();
  app.use(cors())
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
