// ---------------------------------------------------------------------------
// Schema string patterns
// ---------------------------------------------------------------------------

use std::collections::HashMap;
use std::sync::{LazyLock, Mutex};

static SCHEMA_PATTERNS: LazyLock<Mutex<HashMap<&'static str, regex::Regex>>> =
    LazyLock::new(|| Mutex::new(HashMap::new()));

fn matches_pattern(value: &Str, pattern: &'static str) -> bool {
    let Ok(text) = value.to_str() else {
        return false;
    };
    let mut patterns = SCHEMA_PATTERNS.lock().expect("schema pattern cache poisoned");
    let compiled = patterns.entry(pattern).or_insert_with(|| {
        regex::Regex::new(pattern).expect("manifest contains an invalid schema pattern")
    });
    compiled.is_match(text)
}

/// Rejects an encoded string that is not UTF-8 or does not match `pattern`.
pub fn check_pattern(value: &Str, pattern: &'static str) -> EncodeResult<()> {
    if matches_pattern(value, pattern) {
        Ok(())
    } else {
        Err(EncodeError::SchemaConstraint("string does not match pattern"))
    }
}

/// Rejects a decoded string that is not UTF-8 or does not match `pattern`.
pub fn validate_pattern(value: &Str, pattern: &'static str) -> DecodeResult<()> {
    if matches_pattern(value, pattern) {
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
        assert_eq!(validate_pattern(&Str::from_static("abc"), "^[a-z]+$"), Ok(()));
        assert_eq!(
            validate_pattern(&Str::from_static("123"), "^[a-z]+$"),
            Err(DecodeError::SchemaConstraint("string does not match pattern"))
        );
        assert!(check_pattern(&Str::from(bytes::Bytes::from_static(&[0xff])), "^.*$").is_err());
    }
}
