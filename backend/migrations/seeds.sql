-- tiers
insert into tiers(name)
values ('Tier 1');
insert into tiers(name)
values ('Tier 2');
-- teams
insert into teams (name, tier_id)
values ('Alpha', 1);
insert into teams (name, tier_id)
values ('Beta', 1);
insert into teams (name, tier_id)
values ('Charlie', 2);
insert into teams (name, tier_id)
values ('Delta', 2);
--------------------------------------------------------------------------------
-- creating customers
-- companies
insert into companies(name)
values ('Patito Company 1');
insert into companies(name)
values ('Patito Company 2');
insert into companies(name)
values ('Patito Company 3');
insert into companies(name)
values ('Patito Company 4');
-- customers
insert into customers(company_id)
values (1);
insert into customers(company_id)
values (2);
insert into customers(company_id)
values (3);
insert into customers(company_id)
values (4);
-- users for customers
-- customers
insert into users(email, password, customer_id)
values ('email@example1.com', 'password', 1);
insert into users(email, password, customer_id)
values ('email@example2.com', 'password', 2);
insert into users(email, password, customer_id)
values ('email@example3.com', 'password', 3);
insert into users(email, password, customer_id)
values ('email@example4.com', 'password', 4);
insert into users(email, password, customer_id)
values ('email@example5.com', 'password', 4);
-- profiles for customers
insert into profiles(first_name, last_name, user_id)
values ('John', 'Doe', 1);
insert into profiles(first_name, last_name, user_id)
values ('John', 'Smith', 2);
insert into profiles(first_name, last_name, user_id)
values ('Carl', 'Doe', 3);
insert into profiles(first_name, last_name, user_id)
values ('Seidy', 'Doe', 4);
insert into profiles(first_name, last_name, user_id)
values ('Gabriella', 'Smith', 5);
--------------------------------------------------------------------------------
-- agents
insert into agents (team_id, team_lead)
values(1, true);
insert into agents (team_id, team_lead)
values(1, false);
insert into agents (team_id, team_lead)
values(1, false);
insert into agents (team_id, team_lead)
values(1, false);
insert into agents (team_id, team_lead)
values(2, true);
insert into agents (team_id, team_lead)
values(2, false);
insert into agents (team_id, team_lead)
values(2, false);
insert into agents (team_id, team_lead)
values(2, false);
--------------------------------------------------------------------------------
-- statuses
insert into statuses (title, by_default)
values ('Created', true);
insert into statuses (title, by_default)
values ('Escalated', true);
insert into statuses (title, by_default)
values ('Closed', true);
insert into statuses (title, by_default)
values ('Combined', true);
insert into statuses (title, by_default)
values ('Merged', true);
