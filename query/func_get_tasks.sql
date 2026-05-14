/*
	SELECT * FROM get_tasks()
*/

CREATE OR REPLACE FUNCTION get_tasks()
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
	   task.id
	   , task.title
	   , task.subtitle
	   , task.start_at
	   , task.finish_at
	   , task.is_done
	   , task.created_at
	   , task.updated_at
	FROM task
	WHERE created_at::DATE = CURRENT_DATE
	ORDER BY created_at ASC;
$$ LANGUAGE sql;