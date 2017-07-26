package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );

// == FUNCTIONS

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditArticles(
    context * golf.Context,
    data * map[ string ] interface {}
    )
{
    context.Loader( "template" ).Render( "administration_edit_articles.html", *data );
}

// ~~

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditArticle(
    context * golf.Context,
    data * map[ string ] interface {}
    )
{
    context.Loader( "template" ).Render( "administration_edit_article.html", *data );
}
