pub fn add(x: i32, y: i32) -> i32 {
    x + y
}

pub fn multiply(x: i32, y: i32) -> i32 {
    x * y
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
    }

    #[test]
    fn test_multiply() {
        assert_eq!(multiply(2, 3), 6);
    }
}
