package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditUsers( context * golf.Context, data * map[ string ] interface {} ) {
    context.Loader( "template" ).Render( "administration_edit_users.html", *data );
}

// ~~

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditUser( context * golf.Context, data * map[ string ] interface {} ) {
    context.Loader( "template" ).Render( "administration_edit_user.html", *data );
}
