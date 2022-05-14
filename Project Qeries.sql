1. Регистрация (создание новой строки в таблице users)
	INSERT INTO users (name, surname, login, password, status, prioritet) 
		VALUES ('Name', 'Surname', 'Login', 'Password', '1', 'Prioritet');

2. Авторизация (поиск логина и проверка на корректность введённого пароля и поля status по таблице users)
	SELECT password, status FROM users WHERE login='Login';

3. Выставление приоритетов (изменение поля prioritet в таблице users)
	UPDATE users SET prioritet='Prioritet' WHERE id='Id';

4. Расписание на неделю (выставление сотрудникам времени и дат в таблице timetable)
	INSERT INTO timetable (user_id, started_at, finished_at) 
		VALUES ('user_id', 'started_at', 'finished_at');

5. Просьба о переносе смены (создание новой строки в таблице change)
	INSERT INTO change (smena_id, started_at, finished_at, hoster_id, coef, wonted_start, wonted_finish, status) 
		VALUES ('smena_id', 'started_at', 'finished_at', 'hoster_id', 'coef','wonted_start','wonted_finish', '0');

6. Просмотр расписания (вывод смен пользователя из таблицы timetable)
	SELECT started_at, finished_at FROM timetable WHERE user_id='Id';

7. Просмотр всего расписания (с пагинацией)
    7.1 С сортировкой по дате (сортировка ORDER BY data)
	SELECT user_id, started_at, finished_at FROM timetable
		ORDER BY started_at DESC
		LIMIT 20;
    7.2 С происком по фамилии (вывод смен сотрудника) ?? (п.6)

8. Блокировка пользователя (изменение поля status в таблице users)
	UPDATE users SET status=0 WHERE id='Id';

9. Не выходит по состоянию здоровья (транзакция) ???
    1) создание новой строки в таблице illnes (пользователь)
    2) создание новой/ых строки/ок в таблице change (админ)
    3) отмена смен/ы сотрудника в таблице timetable
	BEGIN;
	INSERT INTO illnes (user_id, started_at, finished_at) 
		VALUES ('user_id', 'started_at', 'finished_at');
	SELECT started_at, finished_at FROM timetable WHERE user_id='Id';
	END;

	BEGIN;
	INSERT INTO change (smena_id, started_at, finished_at, hoster_id, coef, status) 
		VALUES ('smena_id', 'started_at', 'finished_at', 'hoster_id', 'coef', '0');
	DELETE FROM timetable WHERE user_id='User_id';
	END;

10. Подтверждение заявки на замену (транзакция)
    1) изменение поля status в таблице change
    2.1) добавление новой смены в таблице timetable (если замена по болезни)
	BEGIN;
	UPDATE change SET status=1 WHERE smena_id='Id';
	INSERT INTO timetable (user_id, started_at, finished_at) 
		VALUES ('user_id', 'started_at', 'finished_at');
	END;
    2.2) обмен сменами в таблице timetable (если обмен)
	BEGIN;
	UPDATE change SET status=1 WHERE smena_id='Id';
	UPDATE timetable SET started_at='Wonted_start' AND finished_at="Wonted_finish' WHERE smena_id='Smena_id';
	UPDATE timetable SET started_at='Started_at' AND finished_at="Finished_at' WHERE user_id='User_id';
	END;