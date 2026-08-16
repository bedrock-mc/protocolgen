// This file is copied into the pinned Endstone checkout by the vanilla-data
// workflow. It is kept separate from protocolgen's Go code because the data
// is only available through Endstone's in-process BDS symbols.
#pragma once

#include <filesystem>

namespace endstone::core::devtools {

// Returns true after all Endstone-native vanilla data files were written.
bool exportHeadlessVanillaData(const std::filesystem::path &base_path);

}  // namespace endstone::core::devtools
