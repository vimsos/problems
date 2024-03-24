// Given a mapping of digits to letters (as in a phone number), and a digit string, return all possible letters the number could represent. You can assume each valid number in the mapping is a single digit.

// For example if {"2": ["a", "b", "c"], 3: ["d", "e", "f"], …} then "23" should return ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

use std::collections::HashMap;

#[allow(dead_code)]
fn possible_phone_words(map: &HashMap<char, Vec<char>>, input: &str) -> Vec<String> {
    fn recurse(map: &HashMap<char, Vec<char>>, input: &str, state: &Vec<String>) -> Vec<String> {
        if input.len() == 0 {
            return state.clone();
        }

        let key = input.chars().nth(0).unwrap();
        let possible_chars = map.get(&key).unwrap();
        let mut new_state = Vec::new();

        for word in state {
            for c in possible_chars {
                let new_word = format!("{}{}", word, c);
                new_state.push(new_word);
            }
        }

        recurse(map, &input[1..], &new_state)
    }

    recurse(map, input, &vec!["".to_string()])
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn possible_phone_texts_test_1() {
        let map = [('2', vec!['a', 'b', 'c']), ('3', vec!['d', 'e', 'f'])]
            .iter()
            .cloned()
            .collect();

        assert_eq!(
            vec!["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"],
            possible_phone_words(&map, "23")
        )
    }
}
