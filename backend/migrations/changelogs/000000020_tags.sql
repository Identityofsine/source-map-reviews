-- +goose Up

INSERT INTO lk_tags (lk_tag, description, short_description)
VALUES
('classic', 'Classic maps that come with the base game of Counter-Strike: Source', 'classic'),
('defuse', 'Defuse maps designed for bomb defusal scenarios; usually prefixed by de', 'defuse'),
('hostage', 'Hostage rescue maps designed for rescuing hostages; usually prefixed by cs', 'hostage'),
('aim', 'Aim training maps designed for practicing aiming skills', 'aim'),
('surf', 'Surf maps designed for surfing mechanics and tricks', 'surf'),
('deathmatch', 'Deathmatch maps designed for free-for-all combat scenarios', 'deathmatch'),
('zombie', 'Zombie survival maps where players fight against zombie hordes', 'zombie'),
('jailbreak', 'Jailbreak maps where players can escape from jail or complete objectives', 'jailbreak');
