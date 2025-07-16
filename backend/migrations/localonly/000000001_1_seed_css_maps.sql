-- Seed script for Counter-Strike Source maps
-- Populates the maps table with official and popular community CS:S maps

-- Official Defusal Maps
-- +goose Up
-- +goose StatementBegin
INSERT INTO maps (map_name, map_path) VALUES
('de_dust2', 'de_dust2.bsp'),
('de_inferno', 'de_inferno.bsp'),
('de_nuke', 'de_nuke.bsp'),
('de_train', 'de_train.bsp'),
('de_aztec', 'de_aztec.bsp'),
('de_cbble', 'de_cbble.bsp'),
('de_chateau', 'de_chateau.bsp'),
('de_dust', 'de_dust.bsp'),
('de_piranesi', 'de_piranesi.bsp'),
('de_port', 'de_port.bsp'),
('de_prodigy', 'de_prodigy.bsp'),
('de_tides', 'de_tides.bsp'),

-- Official Hostage Maps  
('cs_compound', 'cs_compound.bsp'),
('cs_havana', 'cs_havana.bsp'),
('cs_italy', 'cs_italy.bsp'),
('cs_militia', 'cs_militia.bsp'),
('cs_office', 'cs_office.bsp');

INSERT INTO map_tags (map_name, lk_tag) 
VALUES
('de_dust2', 'classic'),
('de_inferno', 'classic'),
('de_nuke', 'classic'),
('de_train', 'classic'),
('de_aztec', 'classic'),
('de_cbble', 'classic'),
('de_chateau', 'classic'),
('de_dust', 'classic'),
('de_piranesi', 'classic'),
('de_port', 'classic'),
('de_prodigy', 'classic'),
('de_tides', 'classic'),
('cs_compound', 'classic'),
('cs_havana', 'classic'),
('cs_italy', 'classic'),
('cs_militia', 'classic'),
('cs_office', 'classic'),

('de_dust2', 'defuse'),
('de_inferno', 'defuse'),
('de_nuke', 'defuse'),
('de_train', 'defuse'),
('de_aztec', 'defuse'),
('de_cbble', 'defuse'),
('de_chateau', 'defuse'),
('de_dust', 'defuse'),
('de_piranesi', 'defuse'),
('de_port', 'defuse'),
('de_prodigy', 'defuse'),
('de_tides', 'defuse'),

('cs_compound', 'hostage'),
('cs_havana', 'hostage'),
('cs_italy', 'hostage'),
('cs_militia', 'hostage'),
('cs_office', 'hostage');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM maps;
-- +goose StatementEnd
