CREATE TABLE if NOT EXISTS files (
    CreatedAt timestamp,
    HashedName VARCHAR(32),
    OriginalName VARCHAR (128),
    ShortUrl VARCHAR (32)
);