package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditSubscribers( context * golf.Context, data * map[ string ] interface {} ) {
    context.Loader( "template" ).Render( "administration_edit_subscribers.html", *data );
}

// ~~

func ( administration_view * ADMINISTRATION_VIEW ) RenderEditSubscriber( context * golf.Context, data * map[ string ] interface {} ) {
    context.Loader( "template" ).Render( "administration_edit_subscriber.html", *data );
}
