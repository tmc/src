#\!/bin/bash
# code-to-gpt - prepare code for LLMs

DIRECTORY="."
EXCLUDE_DIRS=("node_modules" "venv" ".venv")
IGNORED_FILES=("go.sum" "go.work.sum" "yarn.lock" "yarn.error.log" "package-lock.json")
COUNT_TOKENS=false
VERBOSE=false

# Default processing logic
process_file() {
  local file="$1"
  echo "== $file =="
  cat "$file"
  echo ""
}

# Main function
main() {
  echo "Processing code in $DIRECTORY"
  find "$DIRECTORY" -type f -name "*.go" | while read -r file; do
    process_file "$file"
  done
}

# Entry point
main "$@"

