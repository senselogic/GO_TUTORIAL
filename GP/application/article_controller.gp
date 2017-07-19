package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );
    
// == FUNCTIONS

func ( controller * CONTROLLER ) ShowArticle(
    context * golf.Context
    )
{
    var session SESSION;
    var section_array [] SECTION;
    var article ARTICLE;
    var section SECTION;
    var comment_array [] COMMENT;

    article_id := GetIntegerFromString( GetQueryValue( context, "article_id" ) );

    Model.GetSession( context, &session );
    
    Model.GetSectionArray( &section_array );
    Model.GetArticleById( &article, article_id );
    Model.GetSectionById( &section, article.SectionId );
    Model.GetCommentArrayByArticleId( &comment_array, article_id );
    
    article.ImageIndex = article.Id % 20;

    data := map[ string ] interface {}
    {
        "Session" : session,
        "SectionArray" : section_array,
        "Article" : article,
        "Section" : section,
        "CommentArray" : comment_array
    };

    View.RenderShowArticle( context, &data );
}
