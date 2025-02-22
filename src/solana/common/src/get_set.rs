pub fn set_value(storage: &mut u8, value: u8) {
    *storage = value;
}

pub fn get_value(storage: &u8) -> u8 {
    *storage
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_set_get_value() {
        let mut storage: u8 = 0;
        set_value(&mut storage, 42);
        assert_eq!(get_value(&storage), 42);
    }
}
