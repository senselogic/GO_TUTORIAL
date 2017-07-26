package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditArticles( context * golf.Context ) {
    var session SESSION;
    var article_array [] ARTICLE;
    var section_array [] SECTION;
    var user_array [] USER;

    Model.GetSession( context, &session );
    Model.GetArticleArray( &article_array );
    Model.GetSectionArray( &section_array );
    Model.GetUserArray( &user_array );

    data := map[ string ] interface {} {
        "Session" : session, "ArticleArray" : article_array, "SectionArray" : section_array, "UserArray" : user_array };

    AdministrationView.RenderEditArticles( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditArticle( context * golf.Context ) {
    var session SESSION;
    var article ARTICLE;
    var section_array [] SECTION;
    var user_array [] USER;

    id := GetUint32FromString( GetQueryValue( context, "article_id" ) );

    Model.GetSession( context, &session );
    Model.GetArticleById( &article, id );
    Model.GetSectionArray( &section_array );
    Model.GetUserArray( &user_array );

    data := map[ string ] interface {} {
        "Session" : session, "Article" : article, "SectionArray" : section_array, "UserArray" : user_array };

    AdministrationView.RenderEditArticle( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) ChangeArticle( context * golf.Context ) {
    id := GetUint32FromString( GetQueryValue( context, "article_id" ) );

    context.Request.ParseForm();

    title := GetPostValue( context, "title" );
    text := GetPostValue( context, "text" );
    image := GetPostValue( context, "image" );
    date := GetPostValue( context, "date" );
    section_id := GetUint32FromString( GetPostValue( context, "section_id" ) );
    user_id := GetUint32FromString( GetPostValue( context, "user_id" ) );

    Model.ChangeArticle( id, title, text, image, date, section_id, user_id );

    administration_controller.EditArticles( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) AddArticle( context * golf.Context ) {
    context.Request.ParseForm();

    title := GetPostValue( context, "title" );
    text := GetPostValue( context, "text" );
    image := GetPostValue( context, "image" );
    section_id := GetUint32FromString( GetPostValue( context, "section_id" ) );
    user_id := GetUint32FromString( GetPostValue( context, "user_id" ) );

    Model.AddArticle( title, text, image, section_id, user_id );

    administration_controller.EditArticles( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) RemoveArticle( context * golf.Context ) {
    article_id := GetUint32FromString( GetQueryValue( context, "article_id" ) );

    Model.RemoveArticle( article_id );

    administration_controller.EditArticles( context );
}
