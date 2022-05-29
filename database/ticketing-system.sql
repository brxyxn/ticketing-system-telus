BEGIN;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS incidents;
DROP TABLE IF EXISTS history;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS cases;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS agents;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS tiers;
DROP TABLE IF EXISTS statuses;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS companies;
CREATE TABLE "tickets"(
    "ticket_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "case_id" INTEGER NOT NULL,
    "status_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "closed_at" TIMESTAMP(0) WITH TIME zone NOT NULL
);
ALTER TABLE "tickets"
ADD PRIMARY KEY("ticket_id");
CREATE INDEX "tickets_case_id_index" ON "tickets"("case_id");
CREATE INDEX "tickets_status_id_index" ON "tickets"("status_id");
CREATE TABLE "comments"(
    "comment_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "body" TEXT NOT NULL,
    "ticket_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "comments"
ADD PRIMARY KEY("comment_id");
CREATE TABLE "statuses"(
    "status_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "title" VARCHAR(255) NOT NULL,
    "by_default" BOOLEAN NULL
);
ALTER TABLE "statuses"
ADD PRIMARY KEY("status_id");
ALTER TABLE "statuses"
ADD CONSTRAINT "statuses_title_unique" UNIQUE("title");
COMMENT ON COLUMN "statuses"."by_default" IS 'if by_default is true status can''t be deleted, if false it was created by an agent';
CREATE TABLE "history"(
    "event_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "description" VARCHAR(255) NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "made_by" VARCHAR(255) NOT NULL,
    "ticket_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "history"
ADD PRIMARY KEY("event_id");
CREATE INDEX "history_ticket_id_index" ON "history"("ticket_id");
COMMENT ON COLUMN "history"."description" IS 'attach comments made when the event ocurrs';
COMMENT ON COLUMN "history"."status" IS 'use current status when event was created';
COMMENT ON COLUMN "history"."made_by" IS 'change or event created by the user... using user-name';
COMMENT ON COLUMN "history"."ticket_id" IS 'index ticket_id';
CREATE TABLE "cases"(
    "case_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "assigned" BOOLEAN NOT NULL,
    "status" BOOLEAN NOT NULL,
    "agent_id" INTEGER NOT NULL,
    "tier_id" INTEGER NOT NULL
);
ALTER TABLE "cases"
ADD PRIMARY KEY("case_id");
CREATE INDEX "cases_agent_id_index" ON "cases"("agent_id");
CREATE INDEX "cases_tier_id_index" ON "cases"("tier_id");
CREATE TABLE "companies"(
    "company_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR(255) NOT NULL
);
ALTER TABLE "companies"
ADD PRIMARY KEY("company_id");
ALTER TABLE "companies"
ADD CONSTRAINT "companies_name_unique" UNIQUE("name");
CREATE TABLE "customers"(
    "customer_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "company_id" INTEGER NOT NULL
);
ALTER TABLE "customers"
ADD PRIMARY KEY("customer_id");
CREATE INDEX "customers_company_id_index" ON "customers"("company_id");
CREATE TABLE "agents"(
    "agent_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "team_id" INTEGER NOT NULL,
    "team_lead" BOOLEAN NULL
);
ALTER TABLE "agents"
ADD PRIMARY KEY("agent_id");
CREATE INDEX "agents_team_id_index" ON "agents"("team_id");
CREATE TABLE "teams"(
    "team_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR(255) NOT NULL,
    "tier_id" INTEGER NOT NULL
);
ALTER TABLE "teams"
ADD PRIMARY KEY("team_id");
ALTER TABLE "teams"
ADD CONSTRAINT "teams_name_unique" UNIQUE("name");
CREATE INDEX "teams_tier_id_index" ON "teams"("tier_id");
CREATE TABLE "tiers"(
    "tier_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR(255) NOT NULL
);
ALTER TABLE "tiers"
ADD PRIMARY KEY("tier_id");
ALTER TABLE "tiers"
ADD CONSTRAINT "tiers_name_unique" UNIQUE("name");
COMMENT ON COLUMN "tiers"."name" IS 'default tiers: tier 1 and tier 2 *escalation*';
CREATE TABLE "users"(
    "user_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "status" BOOLEAN NULL DEFAULT TRUE,
    "customer_id" INTEGER NULL,
    "agent_id" INTEGER NULL,
    "admin" BOOLEAN NULL DEFAULT FALSE,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "users"
ADD PRIMARY KEY("user_id");
ALTER TABLE "users"
ADD CONSTRAINT "users_email_unique" UNIQUE("email");
CREATE INDEX "users_customer_id_index" ON "users"("customer_id");
CREATE INDEX "users_agent_id_index" ON "users"("agent_id");
CREATE TABLE "profiles"(
    "profile_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX "profiles_user_id_index" ON "profiles"("user_id");
ALTER TABLE "profiles"
ADD PRIMARY KEY("profile_id");
ALTER TABLE "profiles"
ADD CONSTRAINT "profiles_user_id_unique" UNIQUE("user_id");
CREATE TABLE "notifications"(
    "notification_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "receiver_id" INTEGER NOT NULL,
    "sender_id" INTEGER NOT NULL,
    "message_id" INTEGER NOT NULL,
    "ticket_id" INTEGER NOT NULL
);
ALTER TABLE "notifications"
ADD PRIMARY KEY("notification_id");
CREATE INDEX "notifications_receiver_id_index" ON "notifications"("receiver_id");
CREATE INDEX "notifications_sender_id_index" ON "notifications"("sender_id");
CREATE INDEX "notifications_message_id_index" ON "notifications"("message_id");
CREATE INDEX "notifications_ticket_id_index" ON "notifications"("ticket_id");
CREATE TABLE "messages"(
    "message_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "subject" VARCHAR(255) NOT NULL,
    "body" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "messages"
ADD PRIMARY KEY("message_id");
CREATE TABLE "incidents"(
    "incident_id" INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    "user_id" INTEGER NOT NULL,
    "case_id" INTEGER NOT NULL
);
ALTER TABLE "incidents"
ADD PRIMARY KEY("incident_id");
CREATE INDEX "incidents_user_id_index" ON "incidents"("user_id");
CREATE INDEX "incidents_case_id_index" ON "incidents"("case_id");
ALTER TABLE "tickets"
ADD CONSTRAINT "tickets_status_id_foreign" FOREIGN KEY("status_id") REFERENCES "statuses"("status_id");
ALTER TABLE "comments"
ADD CONSTRAINT "comments_ticket_id_foreign" FOREIGN KEY("ticket_id") REFERENCES "tickets"("ticket_id");
ALTER TABLE "comments"
ADD CONSTRAINT "comments_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE "cases"
ADD CONSTRAINT "cases_agent_id_foreign" FOREIGN KEY("agent_id") REFERENCES "agents"("agent_id");
ALTER TABLE "agents"
ADD CONSTRAINT "agents_team_id_foreign" FOREIGN KEY("team_id") REFERENCES "teams"("team_id");
ALTER TABLE "tickets"
ADD CONSTRAINT "tickets_case_id_foreign" FOREIGN KEY("case_id") REFERENCES "cases"("case_id");
ALTER TABLE "incidents"
ADD CONSTRAINT "incidents_case_id_foreign" FOREIGN KEY("case_id") REFERENCES "cases"("case_id");
ALTER TABLE "teams"
ADD CONSTRAINT "teams_tier_id_foreign" FOREIGN KEY("tier_id") REFERENCES "tiers"("tier_id");
ALTER TABLE "notifications"
ADD CONSTRAINT "notifications_sender_id_foreign" FOREIGN KEY("sender_id") REFERENCES "users"("user_id");
ALTER TABLE "incidents"
ADD CONSTRAINT "incidents_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE "notifications"
ADD CONSTRAINT "notifications_receiver_id_foreign" FOREIGN KEY("receiver_id") REFERENCES "users"("user_id");
ALTER TABLE "users"
ADD CONSTRAINT "users_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "customers"("customer_id");
ALTER TABLE "users"
ADD CONSTRAINT "users_agent_id_foreign" FOREIGN KEY("agent_id") REFERENCES "agents"("agent_id");
ALTER TABLE "profiles"
ADD CONSTRAINT "profiles_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE "cases"
ADD CONSTRAINT "cases_tier_id_foreign" FOREIGN KEY("tier_id") REFERENCES "tiers"("tier_id");
ALTER TABLE "notifications"
ADD CONSTRAINT "notifications_message_id_foreign" FOREIGN KEY("message_id") REFERENCES "messages"("message_id");
ALTER TABLE "notifications"
ADD CONSTRAINT "notifications_ticket_id_foreign" FOREIGN KEY("ticket_id") REFERENCES "tickets"("ticket_id");
ALTER TABLE "customers"
ADD CONSTRAINT "customers_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "companies"("company_id");
ALTER TABLE "history"
ADD CONSTRAINT "history_ticket_id_foreign" FOREIGN KEY("ticket_id") REFERENCES "tickets"("ticket_id");
-- TRIGGERS
CREATE OR REPLACE FUNCTION do_updated_at() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';
CREATE TRIGGER set_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE do_updated_at();
CREATE TRIGGER set_updated_at BEFORE
UPDATE ON profiles FOR EACH ROW EXECUTE PROCEDURE do_updated_at();
CREATE TRIGGER set_updated_at BEFORE
UPDATE ON tickets FOR EACH ROW EXECUTE PROCEDURE do_updated_at();
-- closed_at
CREATE OR REPLACE FUNCTION do_closed_at() RETURNS TRIGGER AS $$ BEGIN IF NEW.status_id = 3 THEN NEW.closed_at = now();
END IF;
RETURN NEW;
END;
$$ language 'plpgsql';
CREATE TRIGGER close_ticket BEFORE
UPDATE ON tickets FOR EACH ROW EXECUTE PROCEDURE do_closed_at();
COMMIT;