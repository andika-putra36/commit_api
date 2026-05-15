/*
	CALL delete_task(1)
*/

CREATE OR REPLACE PROCEDURE delete_task(
	p_id INT
)
AS $$
BEGIN
	DELETE FROM tasks
	WHERE id = p_id
;END;
$$ LANGUAGE plpgsql;