package Application;

// == IMPORTS

import (
    "github.com/dinever/golf";
    );
    
// == FUNCTIONS

func ( controller * CONTROLLER ) ShowHome(
    context * golf.Context
    )
{
    var session SESSION;
    var section_array [] SECTION;
    var article_array [] ARTICLE;

    Model.GetSession( context, &session );
    
    Model.GetSectionArray( &section_array );
    section := &section_array[ 0 ];
    Model.GetArticleArrayBySectionId( &article_array, section.Id );
    
    section.ImageIndex = section.Id % 20;

    data := map[ string ] interface {}
    {
        "Session" : session,
        "SectionArray" : section_array,
        "Section" : section,
        "ArticleArray" : article_array
    };

    View.RenderShowSection( context, &data );
}

// ~~

func ( controller * CONTROLLER ) ShowSection(
    context * golf.Context
    )
{
    var session SESSION;
    var section_array [] SECTION;
    var section SECTION;
    var article_array [] ARTICLE;

    section_id := GetIntegerFromString( GetQueryValue( context, "section_id" ) );

    Model.GetSession( context, &session );
    
    Model.GetSectionArray( &section_array );
    Model.GetSectionById( &section, section_id );
    Model.GetArticleArrayBySectionId( &article_array, section_id );
    
    section.ImageIndex = section.Id % 20;

    data := map[ string ] interface {}
    {
        "Session" : session,
        "SectionArray" : section_array,
        "Section" : section,
        "ArticleArray" : article_array
    };

    View.RenderShowSection( context, &data );
}
