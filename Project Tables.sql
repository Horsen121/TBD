CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	surname varchar(255) NOT NULL,
	login varchar(25) NOT NULL,
	password varchar(20) NOT NULL,
	status boolean NOT NULL,
	prioritet boolean NOT NULL
);

CREATE TABLE timetable(
	id SERIAL PRIMARY KEY,
	user_id integer REFERENCES user(id),
	started_at timestamptz NOT NULL,
	finished_at timestamptz NOT NULL
);

CREATE TABLE ill(
	id SERIAL PRIMARY KEY,
	user_id integer REFERENCES user(id),
	d_start timestamptz NOT NULL,
	d_finish timestamptz NOT NULL
);

CREATE TABLE change(
	id SERIAL PRIMARY KEY,
	smena_id integer REFERENCES timetable(id),
	started_at timestamptz REFERENCES timetable(started_at),
	finished_at timestamptz REFERENCES timetable(finished_at),
	hoster_id integer REFERENCES user(u.id),
	coef REAL CHECK (coef>=1),
	wonted_start timestamptz,
	wonted_finish timestamptz,
	status boolean
);