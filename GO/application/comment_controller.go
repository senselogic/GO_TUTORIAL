package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( controller * CONTROLLER ) AddComment( context * golf.Context ) {
    var session SESSION;
    var article ARTICLE;

    article_id := GetIntegerFromString( GetQueryValue( context, "article_id" ) );

    Model.GetSession( context, &session );
    Model.GetArticleById( &article, article_id );

    context.Request.ParseForm();

    text := GetPostValue( context, "text" );

    Model.AddComment( text, article.Id, session.User.Id );

    controller.ShowArticle( context );
}
