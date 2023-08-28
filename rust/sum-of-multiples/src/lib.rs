pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    (1..limit)
        .filter(|multiple| {
            factors
                .iter()
                .any(|&factor| factor != 0 && multiple % factor == 0)
        })
        .sum()
}
