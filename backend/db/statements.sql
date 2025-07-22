-- FetchFlags
SELECT flag, round, team_id, exploit, status
  FROM flags
  WHERE round <= ?
    AND status = 0
  ORDER BY round ASC;

-- UpdateExpiredFlags
UPDATE flags
  SET status = 3
  WHERE round < ?
    AND status = 0;

-- UpdateFlag
UPDATE flags
  SET status = ?
  WHERE flag = ?;

-- InsertFlag
INSERT INTO flags (flag, exploit, team_id, round)
  VALUES (?, ?, ?, ?)
  ON CONFLICT (flag) DO NOTHING;

-- FetchAllExploitData
SELECT exploit, status, COUNT(*)
  FROM flags
  WHERE round >= ?
  GROUP BY exploit, status;

-- FetchExploitData
SELECT status, COUNT(*)
  FROM flags
  WHERE exploit = ?
    AND round >= ?
  GROUP BY status;

-- FetchAllTimelineData
SELECT round, status, COUNT(*)
  FROM flags
  WHERE round >= ?
  GROUP BY round, status;

-- FetchTimelineData
SELECT round, status, COUNT(*)
  FROM flags
  WHERE round >= ?
    AND exploit = ?
  GROUP BY round, status;

-- FetchAllTeamsData
SELECT team_id, status, COUNT(*)
  FROM flags
  WHERE round >= ?
  GROUP BY team_id, status;

-- FetchTeamsData
SELECT team_id, status, COUNT(*)
  FROM flags
  WHERE round >= ?
    AND exploit = ?
  GROUP BY team_id, status;
