package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );

// == FUNCTIONS

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditSections(
    context * golf.Context
    )
{
    var session SESSION;
    var section_array [] SECTION;

    Model.GetSession( context, &session );
    Model.GetSectionArray( &section_array );

    data := map[ string ] interface {}
    {
        "Session" : session,
        "SectionArray" : section_array
    };

    AdministrationView.RenderEditSections( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditSection(
    context * golf.Context
    )
{
    var session SESSION;
    var section SECTION;

    id := GetUint32FromString( GetQueryValue( context, "section_id" ) );

    Model.GetSession( context, &session );
    Model.GetSectionById( &section, id );

    data := map[ string ] interface {}
    {
        "Session" : session,
        "Section" : section
    };

    AdministrationView.RenderEditSection( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) ChangeSection(
    context * golf.Context
    )
{
    id := GetUint32FromString( GetQueryValue( context, "section_id" ) );
    
    context.Request.ParseForm();

    number := GetUint32FromString( GetPostValue( context, "number" ) );
    name := GetPostValue( context, "name" );
    text := GetPostValue( context, "text" );
    image := GetPostValue( context, "image" );

    Model.ChangeSection( id, number, name, text, image );

    administration_controller.EditSections( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) AddSection(
    context * golf.Context
    )
{
    context.Request.ParseForm();

    number := GetUint32FromString( GetPostValue( context, "number" ) );
    name := GetPostValue( context, "name" );
    text := GetPostValue( context, "text" );
    image := GetPostValue( context, "image" );

    Model.AddSection( number, name, text, image );

    administration_controller.EditSections( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) RemoveSection(
    context * golf.Context
    )
{
    section_id := GetUint32FromString( GetQueryValue( context, "section_id" ) );

    Model.RemoveSection( section_id );

    administration_controller.EditSections( context );
}
