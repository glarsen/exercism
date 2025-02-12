use std::collections::HashSet;

#[derive(Debug, PartialEq, Eq)]
pub struct Dna {
    sequence: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna {
    sequence: String,
}

impl Dna {
    // Construct new Dna from '{dna}' string. If string contains invalid
    // nucleotides, return index of first invalid nucleotide
    pub fn new(dna: &str) -> Result<Dna, usize> {
        Ok(Dna {
            sequence: validate(dna, HashSet::from(['G', 'C', 'T', 'A']))?,
        })
    }

    pub fn into_rna(self) -> Rna {
        let sequence = self
            .sequence
            .chars()
            .map(|c| match c {
                'G' => 'C',
                'C' => 'G',
                'T' => 'A',
                'A' => 'U',
                _ => unreachable!(),
            })
            .collect::<String>();

        Rna { sequence }
    }
}

impl Rna {
    // Construct new Rna from '{rna}' string. If string contains invalid
    // nucleotides, return index of first invalid nucleotide
    pub fn new(rna: &str) -> Result<Rna, usize> {
        Ok(Rna {
            sequence: validate(rna, HashSet::from(['C', 'G', 'A', 'U']))?,
        })
    }
}

fn validate(sequence: &str, expected: HashSet<char>) -> Result<String, usize> {
    sequence
        .chars()
        .enumerate()
        .map(|(i, c)| if expected.contains(&c) { Ok(c) } else { Err(i) })
        .collect::<Result<String, usize>>()
}
