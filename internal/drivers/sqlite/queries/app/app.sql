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
    rarity,
    number,
    finish,
    has_vendor,
    language,
    visibility,
    image_path,
    image_url,
    stock
) VALUES (
    @id, @name_en, @name_es, @sku, @url_image, @set_name, @set_code,
    @mana_value, @colors, @types, @rarity, @number, @finish, @has_vendor, @language,
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


-- name: ListAvailableCards :many
SELECT 
    id, name_en, set_code, language, finish, stock
FROM
    cards
WHERE name_en LIKE '%' || @name || '%';

-- name: GetCardsWithPrice :many
SELECT
    c.*,
    p.price
FROM
    cards AS c
JOIN
    cards_price AS p
      ON p.card_id = c.id
     AND p.finish  = c.finish
WHERE
    (@set_code = '' OR c.set_code = @set_code)          
    AND (@name = '' OR c.name_en LIKE '%' || @name || '%')
LIMIT  @limit
OFFSET @offset;

-- name: CountCardsWithPrice :one
SELECT
    COUNT(c.id)
FROM
    cards AS c
JOIN
    cards_price AS p
      ON p.card_id = c.id
     AND p.finish  = c.finish
WHERE
    (@set_code = '' OR c.set_code = @set_code)          
    AND (@name = '' OR c.name_en LIKE '%' || @name || '%');


-- name: GetFilteredCards :many
SELECT
    c.id,
    c.name_en,
    c.language,
    p.price,
    c.url_image
FROM cards AS c
JOIN cards_price AS p
      ON p.card_id = c.id
     AND p.finish  = c.finish
WHERE
    (@cardName = '' OR c.name_en LIKE '%' || @cardName || '%')
    AND (
        (@langEn = 0 AND @langES = 0)
        OR (@langEn = 1 AND c.language = 'English')
        OR (@langES = 1 AND c.language = 'Spanish')
    )
    AND (
        @anyColor = 1
        OR (
            @matchType != 'tight'
            AND (
                (@colorB = 1 AND c.colors LIKE 'B%')
                OR (@colorG = 1 AND c.colors LIKE '%G%')
                OR (@colorR = 1 AND c.colors LIKE '%R%')
                OR (@colorU = 1 AND c.colors LIKE '%U%')
                OR (@colorW = 1 AND c.colors LIKE '%W')
            )
        )
        OR (
            @matchType = 'tight'
            AND (@cardColor = '' OR c.colors = @cardColor)
        )
    )
    AND (@cardType = '' OR c.types = @cardType)
    AND (@cardMv = -1 OR c.mana_value = @cardMv)
    AND (@cardFinish = '' OR c.finish = @cardFinish)
    AND (@cardPriceMin = 0 OR p.price >= @cardPriceMin)
    AND (@cardPriceMax = 0 OR p.price <= @cardPriceMax)
LIMIT @limit
OFFSET @offset;


-- name: CountFilteredCards :one
SELECT
    COUNT( distinct c.id)
FROM cards AS c
JOIN cards_price AS p
      ON p.card_id = c.id
     AND p.finish  = c.finish
WHERE
    (@cardName = '' OR c.name_en LIKE '%' || @cardName || '%')
    AND (
        (@langEn = 0 AND @langES = 0)
        OR (@langEn = 1 AND c.language = 'English')
        OR (@langES = 1 AND c.language = 'Spanish')
    )
    AND (
        @anyColor = 1
        OR (
            @matchType != 'tight'
            AND (
                (@colorB = 1 AND c.colors LIKE 'B%')
                OR (@colorG = 1 AND c.colors LIKE '%G%')
                OR (@colorR = 1 AND c.colors LIKE '%R%')
                OR (@colorU = 1 AND c.colors LIKE '%U%')
                OR (@colorW = 1 AND c.colors LIKE '%W')
            )
        )
        OR (
            @matchType = 'tight'
            AND (@cardColor = '' OR c.colors = @cardColor)
        )
    )
    AND (@cardType = '' OR c.types = @cardType)
    AND (@cardMv = -1 OR c.mana_value = @cardMv)
    AND (@cardFinish = '' OR c.finish = @cardFinish)
    AND (@cardPriceMin = 0 OR p.price >= @cardPriceMin)
    AND (@cardPriceMax = 0 OR p.price <= @cardPriceMax);
