/*
	CALL create_task(
		'Study Go'
		, 'I am working on my first project'
		, '09:00:00'
		, '13:00:00'
	)
*/

CREATE OR REPLACE PROCEDURE create_task(
	p_title 		VARCHAR(255)
	, p_subtitle 	VARCHAR(255)
	, p_start_at 	TIME
	, p_finish_at 	TIME
)
AS $$
BEGIN
	INSERT INTO task (
		title
		, subtitle
		, start_at
		, finish_at
	)
	SELECT
	    p_title
	    , p_subtitle
		, p_start_at
		, p_finish_at;
END;
$$ LANGUAGE plpgsql;