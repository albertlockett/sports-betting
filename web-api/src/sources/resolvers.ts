import _ from "lodash";
import fetch from "node-fetch";

const ENDPOINT = "http://localhost:8080/";

async function fetchThingFromEndpoint(
  indexName: string,
  args: { input?: any }
) {
  const input = args?.input;
  const endpoint = `${ENDPOINT}${indexName}`;
  const response = await fetch(endpoint, {
    method: "POST",
    body: input ? JSON.stringify(input) : undefined,
  });
  const data = await response.json();
  const result = data.map((ev: any) =>
    _.mapKeys(ev, (v, k) => {
      return _.lowerFirst(k);
    })
  );
  return result;
}

export async function fetchExpectedValues(
  context: any,
  args: any,
  info: any
): Promise<any[]> {
  return fetchThingFromEndpoint("expected-values", args);
}

export async function fetchDailyValues(
  context: any,
  args: any,
  info: any
): Promise<any[]> {
  return fetchThingFromEndpoint("daily-summaries", args);
}
