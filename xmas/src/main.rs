use chrono::{Datelike, Local, NaiveDate};
use notify_rust::Notification;

fn main() {
    let today = Local::now().date_naive();

    let current_year = today.year();
    let christmas =
        NaiveDate::from_ymd_opt(current_year, 12, 25).expect("Failed to create Christmas date");

    let target_christmas = if today > christmas {
        NaiveDate::from_ymd_opt(current_year + 1, 12, 25)
            .expect("Failed to create next year's Christmas date")
    } else {
        christmas
    };

    let days_until = target_christmas.signed_duration_since(today).num_days();

    let message = if days_until == 0 {
        "Merry Christmas! 🎄".to_string()
    } else {
        format!("{} days until Christmas! 🎄", days_until)
    };

    Notification::new()
        .summary("xmas")
        .body(&message)
        .icon("christmas-tree")
        .timeout(5000)
        .show()
        .expect("Failed to show notification");
}
