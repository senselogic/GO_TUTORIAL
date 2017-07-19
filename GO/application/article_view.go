package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( view * VIEW ) RenderShowArticle( context * golf.Context, data * map[ string ] interface {} ) {
    context.Loader( "template" ).Render( "show_article.html", *data );
}
