-- Create tables --
CREATE TABLE lc.Employees (
    emp_id int,
    event_day date,
    in_time int,
    out_time int
);

-- Insert data --
INSERT INTO
	lc.Employees (emp_id, event_day, in_time, out_time)
VALUES
	(1, '2020-11-28', 4, 32),
	(1, '2020-11-28', 55, 200),
	(1, '2020-12-03', 1, 42),
	(2, '2020-11-28', 3, 33),
	(2, '2020-12-09', 47, 74);
