-- public.likes definition

-- Drop table

-- DROP TABLE likes;

CREATE TABLE likes (
	id serial4 NOT NULL,
	user_id int4 NOT NULL,
	post_id int4 NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	CONSTRAINT likes_pkey PRIMARY KEY (id),
	CONSTRAINT unique_user_post UNIQUE (user_id, post_id),
	CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES postingan(id) ON DELETE CASCADE,
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(userid) ON DELETE CASCADE
);