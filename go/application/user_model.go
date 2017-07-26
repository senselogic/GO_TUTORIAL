package Application;

// == FUNCTIONS

func ( model * MODEL ) GetUserArray( user_array * [] USER ) {
    Check( database.Select( user_array, "select * from USER order by Email ASC" ) );
}

// ~~

func ( model * MODEL ) GetUserById( user * USER, id uint32 ) {
    Check( database.Get( user, "select * from USER where Id = ?", id ) );
}

// ~~

func ( model * MODEL ) FindUserByPseudonymAndPassword( user * USER, pseudonym string, password string ) bool {
    return database.Get( user, "select * from USER where Pseudonym = ? and Password = ?", pseudonym, password ) == nil;
}

// ~~

func ( model * MODEL ) ChangeUser( id uint32, email string, pseudonym string, password string, it_is_administrator uint32 ) {
    database.MustExec( "update USER set Id = ?, Email = ?, Pseudonym = ?, Password = ?, ItIsAdministrator = ? where Id = ?", id, email, pseudonym, password, it_is_administrator, id );
}

// ~~

func ( model * MODEL ) AddUser( email string, pseudonym string, password string, it_is_administrator uint32 ) {
    database.MustExec( "insert into USER ( Email, Pseudonym, Password, ItIsAdministrator ) values ( ?, ?, ?, ? )", email, pseudonym, password, it_is_administrator );
}

// ~~

func ( model * MODEL ) RemoveUser( id uint32 ) {
    database.MustExec( "delete from USER where Id = ?", id );
}
