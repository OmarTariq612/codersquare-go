CREATE TABLE users (
  id        VARCHAR NOT NULL PRIMARY KEY,
  firstname VARCHAR NOT NULL,
  lastname  VARCHAR NOT NULL,
  username  VARCHAR UNIQUE NOT NULL,
  email     VARCHAR UNIQUE NOT NULL,
  password  VARCHAR NOT NULL,
  created_at INTEGER NOT NULL
);

CREATE TABLE posts (
  id       VARCHAR NOT NULL PRIMARY KEY,
  title    VARCHAR NOT NULL,
  url      VARCHAR UNIQUE NOT NULL,
  user_id   VARCHAR NOT NULL,
  posted_at INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE comments (
  id      VARCHAR NOT NULL PRIMARY KEY,
  user_id  VARCHAR NOT NULL,
  post_id  VARCHAR NOT NULL,
  comment VARCHAR NOT NULL,
  posted_at INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE TABLE likes (
  user_id  VARCHAR NOT NULL,
  post_id  VARCHAR NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (post_id) REFERENCES posts (id),
  PRIMARY KEY (user_id, post_id)
);
