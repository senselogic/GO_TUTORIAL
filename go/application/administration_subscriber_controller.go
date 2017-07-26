package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditSubscribers( context * golf.Context ) {
    var session SESSION;
    var subscriber_array [] SUBSCRIBER;

    Model.GetSession( context, &session );
    Model.GetSubscriberArray( &subscriber_array );

    data := map[ string ] interface {} {
        "Session" : session, "SubscriberArray" : subscriber_array };

    AdministrationView.RenderEditSubscribers( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) EditSubscriber( context * golf.Context ) {
    var session SESSION;
    var subscriber SUBSCRIBER;

    id := GetUint32FromString( GetQueryValue( context, "subscriber_id" ) );

    Model.GetSession( context, &session );
    Model.GetSubscriberById( &subscriber, id );

    data := map[ string ] interface {} {
        "Session" : session, "Subscriber" : subscriber };

    AdministrationView.RenderEditSubscriber( context, &data );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) ChangeSubscriber( context * golf.Context ) {
    id := GetUint32FromString( GetQueryValue( context, "subscriber_id" ) );

    context.Request.ParseForm();

    name := GetPostValue( context, "name" );
    email := GetPostValue( context, "email" );

    Model.ChangeSubscriber( id, name, email );

    administration_controller.EditSubscribers( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) AddSubscriber( context * golf.Context ) {
    context.Request.ParseForm();

    name := GetPostValue( context, "name" );
    email := GetPostValue( context, "email" );

    Model.AddSubscriber( name, email );

    administration_controller.EditSubscribers( context );
}

// ~~

func ( administration_controller * ADMINISTRATION_CONTROLLER ) RemoveSubscriber( context * golf.Context ) {
    subscriber_id := GetUint32FromString( GetQueryValue( context, "subscriber_id" ) );

    Model.RemoveSubscriber( subscriber_id );

    administration_controller.EditSubscribers( context );
}
