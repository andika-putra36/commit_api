

/*
BEGIN;

CALL create_task(
	'Study Postgresql'
	, 'Data will be stored in Postgre'
	, '09:00:00'
	, '13:00:00'
);

CALL create_task(
	'Study Go'
	, 'Data will be transmitted through Go'
	, '13:30:00'
	, '16:00:00'
);

CALL create_task(
	'Study Flutter'
	, 'Data will be shown in Flutter (Mobile)'
	, '16:30:00'
	, '19:00:00'
);

CALL update_task(
	1
	, 'Study Postgresql (Revised)'
	, 'Data will be stored in Postgre'
	, '09:00:00'
	, '13:00:00'
	, true
);

SELECT * FROM get_tasks();

CALL delete_task(1);

SELECT * FROM get_task(4);

ROLLBACK;
*/




