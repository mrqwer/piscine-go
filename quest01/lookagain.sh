#!/bin/bash
find . \( -name "*.sh" \) -exec sh -c 'echo "$(basename {} .sh)"' \;
