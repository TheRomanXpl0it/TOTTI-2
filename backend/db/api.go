package db

import "fmt"

func InsertFlags(received []FlagReceived) error {
	var e error
	for _, r := range received {
		_, err := InsertFlagStmt.Exec(r.Flag, r.Exploit, r.TeamID, r.Round)
		if err != nil {
			if e == nil {
				e = err
			} else {
				e = fmt.Errorf("%v, %v", e, err)
			}
		}
	}
	return nil
}

func FetchAllExploitData(round int) (interface{}, error) {
	rows, err := FetchAllExploitDataStmt.Query(round)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[string]ExploitData)
	for rows.Next() {
		var exploit string
		var status int64
		var count int64
		err := rows.Scan(&exploit, &status, &count)
		if err != nil {
			return nil, err
		}
		if _, ok := data[exploit]; !ok {
			// TODO: make the IsWorking real
			data[exploit] = ExploitData{IsWorking: true, Status: make([]int64, 4)}
		}
		data[exploit].Status[status] = count
	}

	return data, nil
}

func FetchExploitData(round int, exploit string) (interface{}, error) {
	rows, err := FetchExploitDataStmt.Query(exploit, round)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[string]ExploitData)
	// TODO: make the IsWorking real
	data[exploit] = ExploitData{IsWorking: true, Status: make([]int64, 4)}
	for rows.Next() {
		var status int64
		var count int64
		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, err
		}
		data[exploit].Status[status] = count
	}

	return data, nil
}

func FetchAllTimelineData(round int) (interface{}, error) {
	rows, err := FetchAllTimelineDataStmt.Query(round)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[int64]TimelineData)
	for rows.Next() {
		var round int64
		var status int64
		var count int64
		err := rows.Scan(&round, &status, &count)
		if err != nil {
			return nil, err
		}
		if _, ok := data[round]; !ok {
			data[round] = TimelineData{Status: make([]int64, 4)}
		}
		data[round].Status[status] = count
	}

	return data, nil
}

func FetchTimelineData(round int, exploit string) (interface{}, error) {
	rows, err := FetchTimelineDataStmt.Query(round, exploit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[int64]TimelineData)
	for rows.Next() {
		var round int64
		var status int64
		var count int64
		err := rows.Scan(&round, &status, &count)
		if err != nil {
			return nil, err
		}
		if _, ok := data[round]; !ok {
			data[round] = TimelineData{Status: make([]int64, 4)}
		}
		data[round].Status[status] = count
	}

	return data, nil
}

func FetchAllTeamsData(round int) (interface{}, error) {
	rows, err := FetchAllTeamsDataStmt.Query(round)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[int64]TeamsData)
	for rows.Next() {
		var teamID int64
		var status int64
		var count int64
		err := rows.Scan(&teamID, &status, &count)
		if err != nil {
			return nil, err
		}
		if _, ok := data[teamID]; !ok {
			data[teamID] = TeamsData{Status: make([]int64, 4)}
		}
		data[teamID].Status[status] = count
	}

	return data, nil
}

func FetchTeamsData(round int, exploit string) (interface{}, error) {
	rows, err := FetchTeamsDataStmt.Query(round, exploit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[int64]TeamsData)
	for rows.Next() {
		var teamID int64
		var status int64
		var count int64
		err := rows.Scan(&teamID, &status, &count)
		if err != nil {
			return nil, err
		}
		if _, ok := data[teamID]; !ok {
			data[teamID] = TeamsData{Status: make([]int64, 4)}
		}
		data[teamID].Status[status] = count
	}

	return data, nil
}
