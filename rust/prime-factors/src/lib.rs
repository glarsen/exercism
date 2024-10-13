pub fn factors(mut n: u64) -> Vec<u64> {
    let mut factors = Vec::new();
    let mut candidates = 2..;

    while n > 1 {
        let f = candidates.next().unwrap();
        
        while n % f == 0 {
            factors.push(f);
            n /= f;
        }
    }

    factors
}
