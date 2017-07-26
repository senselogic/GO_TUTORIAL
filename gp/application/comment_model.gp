package Application;

// == FUNCTIONS

func ( model * MODEL ) InflateComment(
    comment * COMMENT
    )
{
    comment.Article = new( ARTICLE );
    model.GetArticleById( comment.Article, comment.ArticleId );
    
    comment.User = new( USER );
    model.GetUserById( comment.User, comment.UserId );
}

// ~~

func ( model * MODEL ) InflateCommentArray(
    comment_array * [] COMMENT
    )
{
    for comment_index := range *comment_array 
    {
        comment := &( *comment_array )[ comment_index ];
        
        model.InflateComment( comment );
    }
}

// ~~

func ( model * MODEL ) GetCommentArray(
    comment_array * [] COMMENT
    )
{
    Check( database.Select( comment_array, "select * from COMMENT order by Date DESC" ) );
    
    model.InflateCommentArray( comment_array );
}

// ~~

func ( model * MODEL ) GetCommentById(
    comment * COMMENT,
    id uint32
    )
{
    Check( database.Get( comment, "select * from COMMENT where Id = ?", id ) );
    
    model.InflateComment( comment );
}

// ~~

func ( model * MODEL ) GetCommentArrayByArticleId(
    comment_array * [] COMMENT,
    article_id uint32
    )
{
    Check( database.Select( comment_array, "select * from COMMENT where ArticleId = ? order by Date DESC", article_id ) );
    
    model.InflateCommentArray( comment_array );
}

// ~~

func ( model * MODEL ) ChangeComment(
    id uint32,
    text string,
    date string,
    article_id uint32,
    user_id uint32
    )
{
    database.MustExec( 
        "update COMMENT set Id = ?, Text = ?, Date = ?, ArticleId = ?, UserId = ? where Id = ?", 
        id,
        text,
        date,
        article_id,
        user_id,
        id
        );
}

// ~~

func ( model * MODEL ) AddComment(
    text string,
    article_id uint32,
    user_id uint32
    )
{
    database.MustExec( 
        "insert into COMMENT ( Text, Date, ArticleId, UserId ) values ( ?, NOW(), ?, ? )", 
        text, 
        article_id, 
        user_id 
        );
}

// ~~

func ( model * MODEL ) RemoveComment(
    id uint32
    )
{
    database.MustExec( "delete from COMMENT where Id = ?", id );
}
