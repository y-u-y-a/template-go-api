-- create "companies" table
CREATE TABLE "companies" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create "inquiries" table
CREATE TABLE "inquiries" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "email" character varying NOT NULL, "tel" character varying NOT NULL, "content" character varying NOT NULL, "is_confirm" boolean NOT NULL DEFAULT false, PRIMARY KEY ("id"));
-- create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "given_name" character varying NOT NULL, "family_name" character varying NOT NULL, "gender" bigint NOT NULL, "email" character varying NOT NULL, "birthday" date NOT NULL, "company_id" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_companies_company" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE NO ACTION);
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create index "user_given_name_family_name" to table: "users"
CREATE INDEX "user_given_name_family_name" ON "users" ("given_name", "family_name");