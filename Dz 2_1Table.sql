	1 Table
1. SELECT name, surname FROM students WHERE (score>=4 AND score<=4.5)
	SELECT * FROM students WHERE (score>=4 AND score<=4.5)
2. SELECT * FROM students GROUP BY n_group=n_group/1000
3. SELECT * FROM students ORDER BY n_group, name
4. SELECT * FROM students WHERE score>=4 ORDER BY score DESC
5. SELECT name, risk FROM hobby WHERE name = 'Футбол' or name = 'Хоккей'
6. SELECT hobby_id, student_id FROM students_hobby WHERE start_d > ('2009-1-1'::DATE) and start_d < ('2018-1-1'::DATE) and finish_d is NULL
7. SELECT * FROM students WHERE score>4.5 ORDER BY score DESC
8. SELECT * FROM students WHERE score>4.5 ORDER BY score DESC LIMIT 5
9. SELECT *, CASE WHEN (risk >=8) THEN 'очень высокий'  WHEN risk >=6 and risk<8 THEN 'высокий'  WHEN risk >=4 and risk<6 THEN 'средний' 
	WHEN risk >=2 and risk<4 THEN 'низкий'  WHEN risk<2 THEN 'очень низкий' END
	FROM hobby
10. SELECT * FROM hobby ORDER BY risk DESC LIMIT 3