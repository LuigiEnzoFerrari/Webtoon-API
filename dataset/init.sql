-- CREATE DATABASE webtoons;
-- \connect webtoons;
CREATE TABLE dataset (
    webtoon_id bigint,
    title text,
    genre text,
    thumbnail text,
    summary text,
    episodes bigint,
    "Created by" text,
    view text,
    subscribe text,
    grade double precision,
    released_date text,
    url text,
    cover text,
    likes text,
    "Written by" text,
    "Art by" text,
    "Adapted by" text,
    "Original work by" text,
    "Assisted by" text
);

COPY dataset (
    webtoon_id,
    title,
    genre,
    thumbnail,
    summary,
    episodes,
    "Created by",
    view,
    subscribe,
    grade,
    released_date,
    url,
    cover,
    likes,
    "Written by",
    "Art by",
    "Adapted by",
    "Original work by",
    "Assisted by"
)
FROM '/tmp/data.csv'
DELIMITER ','
CSV HEADER;

