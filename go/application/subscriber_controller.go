package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( controller * CONTROLLER ) AddSubscriber( context * golf.Context ) {
    var session SESSION;

    context.Request.ParseForm();

    name := GetPostValue( context, "name" );
    email := GetPostValue( context, "email" );

    Model.AddSubscriber( name, email );

    Model.GetSession( context, &session );
    session.UserHasSubscribed = true;
    Model.SetSession( context, &session );

    controller.ShowHome( context );
}
