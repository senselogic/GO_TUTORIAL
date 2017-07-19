package Application;

// == FUNCTIONS

func ( model * MODEL ) GetUserArray( user_array * [] USER ) {
    Check( database.Select( user_array, "SELECT * FROM user ORDER BY email ASC" ) );
}

// ~~

func ( model * MODEL ) GetUserById( user * USER, id int ) {
    Check( database.Get( user, "SELECT * FROM user WHERE id = ?", id ) );
}

// ~~

func ( model * MODEL ) FindUserByPseudonymAndPassword( user * USER, pseudonym string, password string ) bool {
    return database.Get( user, "SELECT * FROM user WHERE pseudonym = ? AND password = ?", pseudonym, password ) == nil;
}

// ~~

func ( model * MODEL ) ChangeUser( id int, email string, pseudonym string, password string, it_is_administrator int ) {
    database.MustExec( "UPDATE user SET id = ?, email = ?, pseudonym = ?, password = ?, it_is_administrator = ? WHERE id = ?", id, email, pseudonym, password, it_is_administrator, id );
}

// ~~

func ( model * MODEL ) AddUser( email string, pseudonym string, password string, it_is_administrator int ) {
    database.MustExec( "INSERT INTO user ( email, pseudonym, password, it_is_administrator ) VALUES ( ?, ?, ?, ? )", email, pseudonym, password, it_is_administrator );
}

// ~~

func ( model * MODEL ) RemoveUser( id int ) {
    database.MustExec( "DELETE FROM user WHERE id = ?", id );
}
