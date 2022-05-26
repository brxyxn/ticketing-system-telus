CREATE TABLE IF NOT EXISTS users(
    user_id INT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    username VARCHAR(15) NOT NULL,
    password VARCHAR(75) NOT NULL,
    status BOOLEAN,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);
-- Create a new database called 'profiles'
CREATE TABLE IF NOT EXISTS profiles(
    profile_id INT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    status BOOLEAN,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS tiers(
    tier_id INT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(50) NOT NULL
);
CREATE TABLE IF NOT EXISTS teams(
    team_id INT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    tier_id INT NOT NULL,
    CONSTRAINT fk_tier FOREIGN KEY(tier_id) REFERENCES tiers(tier_id)
);
CREATE TABLE IF NOT EXISTS agents(
    agent_id INT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    team_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(user_id),
    CONSTRAINT fk_team FOREIGN KEY(team_id) REFERENCES teams(team_id)
);
CREATE TABLE IF NOT EXISTS notifications(
    notification_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    email_id INT NOT NULL,
    ticket_id INT NOT NULL,
    type_id INT NOT NULL,
    CONSTRAINT fk_user_sender FOREIGN KEY(sender_id) REFERENCES users(user_id),
    CONSTRAINT fk_user_receiver FOREIGN KEY(receiver_id) REFERENCES users(user_id)
);

COMMIT;