package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );

// == FUNCTIONS

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditSections(
    context * golf.Context,
    data * map[ string ] interface {}
    )
{
    context.Loader( "template" ).Render( "administration_edit_sections.html", *data );
}

// ~~

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditSection(
    context * golf.Context,
    data * map[ string ] interface {}
    )
{
    context.Loader( "template" ).Render( "administration_edit_section.html", *data );
}
