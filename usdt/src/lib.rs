use std::os::raw::c_char;

#[no_mangle]
pub extern "C" fn stacks_enter(span_id: u64, parent_span_id: u64, id: u64, amount: u64, span_name: *const c_char) {    
    probe::probe!(
        stacks_tracing,
        enter,
        span_id,
        parent_span_id,
        id,
        amount,
        span_name
    );
}

#[no_mangle]
pub extern "C" fn stacks_close(span_id: u64) {
    probe::probe!(
        stacks_tracing,
        close,
        span_id
    );
}