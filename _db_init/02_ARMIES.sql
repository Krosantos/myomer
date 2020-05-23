CREATE TABLE "public"."armies" (
    "id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "content" json NOT NULL,
    CONSTRAINT "armies_id" PRIMARY KEY ("id"),
    CONSTRAINT "armies_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE
) WITH (oids = false);
