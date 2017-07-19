package Application;

// == FUNCTIONS

func ( model * MODEL ) InflateArticle( article * ARTICLE ) {
    article.Section = new( SECTION );
    model.GetSectionById( article.Section, article.SectionId );

    article.User = new( USER );
    model.GetUserById( article.User, article.UserId );
}

// ~~

func ( model * MODEL ) InflateArticleArray( article_array * [] ARTICLE ) {
    for article_index := range *article_array {
        article := &( *article_array )[ article_index ];

        model.InflateArticle( article );
    }
}

// ~~

func ( model * MODEL ) GetArticleArray( article_array * [] ARTICLE ) {
    Check( database.Select( article_array, "SELECT * FROM article ORDER BY date DESC" ) );

    model.InflateArticleArray( article_array );
}

// ~~

func ( model * MODEL ) GetArticleById( article * ARTICLE, id int ) {
    Check( database.Get( article, "SELECT * FROM article WHERE id = ?", id ) );

    model.InflateArticle( article );
}

// ~~

func ( model * MODEL ) GetArticleArrayBySectionId( article_array * [] ARTICLE, section_id int ) {
    Check( database.Select( article_array, "SELECT * FROM article WHERE section_id = ? ORDER BY date DESC", section_id ) );

    model.InflateArticleArray( article_array );
}

// ~~

func ( model * MODEL ) ChangeArticle( id int, title string, text string, image string, date string, section_id int, user_id int ) {
    database.MustExec( "UPDATE article SET id = ?, title = ?, text = ?, image = ?, date = ?, section_id = ?, user_id = ? WHERE id = ?", id, title, text, image, date, section_id, user_id, id );
}

// ~~

func ( model * MODEL ) AddArticle( title string, text string, image string, section_id int, user_id int ) {
    database.MustExec( "INSERT INTO article ( title, text, image, date, section_id, user_id ) VALUES ( ?, ?, ?, NOW(), ?, ? )", title, text, image, section_id, user_id, );
}

// ~~

func ( model * MODEL ) RemoveArticle( id int ) {
    database.MustExec( "DELETE FROM article WHERE id = ?", id );
}
