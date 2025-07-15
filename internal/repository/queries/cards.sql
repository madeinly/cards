-- name: GetCard :one
SELECT 
    c.uuid,
    c.name,
    c.setCode,
    c.manaValue,
    c.rarity,
    c.colors,
    c.types,
    s.name AS setName
FROM 
    cardIdentifiers ci
JOIN 
    cards c ON ci.uuid = c.uuid
JOIN 
    sets s ON c.setCode = s.code
WHERE 
    ci.scryfallId = ?;