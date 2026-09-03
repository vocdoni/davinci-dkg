#!/usr/bin/env python3
"""Stage the pinned circuit artifacts as `<sha256>.<ccs|pk|vk>` release assets.

usage: stage-circuit-release.py <artifacts dir> <output dir>
The hashes come from config/circuit_artifacts.go, so what is uploaded is
exactly what nodes will verify against.
"""
import os, re, shutil, sys

src, dst = sys.argv[1], sys.argv[2]
here = os.path.dirname(os.path.abspath(__file__))
pins = open(os.path.join(here, '..', 'config', 'circuit_artifacts.go')).read()
ext = {'Circuit': 'ccs', 'ProvingKey': 'pk', 'VerificationKey': 'vk'}
os.makedirs(dst, exist_ok=True)
for name, kind, h in re.findall(r'(\w+)(Circuit|ProvingKey|VerificationKey)Hash\s*=\s*"([0-9a-f]{64})"', pins):
    path = os.path.join(src, h)
    if not os.path.exists(path):
        sys.exit(f"missing artifact {name} {kind}: {path}")
    shutil.copyfile(path, os.path.join(dst, f"{h}.{ext[kind]}"))
    print(f"{kind:16s} {name:16s} {h[:12]}  {os.path.getsize(path):>10d} bytes")
