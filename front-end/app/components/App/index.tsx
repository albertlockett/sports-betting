import * as React from "react";
import { BrowserRouter as Router } from "react-router-dom";
import { Apollo } from "./apollo";
import Routes from "./routes";

import "./styles.scss";

export default function App(): JSX.Element {
  return (
    <Apollo>
      <div id="app-container">
        <Router>
          <Routes />
        </Router>
      </div>
    </Apollo>
  );
}
