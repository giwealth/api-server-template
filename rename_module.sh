#!/usr/bin/env bash
# Rename Go module path across this repository (after cloning the template).
# Usage:
#   ./rename_module.sh                     # use basename of repo directory as new module path
#   ./rename_module.sh github.com/you/app  # explicit module path
#   ./rename_module.sh --dry-run github.com/you/app
#
# Rewrites every occurrence of the current "module" from go.mod in tracked files.
# Then run: go mod tidy && go build ./...

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

DRY_RUN=0
POS=()
while [[ $# -gt 0 ]]; do
	case "$1" in
		-n|--dry-run) DRY_RUN=1; shift ;;
		-h|--help)
			echo "Usage: $0 [--dry-run] [NEW_MODULE]"
			echo "  NEW_MODULE defaults to basename of repository directory."
			exit 0
			;;
		*) POS+=("$1"); shift ;;
	esac
done

if [[ ! -f go.mod ]]; then
	echo "error: go.mod not found in $ROOT" >&2
	exit 1
fi

OLD_MODULE="$(awk '/^module[[:space:]]+/ { print $2; exit }' go.mod)"
if [[ -z "$OLD_MODULE" ]]; then
	echo "error: cannot parse module line from go.mod" >&2
	exit 1
fi

if [[ ${#POS[@]} -eq 0 ]]; then
	NEW_MODULE="$(basename "$ROOT")"
else
	NEW_MODULE="${POS[0]}"
fi

if [[ -z "$NEW_MODULE" ]]; then
	echo "error: new module path is empty" >&2
	exit 1
fi

if [[ "$OLD_MODULE" == "$NEW_MODULE" ]]; then
	echo "module unchanged: $NEW_MODULE"
	exit 0
fi

if [[ "$NEW_MODULE" =~ [[:space:]] ]]; then
	echo "error: new module path must not contain whitespace" >&2
	exit 1
fi

echo "Renaming module: $OLD_MODULE -> $NEW_MODULE"

export OLD_MODULE NEW_MODULE

patch_file() {
	local f="$1"
	[[ -f "$f" ]] || return 0
	case "$f" in
		rename_module.sh|.git/*|vendor/*|*/vendor/*) return 0 ;;
	esac
	if ! grep -qF "$OLD_MODULE" "$f" 2>/dev/null; then
		return 0
	fi
	if [[ "$DRY_RUN" -eq 1 ]]; then
		echo "would patch: $f"
		return 0
	fi
	perl -0777 -i -pe 'BEGIN { $o = $ENV{"OLD_MODULE"}; $n = $ENV{"NEW_MODULE"}; } s/\Q$o\E/$n/g' "$f"
}

if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
	# Only patch files that contain the old module path.
	while IFS= read -r f; do
		[[ -n "$f" ]] && patch_file "$f"
	done < <(git grep -l -F "$OLD_MODULE" -- . || true)
else
	while IFS= read -r -d '' f; do
		patch_file "$f"
	done < <(find . -type f \( -name '*.go' -o -name 'go.mod' -o -name '*.yaml' -o -name '*.yml' -o -name '*.md' -o -name '*.html' -o -name '.gitignore' \) -print0 2>/dev/null)
fi

unset OLD_MODULE NEW_MODULE

if [[ "$DRY_RUN" -eq 1 ]]; then
	echo "(dry-run: no files modified)"
	exit 0
fi

# Ensure go.mod module line is explicit (perl already replaced substrings)
echo "Done."
