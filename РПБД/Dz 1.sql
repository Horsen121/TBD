1. do $$
BEGIN
  raise notice 'Hello World!';
END;
$$;

2. do $$
BEGIN
  raise notice '%', now()::date;
END;
$$;

3. do $$
DECLARE x INT;
DECLARE y INT;
BEGIN
	x = 2;
	y = 3;
	raise notice '% + % = %', x,y,x+y;
	raise notice '% - % = %', x,y,x-y;
	raise notice '% * % = %', x,y,x*y;
	raise notice '% / % = %', x,y,x/y;
	raise notice '% MOD % = %', x,y,x%y;
END;
$$;

4.1. do $$
DECLARE mark INT;
BEGIN
	mark = 5;
	IF(mark = 2) THEN
		raise notice 'Неуд';
	ELSEIF(mark = 3) THEN
		raise notice 'Удовлетворительно';
	ELSEIF(mark = 4) THEN
		raise notice 'Хорошо';
	ELSEIF(mark = 5) THEN
		raise notice 'Отлично';
	ELSE
		raise notice 'Введённая оценка не верна';
	END IF;
END;
$$;

4.2 do $$
DECLARE mark INT;
BEGIN
	mark = 5;
	CASE
        WHEN mark = 2 THEN
		raise notice 'Неуд';
        WHEN mark = 3 THEN
		raise notice 'Удовлетворительно';
        WHEN mark = 4 THEN
		raise notice 'Хорошо';
        WHEN mark = 5 THEN
		raise notice 'Отлично';
        ELSE
		raise notice 'Введённая оценка не верна';
    END CASE;
END;
$$;

5.1 do $$
BEGIN
	FOR i IN 20..31 LOOP
		raise notice '%^2 = %', i, i*i;
    END LOOP;
END;
$$;

5.2 do $$
DECLARE num INT;
BEGIN
	num = 20;
	WHILE num<31 LOOP
		raise notice '%^2 = %', num, num*num;
		num = num+1;
    END LOOP;
END;
$$;

5.3 do $$
DECLARE num INT;
BEGIN
	num = 20;
	LOOP
		raise notice '%^2 = %', num, num*num;
		num = num+1;
		IF(num=31) THEN
			EXIT;
		END IF;
    END LOOP;
END;
$$;

6. do $$
DECLARE num INT;
BEGIN
	num = 20;
	LOOP
		IF(num%2=0) THEN
			num = num/2;
			raise notice '%', num;
		ELSE
			num = num*3 + 1;
			raise notice '%', num;
		END IF;
		IF(num=1) THEN
			EXIT;
		END IF;
    END LOOP;
END;
$$;

7. CREATE OR REPLACE FUNCTION Luk(cnt int) RETURNS int
AS $$
DECLARE l_2 INT;
DECLARE l_1 INT;
DECLARE tmp INT;
BEGIN
	l_2 = 2;
	l_1 = 1;
	tmp = 0;
	cnt = cnt-2;
	for i IN 1..cnt LOOP
		tmp = l_1 + l_2;
		l_2 = l_1;
		l_1 = tmp;
    END LOOP;
	RETURN tmp;
END
$$ LANGUAGE plpgsql;

8. CREATE OR REPLACE FUNCTION People_year(n_year int) RETURNS int
AS $$
DECLARE cnt INT;
BEGIN
	cnt = 0;
	SELECT COUNT(*) INTO cnt FROM people
	WHERE EXTRACT(YEAR FROM people.birth_date) = n_year;
	RETURN cnt;
END
$$ LANGUAGE plpgsql;

9. CREATE OR REPLACE FUNCTION People_eyes(n_eyes varchar) RETURNS int
AS $$
DECLARE cnt INT;
BEGIN
	cnt = 0;
	SELECT COUNT(*) INTO cnt FROM people
	WHERE people.eyes = n_eyes;
	RETURN cnt;
END
$$ LANGUAGE plpgsql;

10. CREATE OR REPLACE FUNCTION People_yang() RETURNS int
AS $$
DECLARE q_id INT;
BEGIN
	q_id = 0;
	SELECT p.id INTO q_id FROM people AS p
	ORDER BY p.birth_date DESC
	LIMIT 1;
	RETURN q_id;
END
$$ LANGUAGE plpgsql;

11. CREATE OR REPLACE PROCEDURE People_IMT(IN imt int)
LANGUAGE plpgsql
AS $$
DECLARE p people%ROWTYPE;
BEGIN
	FOR p IN SELECT * FROM people
	WHERE people.weight / (people.growth/100 * people.growth/100) > imt
		LOOP
			RAISE NOTICE 'id: %, name: %, surname: %', p.id, p.name, p.surname;
		END LOOP;
END;
$$;

-- Upd 12-15
12. BEGIN;
	CREATE TABLE relations (
  	person_id integer REFERENCES people(id),
  	rel_id integer REFERENCES people(id),
  	rel_type VARCHAR(255) NOT NULL);
END;

13. CREATE OR REPLACE PROCEDURE People_Create(INOUT q_name varchar, q_surname varchar, q_birth_date DATE, q_growth real, q_weight real, q_eyes varchar, q_hair varchar, q_rel_id int, q_rel_type varchar)
LANGUAGE plpgsql
AS $$
DECLARE q_p_id;
BEGIN
	INSERT INTO people (name, surname, birth_date, growth, weight, eyes, hair)
		VALUES(q_name, q_surname, q_birth_date, q_growth, q_weight, q_eyes, q_hair)
		RETURNING p.id INTO q_p_id;
	INSERT INTO relations(person_id, rel_id, rel_type)
		VALUES(q_p_id, q_rel_id, q_rel_type);
END;
$$;

14. BEGIN;
	ALTER TABLE people 
	ADD COLUMN actuality TIMESTAMP;
	COMMIT;
END;

15. CREATE OR REPLACE PROCEDURE People_Actualize(INOUT q_id int, q_growth real, q_weight real)
LANGUAGE plpgsql
AS $$
BEGIN
	UPDATE people
    SET growth = q_growth,
	weight = q_weight,
	actuality = NOW()
    WHERE people.id = q_id;
END;
$$;
