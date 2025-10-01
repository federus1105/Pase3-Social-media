-- public.postingan definition

-- Drop table

-- DROP TABLE postingan;

CREATE TABLE postingan (
	id serial4 NOT NULL,
	user_id int4 NULL,
	image varchar(255) NULL,
	"content" varchar(255) NULL,
	created_at date DEFAULT CURRENT_DATE NULL,
	CONSTRAINT postingan_pkey PRIMARY KEY (id),
	CONSTRAINT postingan_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(userid)
);