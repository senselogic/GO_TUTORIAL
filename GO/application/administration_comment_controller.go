package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditComments( context * golf.Context ) {
    var session SESSION;
    var comment_array [] COMMENT;
    var article_array [] ARTICLE;
    var user_array [] USER

    Model.GetSession( context, &session );
    Model.GetCommentArray( &comment_array );
    Model.GetArticleArray( &article_array );
    Model.GetUserArray( &user_array );

    data := map[ string ] interface {} {
        "Session" : session, "CommentArray" : comment_array, "ArticleArray" : article_array, "UserArray" : user_array };

    AdministrationView.RenderEditComments( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditComment( context * golf.Context ) {
    var session SESSION;
    var comment COMMENT;
    var article_array [] ARTICLE;
    var user_array [] USER;

    id := GetIntegerFromString( GetQueryValue( context, "comment_id" ) );

    Model.GetSession( context, &session );
    Model.GetCommentById( &comment, id );
    Model.GetArticleArray( &article_array );
    Model.GetUserArray( &user_array );

    data := map[ string ] interface {} {
        "Session" : session, "Comment" : comment, "ArticleArray" : article_array, "UserArray" : user_array };

    AdministrationView.RenderEditComment( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) ChangeComment( context * golf.Context ) {
    id := GetIntegerFromString( GetQueryValue( context, "comment_id" ) );

    context.Request.ParseForm();

    text := GetPostValue( context, "text" );
    date := GetPostValue( context, "date" );
    article_id := GetIntegerFromString( GetPostValue( context, "article_id" ) );
    user_id := GetIntegerFromString( GetPostValue( context, "user_id" ) );

    Model.ChangeComment( id, text, date, article_id, user_id );

    administration_controller.EditComments( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) AddComment( context * golf.Context ) {
    context.Request.ParseForm();

    text := GetPostValue( context, "text" );
    article_id := GetIntegerFromString( GetPostValue( context, "article_id" ) );
    user_id := GetIntegerFromString( GetPostValue( context, "user_id" ) );

    Model.AddComment( text, article_id, user_id );

    administration_controller.EditComments( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) RemoveComment( context * golf.Context ) {
    comment_id := GetIntegerFromString( GetQueryValue( context, "comment_id" ) );

    Model.RemoveComment( comment_id );

    administration_controller.EditComments( context );
}
