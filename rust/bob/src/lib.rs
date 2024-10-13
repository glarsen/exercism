pub fn reply(message: &str) -> &str {
    let message: String = message
        .chars()
        .filter(|c| c.is_alphabetic() || c.is_ascii_punctuation())
        .collect();

    let yelling = message.to_uppercase() == message && message.to_lowercase() != message;
    let question = message.ends_with('?');

    if message.is_empty() {
        "Fine. Be that way!"
    } else if yelling && question {
        "Calm down, I know what I'm doing!"
    } else if question {
        "Sure."
    } else if yelling {
        "Whoa, chill out!"
    } else {
        "Whatever."
    }
}
