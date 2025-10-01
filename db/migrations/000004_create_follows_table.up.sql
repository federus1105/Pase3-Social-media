-- public.follows definition

-- Drop table

-- DROP TABLE public.follows;

CREATE TABLE follows (
	id serial4 NOT NULL,
	follower_id int4 NOT NULL,
	following_id int4 NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	CONSTRAINT follows_pkey PRIMARY KEY (id),
	CONSTRAINT unique_follow UNIQUE (follower_id, following_id),
	CONSTRAINT follows_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES users(userid) ON DELETE CASCADE,
	CONSTRAINT follows_following_id_fkey FOREIGN KEY (following_id) REFERENCES users(userid) ON DELETE CASCADE
);