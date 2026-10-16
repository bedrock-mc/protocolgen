/// Client-owned ceiling on recursive schema edges along a decoded value's path.
/// This bounds decoder stack use and the resulting owned tree's clone/drop depth;
/// it is independent of collection/byte limits and is not a native protocol maximum.
pub const MAX_RECURSIVE_VALUE_DEPTH: usize = 64;

impl Reader<'_> {
    /// Decodes one recursive child, restoring depth capacity on success or a decode error.
    /// Refusal consumes no child bytes; cursor changes made by `decode` are not rolled back.
    pub fn with_recursive_value<T>(
        &mut self,
        decode: impl FnOnce(&mut Self) -> DecodeResult<T>,
    ) -> DecodeResult<T> {
        if self.recursive_depth == MAX_RECURSIVE_VALUE_DEPTH {
            return Err(DecodeError::RecursiveDepthExceeded);
        }
        self.recursive_depth += 1;
        let result = decode(self);
        self.recursive_depth -= 1;
        result
    }
}
