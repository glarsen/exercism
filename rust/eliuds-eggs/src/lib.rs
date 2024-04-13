pub fn egg_count(display_value: u32) -> usize {
    let mut n = display_value;
    
    let mut count = 0;
    while n > 0 {
        count += n & 1;
        n >>= 1;
    }
    
    count as usize
}
