package Application;

// == IMPORTS

import (
    _ "github.com/go-sql-driver/mysql";
    "github.com/jmoiron/sqlx";
    );

// == TYPES

type MODEL struct
{
}

// ~~

type SECTION struct
{
    Id int `db:"id"`;
    Number int `db:"number"`;
    Name string `db:"name"`;
    Text string `db:"text"`;
    Image string `db:"image"`;
    ImageIndex int;
}

// ~~

type USER struct
{
    Id int `db:"id"`;
    Email string `db:"email"`;
    Pseudonym string `db:"pseudonym"`;
    Password string `db:"password"`;
    ItIsAdministrator bool `db:"it_is_administrator"`;
}

// ~~

type ARTICLE struct
{
    Id int `db:"id"`;
    Title string `db:"title"`;
    Text string `db:"text"`;
    Image string `db:"image"`;
    Date string `db:"date"`;
    SectionId int `db:"section_id"`;
    UserId int `db:"user_id"`;
    Section * SECTION;
    User * USER;
    ImageIndex int;
}

// ~~

type COMMENT struct
{
    Id int `db:"id"`;
    Text string `db:"text"`;
    Date string `db:"date"`;
    ArticleId int `db:"article_id"`;
    UserId int `db:"user_id"`;
    Article * ARTICLE;
    User * USER;
}

// ~~

type SUBSCRIBER struct
{
    Id int `db:"id"`;
    Name string `db:"name"`;
    Email string `db:"email"`;
}

// ~~

type SESSION struct
{
    User USER;
    UserIsConnected bool;
    UserHasSubscribed bool;
}

// == VARIABLES

var database * sqlx.DB;
