package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == FUNCTIONS

func GetQueryValue( context * golf.Context, name string ) string {
    return context.Param( name );
}

// ~~

func GetPostValue( context * golf.Context, name string ) string {
    value, error := context.Query( name );
    Check( error );

    return value;
}
