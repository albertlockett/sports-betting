import _ from "lodash";
import fetch from "node-fetch";

export async function fetchExpectedValues(
  context: any,
  args: { input?: any },
  info: any
): Promise<any[]> {
  const input = args?.input;
  const endpoint = "http://localhost:8080/expected-values";
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
