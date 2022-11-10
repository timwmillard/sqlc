CREATE TABLE bar (
    ready bool not null,
    name text not null
);

-- -- name: CTEWithAlias :many
-- WITH extra_data AS (
-- 	SELECT
-- 		ready,
-- 		name  AS label
-- 	FROM bar
-- )
-- SELECT *
-- FROM extra_data
-- WHERE label = $1;

-- name: CTEWithUpdate :many
WITH extra_data AS (
    UPDATE bar
    SET ready = true
    WHERE name = $1
    RETURNING *
)
SELECT * 
FROM extra_data
WHERE name = $2;
