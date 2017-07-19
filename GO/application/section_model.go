package Application;

// == FUNCTIONS

func ( model * MODEL ) GetSectionArray( section_array * [] SECTION ) {
    Check( database.Select( section_array, "SELECT * FROM section ORDER BY number ASC" ) );
}

// ~~

func ( model * MODEL ) GetSectionById( section * SECTION, id int ) {
    Check( database.Get( section, "SELECT * FROM section WHERE id = ?", id ) );
}

// ~~

func ( model * MODEL ) ChangeSection( id int, number int, name string, text string, image string ) {
    database.MustExec( "UPDATE section SET id = ?, number = ?, name = ?, text = ?, image = ? WHERE id = ?", id, number, name, text, image, id );
}

// ~~

func ( model * MODEL ) AddSection( number int, name string, text string, image string ) {
    database.MustExec( "INSERT INTO section ( number, name, text, image ) VALUES ( ?, ?, ?, ? )", number, name, text, image );
}

// ~~

func ( model * MODEL ) RemoveSection( id int ) {
    database.MustExec( "DELETE FROM section WHERE id = ?", id );
}
