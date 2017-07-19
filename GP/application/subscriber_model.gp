package Application;

// == FUNCTIONS

func ( model * MODEL ) GetSubscriberArray(
    subscriber_array * [] SUBSCRIBER
    )
{
    Check( database.Select( subscriber_array, "SELECT * FROM subscriber ORDER BY email ASC" ) );
}

// ~~

func ( model * MODEL ) GetSubscriberById(
    subscriber * SUBSCRIBER,
    id int
    )
{
    Check( database.Get( subscriber, "SELECT * FROM subscriber WHERE id = ?", id ) );
}

// ~~

func ( model * MODEL ) ChangeSubscriber(
    id int,
    name string,
    email string
    )
{
    database.MustExec( 
        "UPDATE subscriber SET id = ?, name = ?, email = ? WHERE id = ?", 
        id, 
        name,
        email,
        id 
        );
}

// ~~

func ( model * MODEL ) AddSubscriber(
    name string,
    email string
    )
{
    database.MustExec( 
        "INSERT INTO subscriber ( name, email ) VALUES ( ?, ? )", 
        name, 
        email 
        );
}


// ~~

func ( model * MODEL ) RemoveSubscriber(
    id int
    )
{
    database.MustExec( "DELETE FROM subscriber WHERE id = ?", id );
}
