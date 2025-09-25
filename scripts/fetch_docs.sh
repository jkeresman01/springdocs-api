#!/bin/bash

# fetch_docs.sh — Downloads the Spring Boot AsciiDoc files into ./spring-docs

set -e

DOCS_DIR="spring-docs"
TEMP_DIR=".spring-docs-tmp"

echo "Fetching Spring Boot documentation..."

# Clean previous docs
rm -rf "$DOCS_DIR" "$TEMP_DIR"

# Clone with depth=1 for speed
git clone --depth 1 https://github.com/spring-projects/spring-boot.git "$TEMP_DIR"

# Move only the relevant Antora doc files
mv "$TEMP_DIR/spring-boot-project/spring-boot-docs/src/docs/antora" "$DOCS_DIR"

# Cleanup
rm -rf "$TEMP_DIR"

echo "Docs downloaded to ./$DOCS_DIR"

