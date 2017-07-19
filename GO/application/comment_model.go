package Application;

// == FUNCTIONS

func ( model * MODEL ) InflateComment( comment * COMMENT ) {
    comment.Article = new( ARTICLE );
    model.GetArticleById( comment.Article, comment.ArticleId );

    comment.User = new( USER );
    model.GetUserById( comment.User, comment.UserId );
}

// ~~

func ( model * MODEL ) InflateCommentArray( comment_array * [] COMMENT ) {
    for comment_index := range *comment_array {
        comment := &( *comment_array )[ comment_index ];

        model.InflateComment( comment );
    }
}

// ~~

func ( model * MODEL ) GetCommentArray( comment_array * [] COMMENT ) {
    Check( database.Select( comment_array, "SELECT * FROM comment ORDER BY date DESC" ) );

    model.InflateCommentArray( comment_array );
}

// ~~

func ( model * MODEL ) GetCommentById( comment * COMMENT, id int ) {
    Check( database.Get( comment, "SELECT * FROM comment WHERE id = ?", id ) );

    model.InflateComment( comment );
}

// ~~

func ( model * MODEL ) GetCommentArrayByArticleId( comment_array * [] COMMENT, article_id int ) {
    Check( database.Select( comment_array, "SELECT * FROM comment WHERE article_id = ? ORDER BY date DESC", article_id ) );

    model.InflateCommentArray( comment_array );
}

// ~~

func ( model * MODEL ) ChangeComment( id int, text string, date string, article_id int, user_id int ) {
    database.MustExec( "UPDATE comment SET id = ?, text = ?, date = ?, article_id = ?, user_id = ? WHERE id = ?", id, text, date, article_id, user_id, id );
}

// ~~

func ( model * MODEL ) AddComment( text string, article_id int, user_id int ) {
    database.MustExec( "INSERT INTO comment ( text, date, article_id, user_id ) VALUES ( ?, NOW(), ?, ? )", text, article_id, user_id );
}

// ~~

func ( model * MODEL ) RemoveComment( id int ) {
    database.MustExec( "DELETE FROM comment WHERE id = ?", id );
}
