package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );

// == FUNCTIONS

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditComments(
    context * golf.Context,
    data * map[ string ] interface {}
    )
{
    context.Loader( "template" ).Render( "administration_edit_comments.html", *data );
}

// ~~

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditComment(
    context * golf.Context,
    data * map[ string ] interface {}
    )
{
    context.Loader( "template" ).Render( "administration_edit_comment.html", *data );
}
