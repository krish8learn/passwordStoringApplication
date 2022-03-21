CREATE TABLE "emails" (
  "email_id" varchar PRIMARY KEY,
  "domain_name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "reason" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "browsers" (
  "browser_name" varchar PRIMARY KEY,
  "account_created" boolean,
  "email_id_used" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "applications" (
  "app_id" SERIAL PRIMARY KEY,
  "app_name" varchar NOT NULL,
  "reason" varchar,
  "email_id_used" varchar NOT NULL,
  "browser_stored" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "browsers" ADD FOREIGN KEY ("email_id_used") REFERENCES "emails" ("email_id");

ALTER TABLE "applications" ADD FOREIGN KEY ("email_id_used") REFERENCES "emails" ("email_id");

ALTER TABLE "applications" ADD FOREIGN KEY ("browser_stored") REFERENCES "browsers" ("browser_name");

CREATE INDEX ON "emails" ("email_id");

CREATE INDEX ON "emails" ("domain_name");

CREATE INDEX ON "browsers" ("browser_name");

CREATE INDEX ON "applications" ("app_name");

CREATE INDEX ON "applications" ("email_id_used", "browser_stored");
