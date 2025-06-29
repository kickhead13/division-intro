/*
 * Now let's do something that we actually use in the development
 * of division online... Expanding a very simple API.
 *
 * The code below creates an API server that exposes a single
 * "HTTP GET" API endpoint (i.e. "/hi/{name}") that returns a greeting message on
 * access.
 *
 * Example:
 *  If you run the program as is, and then go to localhost:1234/hi/alex in your
 *  browser, it should display the message "Hi, alex!". You can replace the name
 *  alex with any name and it will greet you accordingly...
 *
 * To do:
 *  Try to understand the code.
 *  Implement a goodbye "GET" endpoint that says "Goodbye, {name}!".
 *  Try implementing a POST endpoint (or DELETE, or PUT...).
 *
 *
 * This is basically 90% of what you are expected to do with Rust for the 
 * purpose of this projects... API's (and some Websockets).
 * */

use actix_web::{middleware::Logger};

#[actix_web::get("/hi/{name}")]
pub async fn get_hi_greeting(
    http: actix_web::HttpRequest
) -> impl actix_web::Responder {
    let name: String = http.match_info().get("name").unwrap().to_string();
    return actix_web::HttpResponse::Ok().body(
        format!("Hi, {}!", name)
    );
}

#[actix_web::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {


    let _ = actix_web::HttpServer::new(move || {
        actix_web::App::new()
            .service(get_hi_greeting) // Here we added the get_hi_greeting function !
            //.service(my_func) to add your function
    })
    .bind(("127.0.0.1", 1234))?
    .run()
    .await;
    Ok(())
}
