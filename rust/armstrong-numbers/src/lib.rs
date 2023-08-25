pub fn is_armstrong_number(num: u32) -> bool {
    let digits: Vec<u64> = num
        .to_string()
        .chars()
        .map(|c| (c.to_string()).parse::<u64>().unwrap())
        .collect();

    let length = digits.len() as u32;

    digits.iter().map(|d| d.pow(length)).sum::<u64>() == num as u64
}
