package Application;

// == IMPORTS

// == IMPORTS

import ( "fmt";
    "strconv" );

// == FUNCTIONS

func Dump( value ... interface {} ) {
    fmt.Printf( "%+v\n", value ... );
}

// ~~

func Check( error_ error ) {
    if error_ != nil {
        fmt.Println( error_ );
    }
}

// ~~

func GetIntegerFromString64( text string ) int64 {
    integer, error := strconv.ParseInt( text, 10, 64 );
    Check( error );

    return integer;
}

// ~~

func GetIntegerFromString( text string ) int {
    integer, error := strconv.ParseInt( text, 10, 64 );
    Check( error );

    return int( integer );
}

// ~~

func GetStringReal64( text string ) float64 {
    real, error := strconv.ParseFloat( text, 64 );
    Check( error );

    return real;
}

// ~~

func GetStringReal( text string ) float32 {
    real, error := strconv.ParseFloat( text, 64 );
    Check( error );

    return float32( real );
}

// ~~

func GetIntegerString64( integer int64 ) string {
    return strconv.FormatInt( integer, 10 );
}

// ~~

func GetIntegerString( integer int ) string {
    return strconv.FormatInt( int64( integer ), 10 );
}

// ~~

func GetRealString64( real float64, digit_count int ) string {
    return strconv.FormatFloat( real, 'f', digit_count, 64 );
}

// ~~

func RealString( real float32, digit_count int ) string {
    return strconv.FormatFloat( float64( real ), 'f', digit_count, 64 );
}
