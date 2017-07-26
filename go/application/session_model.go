package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( model * MODEL ) GetSession( context * golf.Context, session * SESSION ) {
    object, error_ := context.Session.Get( "session" );

    if ( error_ == nil ) { *session, _ = object.( SESSION );
    }
}

// ~~

func ( model * MODEL ) SetSession( context * golf.Context, session * SESSION ) {
    context.Session.Set( "session", *session );
}
