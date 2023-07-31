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
    original_title varchar(100) not null,
    kr_title varchar(100),
    poster_path varchar(255),
    release_date date not null,
    overview varchar(255)
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
    content text,
    foreign key(channel_id) references channel(channel_id),
    email varchar(100),
    foreign key(email) references user(email),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);
create table genre(
    movie_id int,
    foreign key(movie_id) references movie(movie_id),
    genre_id int not null
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