// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerRegistryCleanup {
    pub removed_containers: Vec<FullContainerName>,
}

impl ContainerRegistryCleanup {
    pub const ID: u32 = 317;
}
