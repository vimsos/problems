use std::collections::HashSet;

/// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
/// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

#[allow(dead_code)]
fn sum_to_k(numbers: &[isize], k: isize) -> bool {
    let mut set = HashSet::<isize>::new();

    for n in numbers {
        let complement = k - *n;
        if let Some(_) = set.get(&complement) {
            return true;
        }
        set.insert(*n);
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(&[10,15,3,7], 17, true)]
    #[case(&[10,10,15,3,7], 17, true)]
    #[case(&[10,10,15,3], 17, false)]
    #[case(&[-10,10,15,3], 0, true)]
    fn sum_to_k_test(#[case] numbers: &[isize], #[case] k: isize, #[case] expected: bool) {
        assert_eq!(expected, sum_to_k(numbers, k))
    }
}
