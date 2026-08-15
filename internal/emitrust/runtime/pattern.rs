
// ---------------------------------------------------------------------------
// Schema string patterns
// ---------------------------------------------------------------------------

use std::collections::HashMap;
use std::sync::{LazyLock, Mutex};

static SCHEMA_PATTERNS: LazyLock<Mutex<HashMap<&'static str, regex::Regex>>> =
    LazyLock::new(|| Mutex::new(HashMap::new()));

pub fn assert_pattern(value: &str, pattern: &'static str) {
    assert!(validate_pattern(value, pattern).is_ok(), "string does not match schema pattern");
}

pub fn validate_pattern(value: &str, pattern: &'static str) -> DecodeResult<()> {
    let mut patterns = SCHEMA_PATTERNS.lock().expect("schema pattern cache poisoned");
    let compiled = patterns.entry(pattern).or_insert_with(|| {
        regex::Regex::new(pattern).expect("manifest contains an invalid schema pattern")
    });
    if compiled.is_match(value) {
        Ok(())
    } else {
        Err(DecodeError::SchemaConstraint("string does not match pattern"))
    }
}

#[cfg(test)]
mod pattern_tests {
    use super::*;

    #[test]
    fn patterns_are_enforced() {
        assert_eq!(validate_pattern("abc", "^[a-z]+$"), Ok(()));
        assert_eq!(
            validate_pattern("123", "^[a-z]+$"),
            Err(DecodeError::SchemaConstraint("string does not match pattern"))
        );
    }
}
