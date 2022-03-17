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
9. ?
10. ?
11. ?
12. ?
13. ?
14. ?
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
21. ?
22. ?
23. ?
24. ?
25. ?
26. ?
27. ?
28. ?
29. ?
30. ?
31. ?
32. ?
33. ?
34. ?
35. ?
36. ?
37. ?
38. ?
39. ?
40. ?