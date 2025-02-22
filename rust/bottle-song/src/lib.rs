pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let words = std::collections::HashMap::from([
        (10, "Ten"),
        (9, "Nine"),
        (8, "Eight"),
        (7, "Seven"),
        (6, "Six"),
        (5, "Five"),
        (4, "Four"),
        (3, "Three"),
        (2, "Two"),
        (1, "One"),
        (0, "No"),
    ]);

    ((start_bottles - take_down + 1)..=start_bottles)
        .rev()
        .map(|bottles| {
            let (prefix, suffix) = match bottles {
                2 => (
                    String::from("Two green bottles hanging on the wall,"),
                    String::from("There'll be one green bottle hanging on the wall."),
                ),
                1 => (
                    String::from("One green bottle hanging on the wall,"),
                    String::from("There'll be no green bottles hanging on the wall."),
                ),
                _ => (
                    format!("{} green bottles hanging on the wall,", words[&bottles]),
                    format!(
                        "There'll be {} green bottles hanging on the wall.",
                        words[&(bottles - 1)].to_lowercase()
                    ),
                ),
            };

            format!(
                "{}\n{}\nAnd if one green bottle should accidentally fall,\n{}",
                prefix, prefix, suffix,
            )
        })
        .collect::<Vec<_>>()
        .join("\n\n")
}
