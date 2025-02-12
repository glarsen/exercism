use rand::prelude::*;

// Pick a private key greater than 1 and less than p
pub fn private_key(p: u64) -> u64 {
    rand::thread_rng().gen_range(2..p)
}

// Fast modular exponentiation
// Ref: https://rob.co.bb/posts/2019-02-10-modular-exponentiation-in-rust/
fn mod_pow(mut base: u128, mut exp: u64, modulus: u64) -> u64 {
    if modulus == 1 {
        return 0;
    }
    let mut result = 1;
    base %= modulus as u128;

    while exp > 0 {
        if exp % 2 == 1 {
            result = result * base % modulus as u128;
        }
        exp >>= 1;
        base = base * base % modulus as u128;
    }

    result as u64
}

// Calculate public key using prime numbers {p} and {g}, and private key {a}
// A = gᵃ mod p
pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    mod_pow(g as u128, a, p)
}

// Calculate secret key using prime number {p}, public key {b_pub}, and private key {a}
pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    public_key(p, b_pub, a)
}
