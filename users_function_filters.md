# Users function filters

Let's try to go to the most basic implementation.

What we want is to have a way of calling users code, but that the code is limited by some filters.

```rust
#[repr(C)]
pub struct NetworkFilters {
    pub allowed_ips: *mut cty::c_char,
    pub allowed_domains: *mut cty::c_char,
}

#[repr(C)]
pub struct FilesystemFilters {
    pub allowed_paths_read: *mut cty::c_char,
    pub allowed_paths_write: *mut cty::c_char,
}

#[repr(C)]
pub struct Limits {
    pub open_files: cty::c_int,
    pub memory_bytes: cty::c_int,
    pub cpu_units: cty::c_int,
}

#[repr(C)]
pub struct UserFilters {
    pub network: *mut NetworkFilters,
    pub filesystem: *mut FilesystemFilters,
    pub limits: *mut Limits,
}
```

Then to call some user code, one would do something like the following:

```rust
#[no_mangle]
pub extern "C" fn invoke_user_function(user_fn: fn (*mut UserFilters) -> cty::c_int, filters: UserFilters) -> cty::c_int {
    return user_fn(*mut filters);
}
```

Where the variable `user_fn` would be a closure that has all necessary user state inside and that takes care of actually implementing all the filtering and limits passed in the `filters` object.
