CREATE TABLE "users" (
  "id" VARCHAR(10) PRIMARY KEY UNIQUE,
  "first_name" VARCHAR(50),
  "last_name" VARCHAR(50),
  "age" INT NOT NULL,
  "recording_date" timestamptz
);