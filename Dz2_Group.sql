1. SELECT n_group, COUNT(*) FROM students GROUP BY n_group
2. SELECT n_group, MAX(score) FROM students GROUP BY n_group
3. SELECT surname, COUNT(*) FROM students GROUP BY surname
4. SELECT score, COUNT(*) FROM students GROUP BY score // for score
5. ??
6. ??
7. SELECT n_group, AVG(score) FROM students GROUP BY n_group ORDER BY AVG(score)
8. SELECT n_group, COUNT(*), MAX(score), MIN(score), AVG(score) FROM students GROUP BY n_group
9. SELECT n_group, st_name FROM students WHERE score=(SELECT MAX(score) FROM students) Group by n_group, st_name
10. ??