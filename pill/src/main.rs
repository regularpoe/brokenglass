use rand::Rng;

fn main() {
    let number = rand::thread_rng().gen::<u64>();
    let number_str = number.to_string();
    let first_digit = number_str.chars().next().unwrap();
    let last_digit = number_str.chars().last().unwrap();
    let middle_digit = number_str.chars().nth(number_str.len() / 2).unwrap();

    let new_number: u64 = format!("{}{}{}", first_digit, middle_digit, last_digit)
        .parse()
        .unwrap();

    if new_number % 2 == 0 {
        println!("yes");
    } else {
        println!("no");
    }
}
