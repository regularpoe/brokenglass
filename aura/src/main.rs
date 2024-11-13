use notify_rust::Notification;
use serde::{Deserialize, Serialize};

use std::collections::HashMap;
use std::fs::OpenOptions;
use std::io::Write;
use std::option::Option::None;
use std::process::Command;
use std::thread;
use std::time::Duration;

#[derive(Debug, Serialize, Deserialize, Clone, PartialEq)]
struct Cmus {
    status: String,
    artist: String,
    title: String,
    album: String,
}

impl Cmus {
    fn new(status_map: &HashMap<String, String>) -> Self {
        Cmus {
            status: status_map
                .get("status")
                .unwrap_or(&"Unknow status".to_string())
                .clone(),
            artist: status_map
                .get("artist")
                .unwrap_or(&"Unknown artist".to_string())
                .clone(),
            title: status_map
                .get("title")
                .unwrap_or(&"Unknown Title".to_string())
                .clone(),
            album: status_map
                .get("album")
                .unwrap_or(&"Unknown Album".to_string())
                .clone(),
        }
    }
}

fn get_cmus_status() -> Option<HashMap<String, String>> {
    let output = Command::new("cmus-remote").arg("-Q").output().ok().unwrap();

    if !output.status.success() {
        return None;
    }

    let status_str = String::from_utf8_lossy(&output.stdout);
    let mut status_map = HashMap::new();

    let parts: Vec<&str> = status_str.split("\n").collect();
    let status: Vec<&str> = parts[0].split(' ').collect();

    status_map.insert(status[0].to_string(), status[1].to_string());
    status_map.insert(
        "artist".to_string(),
        parts[4].replace("tag artist", "").to_string(),
    );
    status_map.insert(
        "album".to_string(),
        parts[5].replace("tag album", "").to_string(),
    );
    status_map.insert(
        "title".to_string(),
        parts[6].replace("tag title", "").to_string(),
    );

    Some(status_map)
}

fn log_track(track: &Cmus) -> std::io::Result<()> {
    let json = serde_json::to_string(&track)?;

    let mut file = OpenOptions::new()
        .create(true)
        .append(true)
        .open("cmus_history.jsonl")?;

    writeln!(file, "{}", json)?;
    println!("Logged: {} - {}", track.artist, track.title);
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut last_track: Option<Cmus> = None;

    loop {
        if let Some(status) = get_cmus_status() {
            if status.get("status").map_or(false, |s| s == "playing") {
                let current_track = Cmus::new(&status);

                if last_track.as_ref() != Some(&current_track) {
                    log_track(&current_track)?;
                    let message = format!("{} - {}", &current_track.artist, &current_track.title);

                    Notification::new()
                        .summary("Currently playing..")
                        .body(&message)
                        .timeout(5000)
                        .show()
                        .expect("Failed to show notification");
                    last_track = Some(current_track);
                }
            }
        }

        thread::sleep(Duration::from_secs(5));
    }
}
