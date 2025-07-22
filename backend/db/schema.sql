CREATE TABLE IF NOT EXISTS "flags" (
  "flag" CHAR(32) NOT NULL PRIMARY KEY,
  "round" INT NOT NULL,
  "team_id" INT NOT NULL,
  "exploit" TEXT NOT NULL,
  "status" INT DEFAULT 0 CHECK ("status" IN (0, 1, 2, 3)) NOT NULL
);
