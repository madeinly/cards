CREATE TABLE IF NOT EXISTS cards (
    id          TEXT NOT NULL, -- este id es mtgjson id 
    name_en     TEXT NOT NULL,
    name_es     TEXT NOT NULL,
    sku         TEXT NOT NULL UNIQUE,
    set_name    TEXT NOT NULL,
    set_code    TEXT NOT NULL,
    mana_value  INTEGER NOT NULL,
    colors      TEXT NOT NULL,
    types       TEXT NOT NULL,
    rarity      TEXT NOT NULL,
    number      TEXT NOT NULL,
    finish      TEXT NOT NULL CHECK(finish IN ('foil', 'normal', 'etched')),
    has_vendor  BOOLEAN NOT NULL DEFAULT 0 CHECK(has_vendor IN (0, 1)),
    language    TEXT NOT NULL CHECK(language IN ('Spanish', 'English')),
    visibility  INTEGER NOT NULL CHECK(visibility IN (0, 1)),
    image_path  TEXT, -- Consider NOT NULL if required
    image_url   TEXT NOT NULL,  -- Consider NOT NULL if required
    stock       INTEGER NOT NULL DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id, language, finish)
);

CREATE TABLE IF NOT EXISTS cards_price (
    card_id     TEXT NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    finish      TEXT NOT NULL CHECK(finish IN ('foil', 'etched', 'normal')),
    type        TEXT NOT NULL CHECK(type IN ('buylist', 'retail')),
    price       REAL NOT NULL,
    updated_at  TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (card_id, finish, type)
);

CREATE TABLE IF NOT EXISTS cards_vendor (
    card_id     TEXT NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    vendor_id   TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    qty         INTEGER NOT NULL,
    created_at  TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (card_id, vendor_id)
);

