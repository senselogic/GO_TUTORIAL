package Application;

// == IMPORTS

import ( "github.com/jmoiron/sqlx"; );

// == FUNCTIONS

func ( model * MODEL ) OpenStore() {
    database = sqlx.MustConnect( "mysql", "root:root@tcp(localhost:3306)/blog" );
}

// ~~

func ( model * MODEL ) CloseStore() { }
