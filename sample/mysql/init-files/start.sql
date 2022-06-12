-- Create tables --
CREATE TABLE lc.persons (
    id int,
    email varchar(255)
);

-- Insert data --
INSERT INTO
	lc.persons (id, email)
VALUES
	(1, 'john@example.com'),
	(2, 'bob@example.com'),
	(3, 'john@example.com');
