package db

func FetchFlags(round int64) ([]Flag, error) {
	rows, err := FetchFlagsStmt.Query(round)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flags := make([]Flag, 0)
	for rows.Next() {
		var flag Flag
		err := rows.Scan(
			&flag.Flag,
			&flag.Round,
			&flag.TeamID,
			&flag.Exploit,
			&flag.Status,
		)
		if err != nil {
			return nil, err
		}
		flags = append(flags, flag)
	}

	return flags, nil
}

func UpdateExpiredFlags(round int64) (int64, error) {
	res, err := UpdateExpiredFlagsStmt.Exec(round)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateFlags(responses []Response) error {
	for _, resp := range responses {
		_, err := UpdateFlagStmt.Exec(resp.Status, resp.Flag)
		if err != nil {
			return err
		}
	}

	return nil
}
