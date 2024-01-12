/// Given an array of integers, return a new array such that each element at index i of the new array
/// is the product of all the numbers in the original array except the one at i.
/// For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
/// If our input was [3, 2, 1], the expected output would be [2, 3, 6].

#[allow(dead_code)]
fn product_all_others<const N: usize>(input: &[isize; N]) -> [isize; N] {
    let mut left = [1; N];
    let mut right = [1; N];
    let mut output = [1; N];

    for i in (0..N).skip(1) {
        left[i] = left[i - 1] * input[i - 1]
    }

    for i in (0..N).rev().skip(1) {
        right[i] = right[i + 1] * input[i + 1]
    }

    for i in 0..N {
        output[i] = left[i] * right[i]
    }

    output
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(&[1, 2, 3, 4, 5], [120, 60, 40, 30, 24])]
    #[case(&[3, 2, 1], [2, 3, 6])]
    #[case(&[-3, 2, 1], [2, -3, -6])]
    #[case(&[], [])]
    #[case(&[1], [1])]
    #[case(&[1, 1], [1, 1])]
    fn product_all_others_test<const N: usize>(
        #[case] numbers: &[isize; N],
        #[case] expected: [isize; N],
    ) {
        assert_eq!(expected, product_all_others(numbers))
    }
}
