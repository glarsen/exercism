pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    let students = [
        "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph",
        "Kincaid", "Larry",
    ];

    let index = students.iter().position(|&s| s == student).unwrap() * 2;

    diagram
        .lines()
        .flat_map(|line| {
            line[index..=index + 1].chars().map(|cup| match cup {
                'C' => "clover",
                'G' => "grass",
                'R' => "radishes",
                'V' => "violets",
                _ => unreachable!(),
            })
        })
        .collect()
}
