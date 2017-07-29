package Application;

// == FUNCTIONS

func ( model * MODEL ) InflateArticle(
    article * ARTICLE
    )
{
    article.Section = new( SECTION );
    model.GetSectionById( article.Section, article.SectionId );
        
    article.User = new( USER );
    model.GetUserById( article.User, article.UserId );
}

// ~~

func ( model * MODEL ) InflateArticleArray(
    article_array * [] ARTICLE
    )
{
    for article_index := range *article_array 
    {
        article := &( *article_array )[ article_index ];
        
        model.InflateArticle( article );
    }
}

// ~~

func ( model * MODEL ) GetArticleArray(
    article_array * [] ARTICLE
    )
{
    Check( database.Select( article_array, "select * from ARTICLE order by Date DESC" ) );
    
    model.InflateArticleArray( article_array );
}

// ~~

func ( model * MODEL ) GetArticleById(
    article * ARTICLE,
    id uint32
    )
{
    Check( database.Get( article, "select * from ARTICLE where Id = ?", id ) );
    
    model.InflateArticle( article );
}

// ~~

func ( model * MODEL ) GetArticleArrayBySectionId(
    article_array * [] ARTICLE,
    section_id uint32
    )
{
    Check( database.Select( article_array, "select * from ARTICLE where SectionId = ? order by Date DESC", section_id ) );
    
    model.InflateArticleArray( article_array );
}

// ~~

func ( model * MODEL ) ChangeArticle(
    id uint32,
    title string,
    text string,
    image string,
    date string,
    section_id uint32,
    user_id uint32
    )
{
    database.MustExec( 
        "update ARTICLE set Id = ?, Title = ?, Text = ?, Image = ?, Date = ?, SectionId = ?, UserId = ? where Id = ?", 
        id, 
        title, 
        text, 
        image, 
        date,
        section_id,
        user_id,
        id 
        );
}

// ~~

func ( model * MODEL ) AddArticle(
    title string,
    text string,
    image string,
    section_id uint32,
    user_id uint32
    )
{
    database.MustExec( 
        "insert into ARTICLE ( Title, Text, Image, Date, SectionId, UserId ) values ( ?, ?, ?, NOW(), ?, ? )", 
        title, 
        text, 
        image, 
        section_id,
        user_id,
        );
}

// ~~

func ( model * MODEL ) RemoveArticle(
    id uint32
    )
{
    database.MustExec( "delete from ARTICLE where Id = ?", id );
}
