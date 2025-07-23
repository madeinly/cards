-- name: GetCard :one
SELECT 
    c.uuid,
    c.name,
    c.setCode,
    c.manaValue,
    c.rarity,
    c.colors,
    c.types,
    c.number,
    s.name AS setName
FROM 
    cardIdentifiers ci
JOIN 
    cards c ON ci.uuid = c.uuid
JOIN 
    sets s ON c.setCode = s.code
WHERE 
    ci.scryfallId = ?;

-- name: GetPrice :one
SELECT price
FROM cards_price
WHERE card_id = @cardID;