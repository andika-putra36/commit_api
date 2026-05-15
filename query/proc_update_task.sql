/*
	CALL update_task(
		1
		, 'Study Postgresql (Revised)'
		, 'Data will be stored in Postgre'
		, '09:00:00'
		, '13:00:00'
		, true
	)
*/

CREATE OR REPLACE PROCEDURE update_task(
	p_id			INT
	, p_title 		VARCHAR(255)
	, p_subtitle 	VARCHAR(255)
	, p_start_at 	TIME
	, p_finish_at 	TIME
	, p_is_done		BOOLEAN
)
AS $$
BEGIN
	UPDATE tasks
	SET
		title 			= p_title
		, subtitle 		= p_subtitle
		, start_at 		= p_start_at
		, finish_at 	= p_finish_at
		, is_done		= p_is_done
		, updated_at 	= NOW()
	WHERE id = p_id
;END;
$$ LANGUAGE plpgsql;