use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::{TcpListener, TcpStream};
use tokio_native_tls::native_tls::{self, Identity};
use tokio_native_tls::TlsAcceptor;

use std::env;

async fn list_files() -> Result<std::process::Output, Box<dyn std::error::Error>> {
    let output: std::process::Output = tokio::process::Command::new("sh")
        .arg("-c")
        .arg("ls -lah")
        .output()
        .await?;

    Ok(output)
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:2408".to_string());

    let der = include_bytes!("identity.pfx");
    let cert = Identity::from_pkcs12(der, "dispatch")?;
    let tls_acceptor: tokio_native_tls::TlsAcceptor =
        tokio_native_tls::TlsAcceptor::from(native_tls::TlsAcceptor::builder(cert).build()?);

    let listener = TcpListener::bind(&addr).await?;

    println!("\nwiretrap {} on {}", env!("CARGO_PKG_VERSION"), addr);

    loop {
        let (socket, _) = listener.accept().await?;

        let tls_acceptor = tls_acceptor.clone();

        tokio::spawn(async move {
            process(socket, tls_acceptor).await;
        });
    }
}

async fn process(socket: TcpStream, tls_acceptor: TlsAcceptor) {
    let mut tls_stream = tls_acceptor.accept(socket).await.expect("Failed to accept");

    let mut buffer = [0; 1024];

    loop {
        let n = match tls_stream.read(&mut buffer).await {
            Ok(0) => return,
            Ok(data) => data,
            Err(err) => {
                eprintln!("Failed to read data from the socket: {}", err);
                return;
            }
        };

        let command = String::from_utf8(buffer[0..n].to_vec()).unwrap();

        match command.trim() {
            "foo" => {
                let cmd_output = match list_files().await {
                    Ok(output) => {
                        println!("Command executed successfully");
                        output
                    }
                    Err(err) => {
                        eprintln!("Error executing command: {}", err);
                        return;
                    }
                };

                let response = String::from_utf8_lossy(&cmd_output.stdout).to_string();
                if let Err(err) = tls_stream.write_all(response.as_bytes()).await {
                    eprintln!("Error sending response: {}", err);
                }
                // return;
            }
            "bar" => println!("bar called"),
            "exit" => {
                println!("Exiting..\n");
                if let Err(err) = tls_stream.write(b"Goodbye!\n").await {
                    eprintln!("Error sending response: {}", err);
                }
                return;
            }
            _ => println!("unknown command"),
        }
    }
}
