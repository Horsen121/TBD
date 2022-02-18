CREATE TABLE students(
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	surname varchar(255),
	address varchar(3000),
	score REAL CHECK (score>=2 and score<=5),
	n_group INT CHECK (n_group>=1000 and n_group<=9999)
);

CREATE TABLE hobby(
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	risk INT CHECK (risk>=2 and risk<=10),
);

CREATE TABLE students_hobby(
	student_id integer REFERENCES student(id),
	hobby_id integer REFERENCES hobby(id),
	start_d data NOT NULL,
	finish_d data
);