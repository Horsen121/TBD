1. SELECT st_name, surname, name FROM students AS s, hobby AS h, students_hobby AS sh 
	WHERE s.id=sh.student_id AND h.id=sh.hobby_id 
2. SELECT st_name, surname, score, n_group FROM students AS s, students_hobby AS sh 
	WHERE s.id = (SELECT DISTINCT student_id FROM students_hobby WHERE (finish_d - start_d)=
		(SELECT MAX(finish_d - start_d) FROM students_hobby)) AND sh.student_id=s.id
3. SELECT DISTINCT st_name, surname, score, n_group  FROM  students AS s, hobby AS h, students_hobby AS sh
	WHERE score>(SELECT AVG(score) FROM students) AND 
		(SELECT SUM(h.risk) FROM students AS s, students_hobby AS sh, hobby AS h WHERE sh.student_id=s.id)>9
4. SELECT DISTINCT st_name, surname, h.name, (sh.finish_d - sh.start_d)/30 FROM students AS s, students_hobby AS sh, hobby AS h
 	WHERE finish_d IS NOT NULL AND sh.student_id=s.id AND sh.hobby_id=h.id
5. SELECT st_name, surname, score, n_group FROM students AS s, students_hobby AS sh
 	WHERE finish_d IS NOT NULL AND sh.student_id=s.id AND score>(SELECT AVG(score) FROM students)
6. SELECT n_group, AVG(score) AS avg_score FROM students AS st INNER JOIN students_hobby AS sh ON sh.student_id=st.id GROUP BY n_group
7. SELECT st.id AS student, h.name, h.risk, (CURRENT_DATE - sh.start_d)/30 AS duration FROM students AS st
	RIGHT JOIN students_hobby AS sh ON st.id=sh.student_id
	LEFT JOIN hobby AS h ON h.id=sh.hobby_id
	WHERE sh.finish_d IS NULL
	ORDER BY st.id
8. SELECT st.id AS student, h.name FROM students AS st
	RIGHT JOIN students_hobby AS sh ON st.id=sh.student_id
	LEFT JOIN hobby AS h ON h.id=sh.hobby_id
	WHERE score=(SELECT MAX(score) FROM students)
	ORDER BY st.id
9. SELECT h.name FROM hobby AS h INNER JOIN students_hobby AS sh ON h.id=sh.hobby_id
	INNER JOIN students AS s ON s.id=sh.student_id
	WHERE s.n_group/1000=2 AND s.score BETWEEN 3 AND 4
	GROUP BY h.name
10. SELECT st.n_group / 1000 AS course FROM students AS st
	INNER JOIN (SELECT st.id, st.n_group / 1000 course, COUNT(*) cnt FROM students AS st
		RIGHT JOIN students_hobby sh ON sh.student_id = st.id
		GROUP BY st.id, st.n_group
		HAVING COUNT(*) > 1) tb ON st.n_group / 1000 = tb.course
	GROUP BY st.n_group / 1000, tb.cnt
	HAVING tb.cnt::real / COUNT(*)::real > 0.5;
11. SELECT st.n_group, COUNT(*), tb.fourres FROM students AS st
	INNER JOIN( SELECT n_group, COUNT(*) fourres FROM students
    		WHERE score >= 4
    		GROUP BY n_group) tb ON tb.n_group = st.n_group
	GROUP BY st.n_group, tb.fourres
	HAVING tb.fourres::real/COUNT(*)::real >= 0.6;
12. SELECT st.n_group / 1000 AS course, COUNT(DISTINCT sh.hobby_id) FROM students AS st
	LEFT JOIN students_hobby AS sh ON sh.student_id = st.id
	GROUP BY st.n_group / 1000;
13. SELECT st.id, st.surname, st.name, st.n_group / 1000 AS course FROM students AS st
	LEFT JOIN students_hobby AS sh ON sh.student_id = st.id
	WHERE sh.id IS NULL AND st.score >= 4.5
	ORDER BY st.n_group / 1000, st.age DESC;
14. CREATE OR REPLACE VIEW st_hobby_5 AS
	SELECT * FROM students AS s INNER JOIN students_hobby AS sh ON s.id=sh.student_id
	WHERE sh.finish_d IS NULL AND (CURRENT_DATE - sh.start_d)/30>5
15. SELECT h.name, COUNT(*) FROM hobby AS h, students_hobby AS sh
	WHERE h.id=sh.hobby_id
	GROUP BY h.name
16. SELECT h.name FROM hobby AS h, students_hobby AS sh
	WHERE h.id=sh.hobby_id
	GROUP BY h.name
	ORDER BY COUNT(*) DESC
	LIMIT 1
17. SELECT * FROM students AS st, students_hobby AS sh
	WHERE st.id=sh.student_id AND sh.hobby_id=(SELECT h.id FROM hobby AS h, students_hobby AS sh
		WHERE h.id=sh.hobby_id
		GROUP BY h.id
		ORDER BY COUNT(*) DESC
		LIMIT 1)
18. SELECT id FROM hobby
	ORDER BY risk DESC
	LIMIT 3
19. SELECT st.id AS student, h.name, (CURRENT_DATE - sh.start_d)/30 AS duration FROM students AS st
	RIGHT JOIN students_hobby AS sh ON st.id=sh.student_id
	LEFT JOIN hobby AS h ON h.id=sh.hobby_id
	ORDER BY (CURRENT_DATE - sh.start_d)/30 DESC
	LIMIT 10
20. SELECT n_group FROM students WHERE id IN (SELECT DISTINCT st.id AS student FROM students AS st
	RIGHT JOIN students_hobby AS sh ON st.id=sh.student_id
	LEFT JOIN hobby AS h ON h.id=sh.hobby_id
	LIMIT 10)
	GROUP BY n_group
21. CREATE OR REPLACE VIEW task_21 AS
	SELECT st_name, surname, n_group FROM students
	ORDER BY score DESC
22. CREATE OR REPLACE VIEW task_22 AS
SELECT DISTINCT(n_group/1000) as "Cource", h.name FROM hobby as h INNER JOIN students_hobby as sh ON h.id=sh.hobby_id INNER JOIN students as st ON st.id=sh.student_id
WHERE h.id = (SELECT sh.hobby_id FROM students as st INNER JOIN students_hobby as sh ON st.id=sh.student_id
				WHERE st.n_group/1000=1
			    GROUP BY sh.hobby_id
			  	ORDER BY COUNT(*) DESC
			  	LIMIT 1) AND st.n_group/1000=1
		OR  h.id = (SELECT sh.hobby_id FROM students as st INNER JOIN students_hobby as sh ON st.id=sh.student_id
				WHERE st.n_group/1000=2
			    GROUP BY sh.hobby_id
			  	ORDER BY COUNT(*) DESC
			  	LIMIT 1) AND st.n_group/1000=2
		OR  h.id = (SELECT sh.hobby_id FROM students as st INNER JOIN students_hobby as sh ON st.id=sh.student_id
				WHERE st.n_group/1000=3
			    GROUP BY sh.hobby_id
			  	ORDER BY COUNT(*) DESC
			  	LIMIT 1) AND st.n_group/1000=3
		OR  h.id = (SELECT sh.hobby_id FROM students as st INNER JOIN students_hobby as sh ON st.id=sh.student_id
				WHERE st.n_group/1000=3
			    GROUP BY sh.hobby_id
			  	ORDER BY COUNT(*) DESC
			  	LIMIT 1) AND st.n_group/1000=4
23. CREATE OR REPLACE VIEW task_23 AS
	SELECT h.NAME, COUNT(*) AS popularity, h.risk FROM students AS st
	RIGHT JOIN students_hobby AS sh ON sh.student_id = st.id
	LEFT JOIN hobby h ON h.id = sh.hobby_id
	WHERE st.n_group / 1000 = 2
	GROUP BY h.NAME, h.risk
	ORDER BY COUNT(*) DESC, h.risk DESC
	LIMIT 1;
24. CREATE OR REPLACE VIEW task_24 AS
	SELECT st.n_group / 1000 course, COUNT(*), 
	CASE
    	WHEN tb.MostPop IS NULL THEN 0
    	ELSE tb.MostPop
	END MostPop
	FROM students AS st
	LEFT JOIN (SELECT st.n_group, COUNT(*) MostPop
    		FROM students As st
    		WHERE st.score >= 4.5
    		GROUP BY st.n_group) tb ON tb.n_group = st.n_group
	GROUP BY st.n_group / 1000, tb.MostPop
25. SELECT h.name, COUNT(*) FROM students AS st
	LEFT JOIN students_hobby sh ON st.id = sh.student_id
	LEFT JOIN hobby h ON h.id = sh.hobby_id
	WHERE h.name IS NOT NULL
	GROUP BY h.name
	ORDER BY COUNT(*)
	DESC 
	LIMIT 1;
26. CREATE OR REPLACE VIEW task_26 AS
	SELECT * FROM students AS st
	WITH CHECK OPTION;
27. SELECT LEFT(st.name, 1), MAX(st.score), AVG(st.score), MIN(st.score) FROM students AS st
	GROUP BY LEFT(st.name, 1)
	HAVING MAX(st.score) > 3.6
28. SELECT st.n_group / 1000 AS course, st.surname,  MAX(st.score), MIN(score) FROM students AS st
	GROUP BY st.n_group / 1000, st.surname
29. SELECT st.n_group/1000, COUNT(*) FROM students AS st
	LEFT JOIN students_hobby sh ON sh.student_id = st.id
	WHERE sh.hobby_id IS NOT NULL
	GROUP BY st.n_group/1000;
30. SELECT regexp_split_to_table(st.name,''), MIN(h.risk), MAX(h.risk) FROM students AS st
	RIGHT JOIN students_hobby sh ON sh.student_id = st.id
	LEFT JOIN hobby h ON h.id = sh.hobby_id
	GROUP BY regexp_split_to_table(st.name,'');
31. SELECT EXTRACT(MONTH FROM st.birth_at), AVG(st.score) FROM students AS st
	RIGHT JOIN students_hobby sh ON sh.student_id = st.id
	LEFT JOIN hobby h ON sh.hobby_id = h.id
	WHERE h.name LIKE 'Футбол'
	GROUP BY EXTRACT(MONTH FROM st.birth_at);
32. SELECT st.name AS 'Имя', st.surname As 'Фамилия', st.n_group AS 'Группа' FROM students st
	RIGHT JOIN students_hobby sh ON sh.student_id = st.id
	LEFT JOIN hobby h ON sh.hobby_id = h.id
	GROUP BY st.id;
33. SELECT CASE
	WHEN strpos(st.surname, 'ов') != 0 THEN strpos(st.surname, 'ов')::VARCHAR(255)
	ELSE 'Не найдено'
	END
	FROM students AS st;
34. SELECT rpad(st.surname, 10, '#') FROM students AS st;
35. SELECT replace(st.surname, '#', '') FROM students AS st;
36. SELECT  DATE_PART('days', DATE_TRUNC('month', '2018-04-01'::DATE) + '1 MONTH'::INTERVAL - '1 DAY'::INTERVAL);
37. SELECT date_trunc('week', current_date)::date + 5;
38. SELECT EXTRACT(CENTURY FROM current_date), EXTRACT(WEEK FROM current_date), EXTRACT(DOY FROM current_date);
39. SELECT st.name, st.surname, h.name, CASE
	WHEN sh.finished_at IS NULL THEN 'Занимается'
	ELSE 'Закончил'
	END 
	FROM students AS st
	RIGHT JOIN student_hobby sh ON sh.student_id = st.id
	LEFT JOIN hobby h ON h.id = sh.hobby_id;
40. ?