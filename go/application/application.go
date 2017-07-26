package Application;

// == IMPORTS

import ( "github.com/dinever/golf"; );

// == VARIABLES

var Controller CONTROLLER;
var AdministrationController ADMINISTRATION_CONTROLLER;
var View VIEW;
var AdministrationView ADMINISTRATION_VIEW;
var Model MODEL;

// -- ROUTER

func Run() {
    Model.OpenDatabase();

    defer Model.CloseDatabase();

    Engine := golf.New();

    Engine.View.SetTemplateLoader( "template", "dynamic/" );
    Engine.SessionManager = golf.NewMemorySessionManager();
    Engine.Use( golf.SessionMiddleware );

    Engine.Static( "/static/", "static" );

    Engine.Get( "/", Controller.ShowHome );
    Engine.Get( "/show_section/:section_id/", Controller.ShowSection );
    Engine.Get( "/show_article/:article_id/", Controller.ShowArticle );
    Engine.Post( "/add_comment/:article_id/", Controller.AddComment );
    Engine.Post( "/add_subscriber/", Controller.AddSubscriber );

    Engine.Post( "/connect_user/", Controller.ConnectUser );
    Engine.Get( "/disconnect_user/", Controller.DisconnectUser );

    Engine.Get( "/administration/", AdministrationController.EditSections )

    Engine.Get( "/administration/edit_sections/", AdministrationController.EditSections )
    Engine.Get( "/administration/edit_section/:section_id/", AdministrationController.EditSection );
    Engine.Post( "/administration/change_section/:section_id/", AdministrationController.ChangeSection );
    Engine.Post( "/administration/add_section/", AdministrationController.AddSection );
    Engine.Get( "/administration/remove_section/:section_id/", AdministrationController.RemoveSection );

    Engine.Get( "/administration/edit_articles/", AdministrationController.EditArticles );
    Engine.Get( "/administration/edit_article/:article_id/", AdministrationController.EditArticle );
    Engine.Post( "/administration/change_article/:article_id/", AdministrationController.ChangeArticle );
    Engine.Post( "/administration/add_article/", AdministrationController.AddArticle );
    Engine.Get( "/administration/remove_article/:article_id/", AdministrationController.RemoveArticle );

    Engine.Get( "/administration/edit_comments/", AdministrationController.EditComments );
    Engine.Get( "/administration/edit_comment/:comment_id/", AdministrationController.EditComment );
    Engine.Post( "/administration/change_comment/:comment_id/", AdministrationController.ChangeComment );
    Engine.Post( "/administration/add_comment/", AdministrationController.AddComment );
    Engine.Get( "/administration/remove_comment/:comment_id/", AdministrationController.RemoveComment );

    Engine.Get( "/administration/edit_users/", AdministrationController.EditUsers );
    Engine.Get( "/administration/edit_user/:user_id/", AdministrationController.EditUser );
    Engine.Post( "/administration/change_user/:user_id/", AdministrationController.ChangeUser );
    Engine.Post( "/administration/add_user/", AdministrationController.AddUser );
    Engine.Get( "/administration/remove_user/:user_id/", AdministrationController.RemoveUser );

    Engine.Get( "/administration/edit_subscribers/", AdministrationController.EditSubscribers );
    Engine.Get( "/administration/edit_subscriber/:subscriber_id/", AdministrationController.EditSubscriber );
    Engine.Post( "/administration/change_subscriber/:subscriber_id/", AdministrationController.ChangeSubscriber );
    Engine.Post( "/administration/add_subscriber/", AdministrationController.AddSubscriber );
    Engine.Get( "/administration/remove_subscriber/:subscriber_id/", AdministrationController.RemoveSubscriber );

    Engine.Run( ":9000" );
}
