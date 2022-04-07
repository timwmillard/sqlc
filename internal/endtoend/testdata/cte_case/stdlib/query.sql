CREATE TABLE bar (ready bool not null);

-- name: CTECase :many
WITH extra_data AS (
	SELECT
		ready,
		CASE WHEN ready THEN
			'Ready'
		ELSE 'Not Ready'
		END AS label
	FROM bar
)
SELECT extra_data.ready
FROM extra_data
WHERE label = $1;
