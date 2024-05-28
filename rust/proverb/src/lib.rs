use std::iter::once;

pub fn build_proverb(list: &[&str]) -> String {
    match list.first() {
        None => { "".to_string() },
        Some(first) => {
            list.windows(2)
                .map(|w| {
                    format!("For want of a {} the {} was lost.\n", w[0], w[1])
                })
                .chain(once(format!("And all for the want of a {}.", first)))
                .collect()
        }
    }
}
