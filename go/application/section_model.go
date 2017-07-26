package Application;

// == FUNCTIONS

func ( model * MODEL ) GetSectionArray( section_array * [] SECTION ) {
    Check( database.Select( section_array, "select * from SECTION order by Number ASC" ) );
}

// ~~

func ( model * MODEL ) GetSectionById( section * SECTION, id uint32 ) {
    Check( database.Get( section, "select * from SECTION where Id = ?", id ) );
}

// ~~

func ( model * MODEL ) ChangeSection( id uint32, number uint32, name string, text string, image string ) {
    database.MustExec( "update SECTION set Id = ?, Number = ?, Name = ?, Text = ?, Image = ? where Id = ?", id, number, name, text, image, id );
}

// ~~

func ( model * MODEL ) AddSection( number uint32, name string, text string, image string ) {
    database.MustExec( "insert into SECTION ( Number, Name, Text, Image ) values ( ?, ?, ?, ? )", number, name, text, image );
}

// ~~

func ( model * MODEL ) RemoveSection( id uint32 ) {
    database.MustExec( "delete from SECTION where Id = ?", id );
}
