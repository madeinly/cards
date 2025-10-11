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
    cardidentifiers ci
JOIN 
    cards c ON ci.uuid = c.uuid
JOIN 
    sets s ON c.setCode = s.code
WHERE 
    ci.scryfallId = ?;


-- name: GetCardNameES :one
SELECT name
FROM cardForeignData
WHERE uuid = @id AND language = "Spanish";


-- name: GetSetName :one
SELECT name
FROM sets
WHERE code = @setCode;


-- name: GetSets :many
SELECT code, name
FROM sets
WHERE isOnlineOnly = 0;

-- name: ListAllNames :many
SELECT DISTINCT name
FROM cards
WHERE name LIKE '%'||@cardName||'%';

