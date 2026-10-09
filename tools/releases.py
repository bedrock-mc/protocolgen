#!/usr/bin/env python3
"""Check that the release index agrees with every committed codec and capture."""
import json
from pathlib import Path


def read(path):
    """Read a repository JSON document."""
    return json.loads(path.read_text())


def validate(root):
    """Reject missing snapshots, inconsistent targets and undeclared captures."""
    index = read(root / "data/source/releases.json")
    releases = index["releases"]
    if index["schema_version"] != 1:
        raise ValueError("unsupported release index schema")
    if index["default_protocol"] not in releases:
        raise ValueError("unknown default protocol")
    if "capture" not in releases[index["default_capture"]]:
        raise ValueError("default capture has no adapter")
    committed = {p.parent.name for p in (root / "generated").glob("*/manifest.json")}
    if committed != set(releases):
        raise ValueError(f"release index differs from committed snapshots: {committed ^ set(releases)}")
    for snapshot, release in releases.items():
        target = read(root / "generated" / snapshot / "manifest.json")["target"]
        if (release["snapshot"] != snapshot or target["minecraft_version"] != snapshot
                or target["protocol_version"] != release["protocol_version"]
                or release["channel"] not in ("retail", "preview")
                or release["minecraft_version"] not in release["aliases"]):
            raise ValueError(f"inconsistent release {snapshot}")
        source_path = root / "generated" / snapshot / "vanilla-source.json"
        if source_path.exists() != ("capture" in release):
            raise ValueError(f"capture availability differs for {snapshot}")
        if source_path.exists():
            source = read(source_path)
            if source["minecraft_version"] != snapshot or source["protocol_version"] != release["protocol_version"]:
                raise ValueError(f"capture target differs for {snapshot}")
    lock = read(root / "data/source/lock.json")
    if lock["release"] not in releases:
        raise ValueError("data lock selects an unknown release")
    print(f"Validated {len(releases)} releases and their capture availability")


if __name__ == "__main__":
    validate(Path(__file__).resolve().parent.parent)
