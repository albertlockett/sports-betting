package betdsi

import (
  "encoding/json"
  "fmt"
  "github.com/albertlockett/sports-betting/sources/model"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
  "time"
)

const _SOURCE = "betdsi.com"
const _URL_AUTH_CUSTOMER = "https://retro-ii.betdsi.eu/cloud/api/System/authenticateCustomer"
const _URL_GET_LEAGUE_LINES = "https://retro-ii.betdsi.eu/cloud/api/Lines/Get_LeagueLines"

var teamCodes = map[string]string {
  // AL East
  "Baltimore Orioles": "BAL",
  "Boston Red Sox": "BOS",
  "New York Yankees": "NYY",
  "Tampa Bay Rays": "TB",
  "Toronto Blue Jays": "TOR",
  // AL Central
  "Chicago White Sox": "CSW",
  "Cleveland Indians": "CLE",
  "Cleveland Guardians": "CLE",
  "Detroit Tigers": "DET",
  "Kansas City Royals": "KCR",
  "Minnesota Twins": "MIN",
  // AL West
  "Houston Astros": "HOU",
  "Los Angeles Angels": "LAA",
  "Oakland Athletics": "OAK",
  "Seattle Mariners": "SEA",
  "Texas Rangers": "TEX",
  // NL East
  "Atlanta Braves": "ATL",
  "Miami Marlins": "MIA",
  "New York Mets": "NYM",
  "Philadelphia Phillies": "PHI",
  "Washington Nationals": "WAS",
  // NL Central
  "Chicago Cubs": "CHC",
  "Cincinnati Reds": "CIN",
  "Milwaukee Brewers": "MIL",
  "Pittsburgh Pirates": "PIT",
  "St. Louis Cardinals": "STL",
  "St Louis Cardinals": "STL",
  // NL West
  "Arizona Diamondbacks": "ARZ",
  "Colorado Rockies": "COL",
  "Los Angeles Dodgers": "LAD",
  "San Diego Padres": "SD",
  "San Francisco Giants": "SFG",
}

type _AuthCustomerResp struct {
  Code string
}

type _GetLeagueLinesResp struct {
  Lines []struct {
    Team1ID           string
    Team2ID           string
    MoneyLine1        int32
    MoneyLine2        int32
    MoneyLineDecimal1 float64
    MoneyLineDecimal2 float64
    ScheduleDate      string
  }
}

func FetchLines() ([]*model.Line, error) {
  token, err := getAuthToken()
  if err != nil {
    return nil, err
  }

  return getLines2(token)
}

func getAuthToken() (string, error) {
  resp, err := http.PostForm(_URL_AUTH_CUSTOMER, url.Values{
    "customerID":    {"DSI193788"},
    "state":         {"true"},
    "password":      {"GENERIC1"},
    "multiaccount":  {"0"},
    "response_type": {"code"},
    "client_id":     {"DSI193788"},
    "domain":        {"retro-ii.betdsi.eu"},
    "redirect_uri":  {"retro-ii.betdsi.eu"},
    "token":         {"0"},
    "operation":     {"authenticateCustomer"},
    "RRO":           {"1"},
  })

  if err != nil {
    return "", err
  }

  defer resp.Body.Close()
  // TODO check status should be 200
  body, err := io.ReadAll(resp.Body)

  data := _AuthCustomerResp{}
  err = json.Unmarshal(body, &data)
  if err != nil {
    return "", err
  }

  return data.Code, nil
}

func getLines2(token string) ([]*model.Line, error) {
  method := "POST"
  payload := strings.NewReader(fmt.Sprintf(`customerID=DSI193788+&operation=Get_LeagueLines&sportType=BASEBALL&sportSubType=MLB&period=Game&hourFilter=0&propDescription=Game&wagerType=Straight&keyword=&office=DSIMA&correlationID=&periodNumber=0&grouping=&periods=0&rotOrder=0&placeLateFlag=false&RRO=1&access_token=%s`, token))
  client := &http.Client{}
  req, err := http.NewRequest(method, _URL_GET_LEAGUE_LINES, payload)

  if err != nil {
    return nil, err
  }
  req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
  req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  data := _GetLeagueLinesResp{}
  err = json.Unmarshal(body, &data)
  if err != nil {
    return nil, err
  }

  results := make([]*model.Line, 0)
  now := time.Now()
  schduleDateLayout := "2006-01-02"

  for _, line := range data.Lines {
    timeRaw := strings.Split(line.ScheduleDate, " ")[0]
    eventTime, err := time.Parse(schduleDateLayout, timeRaw)
    if err != nil {
      return nil, err
    }
    event := model.Event{
      HomeTeam: teamCodes[line.Team1ID],
      AwayTeam: teamCodes[line.Team2ID],
      Time: eventTime,
    }
    results = append(results, &model.Line{
      Event:           event,
      Source:          _SOURCE,
      TimeCollected:   now,
      LatestCollected: true,
      Type:            model.LINE_TYPE_MONEYLINE,
      Side:            model.SIDE_HOME,
      LineAmerican:    line.MoneyLine1,
      LineDecimal:     line.MoneyLineDecimal1,
    })
    results = append(results, &model.Line{
      Event:           event,
      Source:          _SOURCE,
      TimeCollected:   now,
      LatestCollected: true,
      Type:            model.LINE_TYPE_MONEYLINE,
      Side:            model.SIDE_AWAY,
      LineAmerican:    line.MoneyLine2,
      LineDecimal:     line.MoneyLineDecimal2,
    })
  }

  return results, nil
}
