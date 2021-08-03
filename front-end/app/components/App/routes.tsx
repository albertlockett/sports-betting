import * as React from "react";
import { Switch, Route } from "react-router-dom";
import { CALENDAR_PAGE_ROUTE, DAILY_GAME_LIST_PAGE } from "./contants";
import CalendarPage from "../CalendarPage";
import DailyGameListPage from "../DailyGameListPage";

export default function (): JSX.Element {
  return (
    <Switch>
      <Route path={DAILY_GAME_LIST_PAGE}>
        <DailyGameListPage />
      </Route>
      <Route path={CALENDAR_PAGE_ROUTE}>
        <CalendarPage />
      </Route>
    </Switch>
  );
}
