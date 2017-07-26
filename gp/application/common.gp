package Application;

// == IMPORTS

import (
    "fmt";    
    "strconv"
    );

// == FUNCTIONS

func Dump(
    value ... interface {}
    )
{
    fmt.Printf( "%+v\n", value ... );
}

// ~~

func Check(
    error_ error
    )
{
    if error_ != nil
    {
        fmt.Println( error_ );
    }
}

// ~~

func GetInt64FromString(
    text string
    ) int64
{
    integer, error := strconv.ParseInt( text, 10, 64 );
    Check( error );
    
    return integer;
}

// ~~

func GetInt32FromString(
    text string
    ) int32
{
    integer, error := strconv.ParseInt( text, 10, 64 );
    Check( error );
    
    return int32( integer );
}


// ~~

func GetUint32FromString(
    text string
    ) uint32
{
    integer, error := strconv.ParseInt( text, 10, 64 );
    Check( error );
    
    return uint32( integer );
}

// ~~

func GetFloat64FromString(
    text string
    ) float64
{
    real, error := strconv.ParseFloat( text, 64 );
    Check( error );
    
    return real;
}

// ~~

func GetFloat32FromString(
    text string
    ) float32
{
    real, error := strconv.ParseFloat( text, 64 );
    Check( error );
    
    return float32( real );
}

// ~~

func GetStringFromInt64(
    integer int64
    ) string
{
    return strconv.FormatInt( integer, 10 );
}

// ~~

func GetStringFromInt32(
    integer int32
    ) string
{
    return strconv.FormatInt( int64( integer ), 10 );
}

// ~~

func GetStringFromFloat64(
    real float64,
    digit_count int
    ) string
{
    return strconv.FormatFloat( real, 'f', digit_count, 64 );
}

// ~~

func GetStringFromFloat32(
    real float32,
    digit_count int
    ) string
{
    return strconv.FormatFloat( float64( real ), 'f', digit_count, 64 );
}
