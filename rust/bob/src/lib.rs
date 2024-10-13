pub fn reply(message: &str) -> &str {
    let message: String = message
        .chars()
        .filter(|c| c.is_alphabetic() || c.is_ascii_punctuation())
        .collect();
    
    if message.is_empty() {
        return "Fine. Be that way!"
    }

    let yelling = message.chars().any(|c| c.is_alphabetic()) && message.to_uppercase() == message;
    let question = message.ends_with('?');

    match (yelling, question) {
        (true, true) => "Calm down, I know what I'm doing!",
        (true, false) => "Whoa, chill out!",
        (false, true) => "Sure.",
        (false, false) => "Whatever."
    }
}
