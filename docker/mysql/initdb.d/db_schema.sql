create table user(
    email varchar(100),
    -- hashed password
    password varchar(60),
    user_name varchar(30),
    photo int,
    primary key (email)
);
create table movie(
    movie_id int,
    primary key(movie_id),
    is_adult bool default false,
    original_title varchar(255) not null,
    kr_title varchar(255),
    poster_path varchar(255),
    release_date date,
    overview text
);
create table channel(
    channel_id int auto_increment,
    primary key(channel_id),
    movie_id int,
    foreign key(movie_id) references movie(movie_id),
    thread_count int default 0,
    subscribe_count int default 0,
    like_count int default 0
);
create table thread(
    thread_id int auto_increment,
    primary key(thread_id),
    channel_id int,
    title varchar(100),
    content text,
    foreign key(channel_id) references channel(channel_id),
    email varchar(100),
    parent int default -1,
    foreign key(email) references user(email),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    is_exposed bool default false
);
create table thread_recommend(
    thread_id int,
    email varchar(100),
    is_recommended bool default false,
    foreign key(thread_id) references thread(thread_id),
    foreign key(email) references user(email),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);
create table genre(
    genre_id int,
    primary key(genre_id),
    genre_name varchar(20)
);
create table genre_relation(
    movie_id int,
    foreign key(movie_id) references movie(movie_id),
    genre_id int,
    foreign key(genre_id) references genre(genre_id)
);
create table user_subscribe(
    email varchar(100),
    foreign key(email) references user(email),
    channel_id int,
    foreign key(channel_id) references channel(channel_id)
);
create table user_scrap(
    email varchar(100),
    foreign key(email) references user(email),
    movie_id int,
    foreign key(movie_id) references movie(movie_id)
);
