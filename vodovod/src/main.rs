extern crate colorful;
use colorful::Color;
use colorful::Colorful;

extern crate select;
use select::document::Document;
use select::predicate::Class;

#[tokio::main]
async fn get_html(url: &str) -> Result<String, reqwest::Error> {
    let rsp = reqwest::get(url).await?;

    let body = rsp.text().await?;

    Ok(body)
}

fn main() {
    println!("\nvodovod {}", env!("CARGO_PKG_VERSION"));

    let url = "http://www.vodovod-pula.hr";

    let raw_html = match get_html(url) {
        Ok(data) => data,
        Err(why) => panic!("Couldn't get URL - error: {}", why),
    };

    let doc = Document::from(raw_html.as_str());

    let mut links: Vec<&str> = Vec::new();

    let mut n = 0;

    for node in doc.find(Class("mod-articles-category-title")) {
        links.push(node.attr("href").unwrap());
    }

    println!("\nHitne intervencije i obavijesti");

    for link in &links {
        match link.contains("veli") {
            true => {
                println!("{}", link.color(Color::Yellow));
                n += 1;
            }
            false => println!("{}", link),
        }
    }

    println!();

    if n != 0 && webbrowser::open(url).is_err() {
        println!("Visit the following URL in your browser");
    }
}
