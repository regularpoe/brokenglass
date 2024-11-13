use image::codecs::png::{CompressionType, FilterType, PngEncoder};
use image::Luma;
use qrcode::{EcLevel, QrCode};

use std::io::{self, BufWriter, Write};
use std::{env, process};

pub type StdErr = Box<dyn std::error::Error>;

pub fn generate_qr_code(text: &str) -> Result<Vec<u8>, StdErr> {
    let qr = QrCode::with_error_correction_level(text, EcLevel::L)?;
    let img_buf = qr.render::<Luma<u8>>().min_dimensions(200, 200).build();
    let mut encoded_buf = Vec::with_capacity(512);
    let encoder = PngEncoder::new_with_quality(
        &mut encoded_buf,
        CompressionType::Default,
        FilterType::NoFilter,
    );
    img_buf.write_with_encoder(encoder)?;
    Ok(encoded_buf)
}

fn main() -> Result<(), StdErr> {
    let mut args = env::args();
    if args.len() != 2 {
        eprintln!("Usage: qr <text>");
        process::exit(1);
    }

    let text = args.nth(1).unwrap();
    let qr_png = generate_qr_code(&text)?;

    let stdout = io::stdout();
    let mut handle = BufWriter::new(stdout.lock());
    handle.write_all(&qr_png)?;

    Ok(())
}
