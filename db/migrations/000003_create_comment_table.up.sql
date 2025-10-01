-- public."comment" definition

-- Drop table

-- DROP TABLE public."comment";

CREATE TABLE comment (
	id serial4 NOT NULL,
	postingan_id int4 NULL,
	user_id int4 NULL,
	teks varchar(255) NULL,
	created_at date DEFAULT CURRENT_DATE NULL,
	CONSTRAINT comment_pkey PRIMARY KEY (id),
	CONSTRAINT comment_postingan_id_fkey FOREIGN KEY (postingan_id) REFERENCES postingan(id),
	CONSTRAINT comment_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(userid)
);