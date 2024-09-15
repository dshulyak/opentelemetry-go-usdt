use std::os::raw::c_char;
use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn stacks_enter(span_id: u64, parent_span_id: u64, id: u64, amount: u64, span_name: *const c_char) {
    let span_name = unsafe {
        assert!(!span_name.is_null());
        CStr::from_ptr(span_name).to_str().unwrap()
    };
    
    probe::probe!(
        stacks_tracing,
        enter,
        span_id,
        parent_span_id,
        id,
        amount,
        span_name.as_ptr()
    );
}

#[no_mangle]
pub extern "C" fn stacks_exit(span_id: u64) {
    probe::probe!(
        stacks_tracing,
        exit,
        span_id
    );
}