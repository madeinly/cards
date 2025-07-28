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
    stock
) VALUES (
    @id, @name_en, @name_es, @sku, @url_image, @set_name, @set_code,
    @mana_value, @colors, @types, @finish, @has_vendor, @language,
    @visibility, @image_path, @image_url, @stock
);

-- name: GetCardStockById :one
SELECT stock
FROM cards
WHERE id = @id AND language = @language AND finish = @finish;

-- name: GetCardHasVendorById :one
SELECT has_vendor
FROM cards
WHERE id = @id;


-- name: UpdateCardStock :exec
UPDATE cards
SET stock = @stock
WHERE id = @id AND language = @language AND finish = @finish;



-- name: CardExists :one
SELECT EXISTS (
    SELECT 1
    FROM cards
    WHERE id = @id AND finish = @finish AND language = @language
);