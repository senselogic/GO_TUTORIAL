package Application;

// == FUNCTIONS

func ( model * MODEL ) GetSubscriberArray(
    subscriber_array * [] SUBSCRIBER
    )
{
    Check( database.Select( subscriber_array, "select * from SUBSCRIBER order by Email ASC" ) );
}

// ~~

func ( model * MODEL ) GetSubscriberById(
    subscriber * SUBSCRIBER,
    id uint32
    )
{
    Check( database.Get( subscriber, "select * from SUBSCRIBER where Id = ?", id ) );
}

// ~~

func ( model * MODEL ) ChangeSubscriber(
    id uint32,
    name string,
    email string
    )
{
    database.MustExec( 
        "update SUBSCRIBER set Id = ?, Name = ?, Email = ? where Id = ?", 
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
        "insert into SUBSCRIBER ( Name, Email ) values ( ?, ? )", 
        name, 
        email 
        );
}


// ~~

func ( model * MODEL ) RemoveSubscriber(
    id uint32
    )
{
    database.MustExec( "delete from SUBSCRIBER where Id = ?", id );
}
