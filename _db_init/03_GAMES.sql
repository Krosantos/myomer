CREATE TABLE "games" (
  "id" uuid NOT NULL,
  "player_01" uuid NOT NULL,
  "player_02" uuid NOT NULL,
  "outcome" character varying,
  "started_at" timestamp NOT NULL,
  "finished_at" timestamp,
  CONSTRAINT "games_id" PRIMARY KEY ("id"),
  CONSTRAINT "games_player_01_id_fkey" FOREIGN KEY (player_01) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "games_player_02_id_fkey" FOREIGN KEY (player_02) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);
