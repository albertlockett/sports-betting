import process from "process";

export default {
  endpoint: process.env["ENDPOINT"] || "http://localhost:8080/",
};
