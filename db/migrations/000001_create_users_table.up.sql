-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE users (
	userid serial4 NOT NULL,
	email varchar(50) NULL,
	"password" varchar(255) NULL,
	CONSTRAINT users_pkey PRIMARY KEY (userid)
);