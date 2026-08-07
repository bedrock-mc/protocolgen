// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AvailableCommands {
    pub enum_values: Vec<String>,
    pub chained_subcommand_values: Vec<String>,
    pub post_fixes: Vec<String>,
    pub enum_data: Vec<AvailableCommandsEnumData>,
    pub chained_subcommand_data: Vec<AvailableCommandsChainedSubcommandData>,
    pub commands: Vec<AvailableCommandsPacketCommandData>,
    pub soft_enums: Vec<AvailableCommandsSoftEnumData>,
    pub constraints: Vec<AvailableCommandsConstrainedValueData>,
}

impl AvailableCommands {
    pub const ID: u32 = 76;
}
