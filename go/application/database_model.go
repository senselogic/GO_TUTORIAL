package Application;

// == IMPORTS

import ( "github.com/jmoiron/sqlx"; );

// == FUNCTIONS

func ( model * MODEL ) OpenDatabase() {
    database = sqlx.MustConnect( "mysql", "root:root@tcp(localhost:3306)/BLOG" );
}

// ~~

func ( model * MODEL ) CloseDatabase() { }
