package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );
    
// == FUNCTIONS

func ( controller * CONTROLLER ) ConnectUser(
    context * golf.Context
    )
{
    var session SESSION;

    context.Request.ParseForm();

    pseudonym := GetPostValue( context, "pseudonym" );
    password := GetPostValue( context, "password" );

    Model.GetSession( context, &session );
    
    if ( Model.FindUserByPseudonymAndPassword( &session.User, pseudonym, password ) )
    {
        session.UserIsConnected = true;
        Model.SetSession( context, &session );
    }

    controller.ShowHome( context );
}

// ~~

func ( controller * CONTROLLER ) DisconnectUser(
    context * golf.Context
    )
{
    var session SESSION;

    Model.GetSession( context, &session );
    session.UserIsConnected = false;
    Model.SetSession( context, &session );
    
    controller.ShowHome( context );
}
