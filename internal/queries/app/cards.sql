-- name: GetCard :one
SELECT *
FROM cards
WHERE id = @id;

-- name: GetPrice :one
SELECT price
FROM cards_price
WHERE card_id = @cardID AND finish = @finish;


-- name: CreateCard :exec
INSERT INTO cards (
    id,
    name_en,
    name_es,
    sku,
    url_image,
    set_name,
    set_code,
    mana_value,
    colors,
    types,
    finish,
    has_vendor,
    language,
    visibility,
    image_path,
    image_url,
    note,
    stock
) VALUES (
    @id, @name_es, @name_en, @sku, @url_image, @set_name, @set_code,
    @mana_value, @colors, @types, @finish, @has_vendor, @language,
    @visibility, @image_path, @image_url, @note, @stock
);

-- name: GetCardStockById :one
SELECT stock
FROM cards
WHERE id = @id;

-- name: GetCardHasVendorById :one
SELECT has_vendor
FROM cards
WHERE id = @id