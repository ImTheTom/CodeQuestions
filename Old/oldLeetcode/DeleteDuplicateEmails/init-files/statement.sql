DELETE FROM Person
	WHERE id NOT IN (SELECT *
    FROM (SELECT MIN(p.id)
            FROM Person p
        GROUP BY p.email) x);
