-- public.account definition

-- Drop table

-- DROP TABLE public.account;

CREATE TABLE account (
	id serial4 NOT NULL,
	username varchar(255) NULL,
	avatar varchar(255) NULL,
	bio varchar(255) NULL,
	user_id int4 NULL,
	CONSTRAINT account_pkey PRIMARY KEY (id),
	CONSTRAINT account_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(userid)
);