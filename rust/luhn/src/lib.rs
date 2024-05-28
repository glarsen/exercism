/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    code.chars()
        .rev() // Starting from the right
        .filter(|c| !c.is_whitespace()) // Ignore whitespaces
        .try_fold((0, 0), |(sum, count), val| {
            val.to_digit(10)
                // Double every second digit
                .map(|n| if count % 2 == 1 {n * 2} else {n})
                // If the resulting digit is greater than 9, subtract 9
                .map(|n| if n > 9 {n - 9} else {n})
                // Add the (modified) digit to the sum and add it to the running count
                .map(|n| (n + sum, count + 1))
        })
        // Resulting sum of at least two digits must be evenly divisible by 10
        .map_or(false, |(sum, count)| sum % 10 == 0 && count > 1)
}
