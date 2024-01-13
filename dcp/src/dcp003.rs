/// Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
/// For example, given the following Node class
///
/// class Node:
///     def __init__(self, val, left=None, right=None):
///         self.val = val
///         self.left = left
///         self.right = right
///
/// The following test should pass:
///
/// node = Node('root', Node('left', Node('left.left')), Node('right'))
/// assert deserialize(serialize(node)).left.left.val == 'left.left'

type NodeRef = Option<Box<Node>>;

struct Node {
    value: String,
    left: NodeRef,
    right: NodeRef,
}

#[allow(dead_code)]
impl Node {
    fn new(value: impl ToString, left: NodeRef, right: NodeRef) -> Self {
        Self {
            value: value.to_string(),
            left,
            right,
        }
    }

    fn leaf(value: impl ToString) -> Self {
        Self {
            value: value.to_string(),
            left: None,
            right: None,
        }
    }

    fn serialize(&self) -> String {
        fn recurse(node: &Node, output: &mut String) {
            output.push_str(&node.value);
            output.push(',');

            for maybe_node in [&node.left, &node.right] {
                if let Some(node) = maybe_node {
                    recurse(node, output);
                } else {
                    output.push_str("null,");
                }
            }
        }

        let mut output = String::new();
        recurse(&self, &mut output);

        output
    }

    fn deserialize(input: &str) -> NodeRef {
        fn recurse(rest: &str) -> (NodeRef, &str) {
            if rest.is_empty() {
                return (None, rest);
            }

            let boundary = rest.find(',').unwrap();
            let value = &rest[0..boundary];
            let rest = &rest[boundary + 1..];

            if value == "null" {
                return (None, rest);
            }

            let (left, rest) = recurse(rest);
            let (right, rest) = recurse(rest);

            (Some(Box::new(Node::new(value, left, right).into())), rest)
        }

        recurse(input).0
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use rstest::rstest;

    #[rstest]
    #[case(Node::new("root", None, None))]
    #[case(Node::new("root", None, Some(Node::leaf("right").into())))]
    #[case(Node::new("root", Some(Node::leaf("right").into()), None))]
    #[case(Node::new(
        "root",
       Some(Node::new("left", Some(Node::leaf("left.left").into()), None).into()),
      Some(Node::leaf("right").into())  
    ))]
    fn binary_tree_serialize_deserialize_test(#[case] node: Node) {
        let serialized = node.serialize();
        let deserialized = Node::deserialize(&serialized);
        let roundtrip = deserialized.unwrap().serialize();

        assert_eq!(serialized, roundtrip)
    }
}
