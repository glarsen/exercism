pub fn collatz(n: u64) -> Option<u64> {
    match n {
        0 => None,
        1 => Some(0),
        _ => Some(1 + collatz(if n % 2 == 0 { n / 2 } else { 3 * n + 1 })?),
    }
}
