package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( view * VIEW ) RenderShowSection( context * golf.Context, data * map[ string ] interface {} ) {
    context.Loader( "template" ).Render( "show_section.html", *data );
}
