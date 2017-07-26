type SECTION struct {
    Id uint32 `db:"Id"`;
    Number uint32 `db:"Number"`;
    Name string `db:"Name"`;
    Text string `db:"Text"`;
    Image string `db:"Image"`;
    ImageIndex uint32;
}

// ~~

type USER struct {
    Id uint32 `db:"Id"`;
    Email string `db:"Email"`;
    Pseudonym string `db:"Pseudonym"`;
    Password string `db:"Password"`;
    ItIsAdministrator bool `db:"ItIsAdministrator"`;
}

// ~~

type ARTICLE struct {
    Id uint32 `db:"Id"`;
    SectionId uint32 `db:"SectionId"`;
    UserId uint32 `db:"UserId"`;
    Title string `db:"Title"`;
    Text string `db:"Text"`;
    Image string `db:"Image"`;
    Date string `db:"Date"`;
    Section * SECTION;
    User * USER;
    ImageIndex uint32;
}

// ~~

type COMMENT struct {
    Id uint32 `db:"Id"`;
    ArticleId uint32 `db:"ArticleId"`;
    UserId uint32 `db:"UserId"`;
    Text string `db:"Text"`;
    Date string `db:"Date"`;
    Article * ARTICLE;
    User * USER;
}

// ~~

type SUBSCRIBER struct {
    Id uint32 `db:"Id"`;
    Name string `db:"Name"`;
    Email string `db:"Email"`;
}

// ~~

