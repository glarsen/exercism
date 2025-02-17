use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    nucleotide_counts(dna)?
        .get(&nucleotide)
        .ok_or(nucleotide)
        .copied()
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts = HashMap::from([('A', 0), ('C', 0), ('G', 0), ('T', 0)]);

    for nucleotide in dna.chars() {
        match counts.get_mut(&nucleotide) {
            Some(count) => *count += 1,
            None => return Err(nucleotide),
        }
    }

    Ok(counts)
}
