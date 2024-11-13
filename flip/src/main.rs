use base64::{engine::general_purpose, Engine as _};
use clipboard::{ClipboardContext, ClipboardProvider};
use rdev::{listen, Event, EventType, Key};

use std::sync::{Arc, Mutex};
use std::thread;

fn encode_clipboard(ctx: Arc<Mutex<ClipboardContext>>) {
    let mut ctx = ctx.lock().unwrap();
    if let Ok(content) = ctx.get_contents() {
        let encoded_content = general_purpose::STANDARD.encode(content);
        ctx.set_contents(encoded_content).unwrap();
        println!("Clipboard content encoded to Base64!");
    } else {
        println!("Failed to read clipboard content.");
    }
}

fn handle_event(
    event: Event,
    ctx: Arc<Mutex<ClipboardContext>>,
    ctrl_pressed: &mut bool,
    shift_pressed: &mut bool,
) {
    match event.event_type {
        EventType::KeyPress(Key::ControlLeft) => *ctrl_pressed = true,
        EventType::KeyRelease(Key::ControlLeft) => *ctrl_pressed = false,
        EventType::KeyPress(Key::ShiftLeft) => *shift_pressed = true,
        EventType::KeyRelease(Key::ShiftLeft) => *shift_pressed = false,
        EventType::KeyPress(Key::KeyC) => {
            if *ctrl_pressed && *shift_pressed {
                encode_clipboard(ctx.clone());
            }
        }
        _ => (),
    }
}

fn main() {
    let ctx = Arc::new(Mutex::new(ClipboardProvider::new().unwrap()));
    let ctx_clone = ctx.clone();

    let mut ctrl_pressed = false;
    let mut shift_pressed = false;

    let handle = thread::spawn(move || {
        listen(move |event| {
            handle_event(
                event,
                ctx_clone.clone(),
                &mut ctrl_pressed,
                &mut shift_pressed,
            )
        })
        .unwrap();
    });

    handle.join().unwrap();
}
