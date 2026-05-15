/*
	SELECT * FROM get_task(1)
*/

CREATE OR REPLACE FUNCTION get_task(
	p_id INT
)
RETURNS TABLE (
	id 				INT
	, title 		VARCHAR(255)
	, subtitle 		VARCHAR(255)
	, start_at 		TIME
	, finish_at 	TIME
	, is_done 		BOOLEAN
	, created_at 	TIMESTAMP
	, updated_at 	TIMESTAMP
)
AS $$
	SELECT
	   tasks.id
	   , tasks.title
	   , tasks.subtitle
	   , tasks.start_at
	   , tasks.finish_at
	   , tasks.is_done
	   , tasks.created_at
	   , tasks.updated_at
	FROM tasks
	WHERE 
		id 						= p_id
		AND created_at::DATE 	= CURRENT_DATE
	ORDER BY created_at ASC
	LIMIT 1
;
$$ LANGUAGE sql;