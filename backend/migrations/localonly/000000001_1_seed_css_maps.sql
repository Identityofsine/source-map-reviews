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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM maps;
-- +goose StatementEnd