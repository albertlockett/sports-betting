/* eslint-disable */
const express = require("express");
const path = require("path");
const app = express();

// eslint-disable-next-line
const DIRNAME = __dirname;

console.log(path.join(DIRNAME, "dist"));

const fuck = path.join(DIRNAME,"..", "dist")
app.use("/",            express.static(fuck));
app.get('/daily-games', function(req, res) {
  res.sendFile(path.join(fuck+ '/index.html'));
});

app.listen(3000, () => {
  // eslint-disable-next-line no-undef
  console.log("make app go on internet");
});
