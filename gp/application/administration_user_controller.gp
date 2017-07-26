package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );

// == FUNCTIONS

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditUsers(
    context * golf.Context
    )
{
    var session SESSION;
    var user_array [] USER;

    Model.GetSession( context, &session );
    Model.GetUserArray( &user_array );

    data := map[ string ] interface {}
    {
        "Session" : session,
        "UserArray" : user_array
    };

    AdministrationView.RenderEditUsers( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditUser(
    context * golf.Context
    )
{
    var session SESSION;
    var user USER;

    id := GetUint32FromString( GetQueryValue( context, "user_id" ) );

    Model.GetSession( context, &session );
    Model.GetUserById( &user, id );

    data := map[ string ] interface {}
    {
        "Session" : session,
        "User" : user
    };

    AdministrationView.RenderEditUser( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) ChangeUser(
    context * golf.Context
    )
{
    id := GetUint32FromString( GetQueryValue( context, "user_id" ) );
    
    context.Request.ParseForm();

    email := GetPostValue( context, "email" );
    pseudonym := GetPostValue( context, "pseudonym" );
    password := GetPostValue( context, "password" );
    it_is_administrator := GetUint32FromString( GetPostValue( context, "it_is_administrator" ) );

    Model.ChangeUser( id, email, pseudonym, password, it_is_administrator );

    administration_controller.EditUsers( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) AddUser(
    context * golf.Context
    )
{
    context.Request.ParseForm();

    email := GetPostValue( context, "email" );
    pseudonym := GetPostValue( context, "pseudonym" );
    password := GetPostValue( context, "password" );
    it_is_administrator := GetUint32FromString( GetPostValue( context, "it_is_administrator" ) );

    Model.AddUser( email, pseudonym, password, it_is_administrator );

    administration_controller.EditUsers( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) RemoveUser(
    context * golf.Context
    )
{
    user_id := GetUint32FromString( GetQueryValue( context, "user_id" ) );

    Model.RemoveUser( user_id );

    administration_controller.EditUsers( context );
}
