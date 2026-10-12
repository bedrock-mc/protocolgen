//! A global allocator that can refuse large requests, so the allocation
//! failure probe sees a decoder's fallible paths without exhausting memory.

use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicUsize, Ordering};

static LIMIT: AtomicUsize = AtomicUsize::new(usize::MAX);

struct Refusing;

#[global_allocator]
static GLOBAL: Refusing = Refusing;

fn refuses(size: usize) -> bool {
    size > LIMIT.load(Ordering::Relaxed)
}

// SAFETY: every call forwards to `System` unchanged, or returns null, which
// `GlobalAlloc` permits for any request.
unsafe impl GlobalAlloc for Refusing {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        if refuses(layout.size()) {
            return std::ptr::null_mut();
        }
        unsafe { System.alloc(layout) }
    }

    unsafe fn alloc_zeroed(&self, layout: Layout) -> *mut u8 {
        if refuses(layout.size()) {
            return std::ptr::null_mut();
        }
        unsafe { System.alloc_zeroed(layout) }
    }

    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) {
        unsafe { System.dealloc(ptr, layout) }
    }

    unsafe fn realloc(&self, ptr: *mut u8, layout: Layout, new_size: usize) -> *mut u8 {
        if refuses(new_size) {
            return std::ptr::null_mut();
        }
        unsafe { System.realloc(ptr, layout, new_size) }
    }
}

/// Runs `probe` with every allocation larger than `bytes` refused. The
/// runner is single-threaded, so nothing else allocates meanwhile.
pub fn refusing_above<T>(bytes: usize, probe: impl FnOnce() -> T) -> T {
    struct Restore;
    impl Drop for Restore {
        fn drop(&mut self) {
            LIMIT.store(usize::MAX, Ordering::Relaxed);
        }
    }
    LIMIT.store(bytes, Ordering::Relaxed);
    let _restore = Restore;
    probe()
}
